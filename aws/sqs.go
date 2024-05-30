package aws

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	conf "github.com/oneee-playground/r2d2-image-builder/config"
	"github.com/pkg/errors"
)

var awsConfig aws.Config

var sqsOnce sync.Once

var sqsClient *sqs.Client

func SQSClient() *sqs.Client {
	sqsOnce.Do(func() {
		sqsClient = sqs.NewFromConfig(awsConfig)
	})
	return sqsClient
}

func LoadConfig(ctx context.Context) error {
	conf, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(conf.AWSRegion))
	if err != nil {
		return errors.Wrap(err, "loading aws config")
	}

	awsConfig = conf

	return nil
}
