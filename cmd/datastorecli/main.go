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

	rootCmd.AddCommand(cli.ConnectableCommandFunc(cli.QueryCommand)())
	rootCmd.AddCommand(cli.ConnectableCommandFunc(cli.GetCommand)())
	rootCmd.AddCommand(cli.ConnectableCommandFunc(cli.PutCommand)())
	rootCmd.AddCommand(cli.ConnectableCommandFunc(cli.DeleteCommand)())
	rootCmd.AddCommand((func() *cobra.Command {
		keyCommand := &cobra.Command{
			Use: "key",
		}
		keyCommand.AddCommand(cli.WithNamespace(cli.KeyEncodeCommand)())
		keyCommand.AddCommand(cli.KeyDecodeCommand())
		return keyCommand
	})())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
