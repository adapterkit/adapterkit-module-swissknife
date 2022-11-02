package swissknife_test

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/stretchr/testify/assert"

	swissknife "github.com/pmg-tools/adapterkit-module-swissknife"
)

func Dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	service := &swissknife.Service{}
	swissknife.RegisterSwissknifeSvcServer(server, service)
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestSwissknifeSvcConvHexa_grpc(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(Dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := swissknife.NewSwissknifeSvcClient(conn)
	patient := []struct {
		input    string
		expected string
	}{
		{"hello", "68656c6c6f"},
		{"hello world", "68656c6c6f20776f726c64"},
		{"Zetest", "5a6574657374"},
	}

	for _, p := range patient {
		req := &swissknife.ConvHexaReq{Input: p.input}
		resp, err := client.ConvHexa(ctx, req)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, p.expected, resp.GetOutput(), "%s must be equal to %s", resp.GetOutput(), p.expected)

	}
}

func TestSwissknifeSvcConvBase64_grpc(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(Dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := swissknife.NewSwissknifeSvcClient(conn)

	patient := []struct {
		input    string
		expected string
	}{
		{"hello", "aGVsbG8="},
		{"hello world", "aGVsbG8gd29ybGQ="},
		{"Zetest", "WmV0ZXN0"},
	}

	for _, p := range patient {
		req := &swissknife.ConvBase64Req{Input: p.input}
		resp, err := client.ConvBase64(ctx, req)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, p.expected, resp.GetOutput(), "%s must be equal to %s", resp.GetOutput(), p.expected)
	}
}
