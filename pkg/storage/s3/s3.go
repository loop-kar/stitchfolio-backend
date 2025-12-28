package s3

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/imkarthi24/sf-backend/pkg/errs"
	"github.com/imkarthi24/sf-backend/pkg/storage"
)

// S3Storage implements the storage.StorageProvider interface using AWS S3
type S3Storage struct {
	client     *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
	config     *S3Config
}

// NewS3Storage creates a new S3Storage instance
func NewS3Storage(config *S3Config) (S3Storage, error) {
	client, err := newS3Client(config)
	if err != nil {
		return S3Storage{}, fmt.Errorf("failed to create S3 client: %w", err)
	}

	return S3Storage{
		client:     client,
		uploader:   manager.NewUploader(client),
		downloader: manager.NewDownloader(client),
		config:     config,
	}, nil
}

// Upload implements the StorageProvider.Upload method
func (s S3Storage) Upload(ctx context.Context, key string, reader io.Reader, contentType string) error {
	_, err := s.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.config.Bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return fmt.Errorf("failed to upload file %s: %w", key, err)
	}

	return nil
}

// Download implements the StorageProvider.Download method
func (s S3Storage) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	// Create a buffer to write the file to
	buff := manager.NewWriteAtBuffer([]byte{})

	// Download the file
	_, err := s.downloader.Download(ctx, buff, &s3.GetObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to download file %s: %w", key, err)
	}

	// Create a reader from the buffer
	return io.NopCloser(strings.NewReader(string(buff.Bytes()))), nil
}

// Delete implements the StorageProvider.Delete method
func (s S3Storage) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", key, err)
	}

	return nil
}

// List implements the StorageProvider.List method
func (s S3Storage) List(ctx context.Context, prefix string) ([]storage.FileInfo, error) {
	resp, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.config.Bucket),
		Prefix: aws.String(prefix),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list files with prefix %s: %w", prefix, err)
	}

	files := make([]storage.FileInfo, 0, len(resp.Contents))
	for _, obj := range resp.Contents {
		files = append(files, storage.FileInfo{
			Key:          *obj.Key,
			Size:         *obj.Size,
			LastModified: *obj.LastModified,
			ContentType:  getContentType(obj.Key),
		})
	}

	return files, nil
}

// GetURL implements the StorageProvider.GetURL method
func (s S3Storage) GetURL(ctx context.Context, key string, expires time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(s.client)

	presignedReq, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expires))

	if err != nil {
		return "", fmt.Errorf("failed to generate pre-signed URL for %s: %w", key, err)
	}
	// fmt.Printf("Client Config: Endpoint=%s, Bucket=%s\n", s.config.Endpoint, s.config.Bucket)
	// fmt.Println("presignedReq", presignedReq)
	// fmt.Println("presignedReq url", presignedReq.URL)

	return presignedReq.URL, nil
}

// Exists implements the StorageProvider.Exists method
func (s S3Storage) Exists(ctx context.Context, key string) (bool, error) {
	_, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		var nsk *types.NotFound
		if strings.Contains(err.Error(), "NotFound") || strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "No such key") || strings.Contains(err.Error(), "not exist") || err == nsk {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if file %s exists: %w", key, err)
	}

	return true, nil
}

// GetBucket implements the StorageProvider.GetBucket method
func (s S3Storage) GetBucket() string {
	return s.config.Bucket
}

// Helper function to guess content type from file extension
func getContentType(key *string) string {
	ext := filepath.Ext(*key)
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".pdf":
		return "application/pdf"
	case ".doc", ".docx":
		return "application/msword"
	case ".xls", ".xlsx":
		return "application/vnd.ms-excel"
	case ".txt":
		return "text/plain"
	default:
		return "application/octet-stream"
	}
}

// IsURLExpired checks if a pre-signed S3 URL has expired
func (s S3Storage) IsURLExpired(ctx context.Context, url string, key string) (bool, error) {
	// Extract the X-Amz-Date and X-Amz-Expires parameters from the URL
	parts := strings.Split(url, "?")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid pre-signed URL format")
	}

	params := strings.Split(parts[1], "&")
	var dateStr, expiresStr string

	for _, param := range params {
		if strings.HasPrefix(param, "X-Amz-Date=") {
			dateStr = strings.TrimPrefix(param, "X-Amz-Date=")
		} else if strings.HasPrefix(param, "X-Amz-Expires=") {
			expiresStr = strings.TrimPrefix(param, "X-Amz-Expires=")
		}
	}

	if dateStr == "" || expiresStr == "" {
		return false, fmt.Errorf("missing required parameters in pre-signed URL")
	}

	// Parse the date (format: YYYYMMDDTHHMMSSZ)
	date, err := time.Parse("20060102T150405Z", dateStr)
	if err != nil {
		return false, fmt.Errorf("failed to parse X-Amz-Date: %w", err)
	}

	// Parse the expiration seconds
	expiresSeconds, err := time.ParseDuration(expiresStr + "s")
	if err != nil {
		return false, fmt.Errorf("failed to parse X-Amz-Expires: %w", err)
	}

	// Calculate expiration time and compare with current time
	expirationTime := date.Add(expiresSeconds)
	//if expired, get new url
	if time.Now().UTC().After(expirationTime) {
		return true, nil
	}

	return false, nil
}

func (s S3Storage) GetCurrentOrRenewedURL(ctx *context.Context, url string, key string) (string, error) {
	isExpired, err := s.IsURLExpired(*ctx, url, key)
	if err != nil {
		return "", errs.Wrap(err, "unable to check if URL expired")
	}

	if isExpired {
		// Get renewed S3 URL with new expiry
		renewedUrl, err := s.GetURL(*ctx, key, 24*3*time.Hour)
		if err != nil {
			return "", errs.Wrap(err, "unable to get new presigned url")
		}
		return renewedUrl, nil
	}

	return url, nil

}
