package aws

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/oneee-playground/r2d2-image-builder/config"
)

var sqsOnce sync.Once

var sqsClient *sqs.Client

func SQSClient() *sqs.Client {
	sqsOnce.Do(func() {
		sqsClient = sqs.NewFromConfig(getAWSConfig())
	})
	return sqsClient
}

func getAWSConfig() aws.Config {
	credProvider := credentials.NewStaticCredentialsProvider(config.AWSAccessKeyID, config.AWSSecretKey, "")

	conf := aws.Config{
		Region:      config.AWSRegion,
		Credentials: aws.NewCredentialsCache(credProvider),
	}

	return conf
}
