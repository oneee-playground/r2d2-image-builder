package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	aws_util "github.com/oneee-playground/r2d2-image-builder/aws"
	"github.com/pkg/errors"
)

type BuildImageEvent struct {
	ID      uuid.UUID     `json:"id"`
	Success bool          `json:"success"`
	Took    time.Duration `json:"took"`
	Extra   string        `json:"extra"`
}

func Publish(ctx context.Context, id uuid.UUID, took time.Duration, err error) error {
	client := aws_util.SQSClient()

	payload, err := json.Marshal(buildEvent(id, took, err))
	if err != nil {
		return errors.Wrap(err, "marshalling payload")
	}

	input := &sqs.SendMessageInput{
		MessageBody: aws.String(string(payload)),
		QueueUrl:    aws.String("notfound"),
	}

	if _, err := client.SendMessage(ctx, input); err != nil {
		return errors.Wrap(err, "sending message")
	}

	return nil
}

func buildEvent(id uuid.UUID, took time.Duration, err error) *BuildImageEvent {
	res := BuildImageEvent{ID: id, Took: took, Success: true}
	if err != nil {
		res.Success = false
		res.Extra = err.Error()
	}
	return &res
}
