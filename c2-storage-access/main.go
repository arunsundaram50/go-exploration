package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// getS3Client initializes an S3 client based on the AWS_PROFILE environment variable.
func getAwsConfig() (aws.Config, error) {
	profile := os.Getenv("AWS_PROFILE")

	// AWS configuration options
	opts := []func(*config.LoadOptions) error{
		config.WithSharedConfigProfile(profile),
	}

	// If LocalStack is detected, set the custom endpoint
	if profile == "localstack" {
		opts = append(opts, config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: "http://localhost:4566",
				}, nil
			},
		)))
	}

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		return *aws.NewConfig(), fmt.Errorf("error loading AWS config: %w", err)
	}

	return cfg, nil
}

// readS3Object fetches and prints an S3 object
func readS3Object(bucket, key string) error {
	cfg, err := getAwsConfig()
	if err != nil {
		return err
	}

	// GOTCHA
	// localStack defaults to path-style access (http://localhost:4566/my-bucket/test.txt)
	// AWS defaults to virtual-hosted style access (http://my-bucket.localhost:4566/test.txt)
	// Set o.UsePathStyle to true, so that it works when AWS_PROFILE is set to localstack or cloud
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	// Get the object from S3
	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("error fetching object: %w", err)
	}
	defer resp.Body.Close()

	// Stream output to stdout
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return fmt.Errorf("error reading object data: %w", err)
	}

	return nil
}

func main() {
	bucket := "my-bucket"
	key := "test.txt"

	if err := readS3Object(bucket, key); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
