package utils

import (
	"context"

	"github.com/micro/go-micro/server"
)

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context. If valid, the call is passed along to
// the handler. If not, an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, res interface{}) error {
		// TODO: Validate that JWT is valid

		return nil
	}
}
