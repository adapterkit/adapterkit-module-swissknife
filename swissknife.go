package swissknife

import (
	"context"
	"encoding/base64"
	"encoding/hex"
)

type Service struct {
	UnimplementedSwissknifeSvcServer
}

func (s *Service) ConvHexa(_ context.Context, req *ConvHexa_Request) (*ConvHexa_Response, error) {
	in := req.GetInput()
	out := hex.EncodeToString([]byte(in))
	return &ConvHexa_Response{Output: out}, nil
}

func (s *Service) ConvBase64(_ context.Context, req *ConvBase64_Request) (*ConvBase64_Response, error) {
	in := req.GetInput()
	out := base64.StdEncoding.EncodeToString([]byte(in))
	return &ConvBase64_Response{Output: out}, nil
}
