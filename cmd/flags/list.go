package flags

import (
	"context"
	"fmt"
	"io"

	"scratch-kong/internal/client"
)

type ListCmd struct {
	Project string `kong:"required,type='string',help='project key'"`
}

func (cmd *ListCmd) Run(outwriter io.Writer, client client.Client) error {
	flag, err := client.ListFlags(
		context.Background(),
		cmd.Project,
	)
	if err != nil {
		return err
	}

	fmt.Fprintf(outwriter, string(flag)+"\n")

	return nil
}
