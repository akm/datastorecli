package main

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func validateFirstArgAsKind(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.Errorf("kind is required")
	}
	return nil
}
