package helloworld

import (
	"context"
	"github.com/hum2/backend/internal/interface/controller/batch"
)

type Controller struct{}

type response struct {
	Text string `json:"text"`
}

func New() batch.LambdaController {
	return &Controller{}
}

func (c *Controller) Handler(ctx context.Context, req interface{}) (interface{}, error) {
	return response{
		Text: "Hello AWS Lambda World.",
	}, nil
}
