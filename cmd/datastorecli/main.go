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
		cli.Query(),
		cli.Get(),
		cli.Put(),
		cli.Delete(),
		cli.Group("key",
			cli.KeyEncode(),
			cli.KeyDecode(),
		),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
