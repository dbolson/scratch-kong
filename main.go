package main

import (
	"os"
	"scratch-kong/cmd/root"
	"scratch-kong/internal/client"
)

func main() {
	// var (
	// 	rootCmd root.RootCmd
	// 	// globals Globals
	// )

	ctx := root.NewRootCmd(client.LDClientFn(), os.Args[1:])
	// ctx := kong.Parse(&rootCmd,
	// 	// kong.Bind(&globals),
	// 	kong.Configuration(kongyaml.Loader, "./config.yml"),
	// 	kong.ConfigureHelp(kong.HelpOptions{
	// 		Compact: true,
	// 	}),
	// 	kong.Description("LaunchDarkly CLI to control your feature flags"),
	// 	kong.Name("ldcli"),
	// 	kong.UsageOnError(),
	// )

	// c := client.NewLDClient(
	// 	rootCmd.AccessToken,
	// 	rootCmd.BaseURI,
	// 	"0.1.0",
	// )
	// ctx.BindTo(&c, (*client.Client)(nil))

	// err := ctx.Run(&Globals{AccessToken: rootCmd.AccessToken})
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
