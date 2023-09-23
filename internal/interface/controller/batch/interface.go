package batch

import "context"

// LambdaController is a common handler of lambda
type LambdaController interface {
	Handler(ctx context.Context, req interface{}) (interface{}, error)
}
