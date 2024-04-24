package flags_test

import (
	"context"

	"github.com/stretchr/testify/mock"

	"scratch-kong/internal/client"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) GetFlag(ctx context.Context, projKey string, envKey string, key string) ([]byte, error) {
	args := c.Called(projKey, envKey, key)

	return args.Get(0).([]byte), args.Error(1)
}

func (c *MockClient) ListFlags(ctx context.Context, projKey string) ([]byte, error) {
	args := c.Called(projKey)

	return args.Get(0).([]byte), args.Error(1)
}

func MockClientFn(c *MockClient) client.ClientFn {
	return func(accessToken string, baseURI string, version string) client.Client {
		return c
	}
}
