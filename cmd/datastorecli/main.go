package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "datastorecli",
	}

	rootCmd.AddCommand(connectableCommandFunc(queryCommand)())
	rootCmd.AddCommand(connectableCommandFunc(getCommand)())
	rootCmd.AddCommand((func() *cobra.Command {
		keyCommand := &cobra.Command{
			Use: "key",
		}
		keyCommand.AddCommand(keyEncodeCommand())
		keyCommand.AddCommand(keyDecodeCommand())
		return keyCommand
	})())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
