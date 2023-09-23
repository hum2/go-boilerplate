package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hum2/backend/internal/interface/controller/batch/helloworld"
)

func main() {
	app := helloworld.New()
	lambda.Start(app.Handler)
}
