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
		cli.Connectable(cli.QueryCommand)(),
		cli.Connectable(cli.GetCommand)(),
		cli.Connectable(cli.PutCommand)(),
		cli.Connectable(cli.DeleteCommand)(),
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
