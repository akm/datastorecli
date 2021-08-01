package main

import (
	"fmt"
	"os"

	"github.com/akm/datastorecli/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "datastorecli",
	}

	rootCmd.AddCommand(
		cli.ConnectableCommandFunc(cli.QueryCommand)(),
		cli.ConnectableCommandFunc(cli.GetCommand)(),
		cli.ConnectableCommandFunc(cli.PutCommand)(),
		cli.ConnectableCommandFunc(cli.DeleteCommand)(),
		cli.GroupCommand("key",
			cli.WithNamespace(cli.KeyEncodeCommand)(),
			cli.KeyDecodeCommand(),
		),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
