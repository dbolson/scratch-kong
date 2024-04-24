package flags_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"scratch-kong/cmd/root"
	"scratch-kong/internal/client"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) ListProjects(ctx context.Context) ([]byte, error) {
	args := c.Called()

	return args.Get(0).([]byte), args.Error(1)
}

func (c *MockClient) GetFlag(ctx context.Context, projKey string, envKey string, key string) ([]byte, error) {
	args := c.Called(projKey, envKey, key)

	return args.Get(0).([]byte), args.Error(1)
}

func MockClientFn(c *MockClient) client.ClientFn {
	return func(accessToken string, baseURI string, version string) client.Client {
		return c
	}
}

func TestGet(t *testing.T) {
	mockArgs := []interface{}{
		"test-proj-key",
		"test-env-key",
		"test-flag-key",
	}

	t.Run("with valid flags calls API", func(t *testing.T) {
		output := new(bytes.Buffer)
		validResponse := `{"valid": true}`
		client := MockClient{}
		client.
			On("GetFlag", mockArgs...).
			Return([]byte(validResponse), nil)
		args := []string{
			"flags", "get",
			"--access-token", "test-access-token",
			"--environment", "test-env-key",
			"--project", "test-proj-key",
			"--flag", "test-flag-key",
		}
		ctx := root.NewRootCmd(MockClientFn(&client), args, kong.Writers(output, output))

		err := ctx.Run()

		require.NoError(t, err)
		assert.Equal(t, validResponse+"\n", output.String())
	})

	t.Run("with an error response is an error", func(t *testing.T) {
		output := new(bytes.Buffer)
		client := MockClient{}
		client.
			On("GetFlag", mockArgs...).
			Return([]byte(`{}`), errors.New("an error"))
		args := []string{
			"flags", "get",
			"--access-token", "test-access-token",
			"--environment", "test-env-key",
			"--project", "test-proj-key",
			"--flag", "test-flag-key",
		}
		ctx := root.NewRootCmd(MockClientFn(&client), args, kong.Writers(output, output))

		err := ctx.Run()

		assert.EqualError(t, err, "an error")
		assert.Empty(t, output.String())
	})
}
