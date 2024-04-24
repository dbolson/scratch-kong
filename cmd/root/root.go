package root

import (
	"io"
	"scratch-kong/cmd/flags"
	"scratch-kong/internal/client"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
)

type RootCmd struct {
	GlobalFlags
	Flags flags.FlagsCmd `kong:"cmd,help='Make requests (list, create, etc.) on flags',group='flags'"`
}

type GlobalFlags struct {
	AccessToken string `kong:"required,help='LaunchDarkly API token with write-level access',type='string',env='ACCESS_TOKEN',envprefix='LD_'"`
	BaseURI     string `kong:"help='LaunchDarkly base URI',default='https://app.launchdarkly.com',type='string',env='BASE_URI',envprefix='LD_'"`
}

func NewRootCmd(
	clientFn client.ClientFn,
	args []string,
	options ...kong.Option,
) *kong.Context {
	var rootCmd RootCmd

	defaultOpts := []kong.Option{
		kong.Configuration(kongyaml.Loader, "./config.yml"),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Description("LaunchDarkly CLI to control your feature flags"),
		kong.Name("ldcli"),
		kong.UsageOnError(),
	}
	options = append(options, defaultOpts...)

	parser, err := kong.New(&rootCmd, options...)
	if err != nil {
		panic(err)
	}

	ctx, err := parser.Parse(args)
	parser.FatalIfErrorf(err)

	c := clientFn(
		rootCmd.AccessToken,
		rootCmd.BaseURI,
		"0.1.0",
	)
	ctx.BindTo(c, (*client.Client)(nil))
	ctx.BindTo(parser.Stdout, (*io.Writer)(nil))

	return ctx
}
