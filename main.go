package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/oneee-playground/r2d2-image-builder/function"
	"github.com/oneee-playground/r2d2-image-builder/util"
	"github.com/sirupsen/logrus"
)

func init() {
	// TODO: Maybe change this.
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	if err := util.InitFS(); err != nil {
		logrus.Fatal(err)
	}

	lambda.Start(function.HandleBuildImage)
}
