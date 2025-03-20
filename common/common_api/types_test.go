package common_api

import (
	"context"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MockServiceServer is a mock implementation of ServiceServer interface
type MockServiceServer struct {
	mock.Mock
}

// Register implements ServiceServer.Register
func (m *MockServiceServer) Register(grpcServer *grpc.Server) {
	m.Called(grpcServer)
}

// RegisterFromEndpoint implements ServiceServer.RegisterFromEndpoint
func (m *MockServiceServer) RegisterFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	args := m.Called(ctx, mux, endpoint, opts)
	return args.Error(0)
}

func TestServiceServerInterface(t *testing.T) {
	// Create a mock instance
	mockServer := new(MockServiceServer)

	// Test that our mock implements the ServiceServer interface
	var _ ServiceServer = mockServer

	// Create test parameters
	grpcServer := grpc.NewServer()
	ctx := context.Background()
	mux := runtime.NewServeMux()
	endpoint := "localhost:8080"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Setup expectations
	mockServer.On("Register", grpcServer).Return()
	mockServer.On("RegisterFromEndpoint", ctx, mux, endpoint, opts).Return(nil)

	// Call the methods
	mockServer.Register(grpcServer)
	err := mockServer.RegisterFromEndpoint(ctx, mux, endpoint, opts)

	// Assert expectations
	assert.NoError(t, err)
	mockServer.AssertExpectations(t)
}

func TestServiceServerUsage(t *testing.T) {
	// This test demonstrates how ServiceServer would be used in an actual application

	// Create a mock implementation
	mockService := new(MockServiceServer)

	// Setup some function that uses ServiceServer interface
	registerServices := func(services []ServiceServer, grpcServer *grpc.Server) {
		for _, service := range services {
			service.Register(grpcServer)
		}
	}

	setupGateway := func(services []ServiceServer, ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		for _, service := range services {
			if err := service.RegisterFromEndpoint(ctx, mux, endpoint, opts); err != nil {
				return err
			}
		}
		return nil
	}

	// Create test parameters
	grpcServer := grpc.NewServer()
	ctx := context.Background()
	mux := runtime.NewServeMux()
	endpoint := "localhost:8080"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Setup expectations
	mockService.On("Register", grpcServer).Return()
	mockService.On("RegisterFromEndpoint", ctx, mux, endpoint, opts).Return(nil)

	// Call the functions
	services := []ServiceServer{mockService}
	registerServices(services, grpcServer)
	err := setupGateway(services, ctx, mux, endpoint, opts)

	// Assert expectations
	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}
