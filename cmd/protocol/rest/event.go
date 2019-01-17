package rest

import (
	"context"
	"gemini/pb/v1/event"
)

// WithEventServer ..
func WithEventServer() OptionFunc {
	return func(ctx context.Context, port string) {
		event.RegisterEventServiceHandlerFromEndpoint(ctx, mux, port, dialOpts)
	}
}
