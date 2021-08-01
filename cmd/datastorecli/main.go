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
		(func() *cobra.Command {
			keyCommand := &cobra.Command{Use: "key"}
			keyCommand.AddCommand(
				cli.WithNamespace(cli.KeyEncodeCommand)(),
				cli.KeyDecodeCommand(),
			)
			return keyCommand
		})(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
