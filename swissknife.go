package swissknife

import (
	"context"
	"encoding/base64"
	"encoding/hex"
)

type service struct {
	UnimplementedSwissknifeSvcServer
}

func (s *service) ConvHexa(_ context.Context, req *ConvHexaReq) (*ConvHexaRes, error) {
	in := req.GetInput()
	out := hex.EncodeToString([]byte(in))
	return &ConvHexaRes{Output: out}, nil
}

func (s *service) ConvBase64(_ context.Context, req *ConvBase64Req) (*ConvBase64Res, error) {
	in := req.GetInput()
	out := base64.StdEncoding.EncodeToString([]byte(in))
	return &ConvBase64Res{Output: out}, nil
}

func New() SwissknifeSvcServer {
	return &service{}
}
