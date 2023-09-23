package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	lambda.Start(app.Ctr.Handler)
}
