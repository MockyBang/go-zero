package serverinterceptors

import (
	"context"
	"testing"

	"github.com/MockyBang/go-zero/core/prometheus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestUnaryPromMetricInterceptor_Disabled(t *testing.T) {
	_, err := UnaryPrometheusInterceptor(context.Background(), nil, &grpc.UnaryServerInfo{
		FullMethod: "/",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})
	assert.Nil(t, err)
}

func TestUnaryPromMetricInterceptor_Enabled(t *testing.T) {
	prometheus.StartAgent(prometheus.Config{
		Host: "localhost",
		Path: "/",
	})
	_, err := UnaryPrometheusInterceptor(context.Background(), nil, &grpc.UnaryServerInfo{
		FullMethod: "/",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})
	assert.Nil(t, err)
}
