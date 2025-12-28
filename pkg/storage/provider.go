package storage

import (
	"context"
	"io"
	"time"
)

type StorageProvider interface {
	// Upload stores a file in the storage
	Upload(ctx context.Context, key string, reader io.Reader, contentType string) error

	// Download retrieves a file from storage
	Download(ctx context.Context, key string) (io.ReadCloser, error)

	// Delete removes a file from storage
	Delete(ctx context.Context, key string) error

	// List returns information about files with the given prefix
	List(ctx context.Context, prefix string) ([]FileInfo, error)

	// Exists checks if a file exists
	Exists(ctx context.Context, key string) (bool, error)
}

type CloudStorageProvider interface {
	StorageProvider

	// GetURL returns a URL to access the file (can be pre-signed if needed)
	GetURL(ctx context.Context, key string, expires time.Duration) (string, error)

	// GetBucket returns the name of the bucket
	GetBucket() string

	// IsURLExpired checks if the URL is expired
	IsURLExpired(ctx context.Context, url string, key string) (bool, error)

	// GetCurrentOrRenewedURL gets the current presigned url.
	// If expired, it generates a new URL and returns it
	GetCurrentOrRenewedURL(ctx *context.Context, url string, key string) (string, error)
}
