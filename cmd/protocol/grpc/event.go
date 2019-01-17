package grpc

import (
	"gemini/pb/v1/event"

	"google.golang.org/grpc"
)

// WithEventServer ..
func WithEventServer(eventSrv event.EventServiceServer) OptionFunc {
	return func(srv *grpc.Server) {
		event.RegisterEventServiceServer(srv, eventSrv)
	}
}
