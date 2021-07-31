package main

import (
	"fmt"
	"os"

	"github.com/akm/datastorecli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func keyDecodeCommand() *cobra.Command {
	validateArgs := func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.Errorf("encoded-key is required")
		}
		return nil
	}

	r := &cobra.Command{
		Use:  "decode ENCODED-KEY",
		Args: validateArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			key, err := datastorecli.DecodeKey(args[0])
			if err != nil {
				return err
			}
			fmt.Fprint(os.Stdout, key.String())
			return nil
		},
	}
	return r
}
