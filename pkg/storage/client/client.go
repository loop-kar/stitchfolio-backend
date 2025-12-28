package client

import (
	"github.com/imkarthi24/sf-backend/pkg/storage"
	"github.com/imkarthi24/sf-backend/pkg/storage/s3"
)

var storageClient storage.StorageProvider

func ProvideCloudStorageClient(config s3.S3Config) (storage.CloudStorageProvider, error) {
	s3Storage, err := s3.NewS3Storage(&config)
	if err != nil {
		return nil, err
	}

	storageClient = s3Storage
	return s3Storage, nil
}

func ProvideLocalStorageClient(config s3.S3Config) (storage.StorageProvider, error) {
	s3Storage, err := s3.NewS3Storage(&config)
	if err != nil {
		return nil, err
	}
	storageClient = s3Storage
	return s3Storage, nil
}

func GetClient() storage.StorageProvider {
	return storageClient
}
