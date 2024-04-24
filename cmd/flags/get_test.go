package flags_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"scratch-kong/cmd/root"
)

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
