package hellowire

import (
	"context"
	"github.com/hum2/backend/internal/interface/controller/batch"
	usecase "github.com/hum2/backend/internal/usecase/batch/hellowire"
)

type Controller struct {
	usecase *usecase.Usecase
}

type response struct {
	Text string `json:"text"`
}

func New(u *usecase.Usecase) batch.LambdaController {
	return &Controller{
		usecase: u,
	}
}

func (c *Controller) Handler(ctx context.Context, req interface{}) (interface{}, error) {
	text := c.usecase.FindAll(ctx)
	return response{
		Text: text,
	}, nil
}
