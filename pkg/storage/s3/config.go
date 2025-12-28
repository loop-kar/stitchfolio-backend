package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

// S3Config holds configuration for S3 storage
type S3Config struct {
	Region          string
	Bucket          string
	AccessKeyID     string
	SecretAccessKey string
	Endpoint        string
	UsePathStyle    bool
	ForceHTTPS      bool
}

// newS3Client creates and configures a new S3 client
func newS3Client(cfg *S3Config) (*s3.Client, error) {
	optFuncs := []func(*config.LoadOptions) error{
		config.WithRegion(cfg.Region),
		config.WithEndpointResolver(
			aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           cfg.Endpoint, // your R2 endpoint
					SigningRegion: "us-east-1",  // required for signing
				}, nil
			}),
		),
	}

	// If credentials are provided, use them
	if cfg.AccessKeyID != "" && cfg.SecretAccessKey != "" {
		optFuncs = append(optFuncs, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		))
	}

	awsCfg, err := config.LoadDefaultConfig(context.Background(), optFuncs...)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// s3Opts := []func(*s3.Options){
	// 	// Configure for local development or testing if endpoint is provided
	// 	func(o *s3.Options) {
	// 		if cfg.Endpoint != "" {
	// 			o.BaseEndpoint = aws.String(cfg.Endpoint)
	// 		}
	// 		if cfg.UsePathStyle {
	// 			o.UsePathStyle = true
	// 		}
	// 	},
	// }

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = false
	})

	return client, nil
}

// func newS3Client(cfg *S3Config) (*s3.Client, error) {
// 	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion("auto"), // R2 accepts 'auto'
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
// 			cfg.AccessKeyID,
// 			cfg.SecretAccessKey,
// 			"",
// 		)),
// 		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
// 			return aws.Endpoint{
// 				URL:               cfg.Endpoint, // your R2 endpoint
// 				SigningRegion:     "us-east-1",  // force signing region
// 				HostnameImmutable: true,         // Optional: disables region-based hostname rewriting
// 			}, nil
// 		})),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load AWS config: %w", err)
// 	}

// 	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
// 		o.UsePathStyle = cfg.UsePathStyle
// 	})
// 	return client, nil
// }

// LoadS3ConfigFromEnv loads S3 configuration from environment variables
func LoadS3ConfigFromEnv() (*S3Config, error) {
	// In a real application, you'd use something like viper or godotenv
	// to load config from environment variables
	// This is a placeholder implementation

	return &S3Config{
		Region:          "us-east-1", //viper.GetString("s3Config.region"),
		Bucket:          viper.GetString("s3Config.bucket"),
		AccessKeyID:     viper.GetString("s3Config.accessKey"),
		SecretAccessKey: viper.GetString("s3Config.secretAccessKey"),
		Endpoint:        viper.GetString("s3Config.endpoint"),
		UsePathStyle:    viper.GetBool("s3Config.usePathStyle"),
		ForceHTTPS:      viper.GetBool("s3Config.forceHTTPS"),
	}, nil
}

func ProvideS3Config(config *S3Config) S3Config {
	return *config
}

// Helper functions for environment variables
func getEnvOrDefault(key, defaultValue string) string {
	// In a real app, use os.Getenv
	return os.Getenv(key)
}

func getEnvAsBool(key string, defaultValue bool) bool {
	// In a real app, parse from os.Getenv
	return defaultValue
}
