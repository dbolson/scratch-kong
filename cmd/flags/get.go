package flags

import (
	"context"
	"fmt"
	"io"

	"scratch-kong/internal/client"
)

type GetCmd struct {
	Environment string `kong:"required,type='string',help='environment key'"`
	Flag        string `kong:"required,type='string',help='flag key'"`
	Project     string `kong:"required,type='string',help='project key'"`
}

func (cmd *GetCmd) Run(outwriter io.Writer, client client.Client) error {
	flag, err := client.GetFlag(
		context.Background(),
		cmd.Project,
		cmd.Environment,
		cmd.Flag,
	)
	if err != nil {
		return err
	}

	fmt.Fprintf(outwriter, string(flag)+"\n")

	return nil
}

func (cmd *GetCmd) Help() string {
	return "Get a flag to check its details"
}
