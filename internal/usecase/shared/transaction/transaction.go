package transaction

import "context"

// Handler is an interface of transaction
type Handler interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error)
}
