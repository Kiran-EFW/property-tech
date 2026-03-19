package storage

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

// StorageProvider defines the interface for object storage operations.
type StorageProvider interface {
	Upload(ctx context.Context, key string, body io.Reader, contentType string) (string, error)
	Delete(ctx context.Context, key string) error
	GetPresignedURL(ctx context.Context, key string, expiry time.Duration) (string, error)
}

// R2Provider implements StorageProvider using Cloudflare R2 (S3-compatible).
type R2Provider struct {
	client    *s3.Client
	bucket    string
	publicURL string
}

// NewR2Provider creates a new R2Provider.
func NewR2Provider(accountID, accessKeyID, accessKeySecret, bucket, publicURL string) (*R2Provider, error) {
	r2Endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKeyID,
			accessKeySecret,
			"",
		)),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, fmt.Errorf("r2: failed to load config: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(r2Endpoint)
		o.UsePathStyle = true
	})

	return &R2Provider{
		client:    client,
		bucket:    bucket,
		publicURL: publicURL,
	}, nil
}

// Upload uploads an object to R2 and returns its public URL.
func (r *R2Provider) Upload(ctx context.Context, key string, body io.Reader, contentType string) (string, error) {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	}

	_, err := r.client.PutObject(ctx, input)
	if err != nil {
		return "", fmt.Errorf("r2 upload: %w", err)
	}

	publicURL := fmt.Sprintf("%s/%s", r.publicURL, key)

	log.Info().
		Str("key", key).
		Str("bucket", r.bucket).
		Str("url", publicURL).
		Msg("file uploaded to R2")

	return publicURL, nil
}

// Delete removes an object from R2.
func (r *R2Provider) Delete(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	}

	_, err := r.client.DeleteObject(ctx, input)
	if err != nil {
		return fmt.Errorf("r2 delete: %w", err)
	}

	log.Info().
		Str("key", key).
		Str("bucket", r.bucket).
		Msg("file deleted from R2")

	return nil
}

// GetPresignedURL generates a presigned URL for temporary access to a private object.
func (r *R2Provider) GetPresignedURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(r.client)

	input := &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	}

	presignResult, err := presignClient.PresignGetObject(ctx, input, s3.WithPresignExpires(expiry))
	if err != nil {
		return "", fmt.Errorf("r2 presign: %w", err)
	}

	return presignResult.URL, nil
}

// DetectContentType detects the MIME type of the given data.
func DetectContentType(data []byte) string {
	return http.DetectContentType(data)
}
