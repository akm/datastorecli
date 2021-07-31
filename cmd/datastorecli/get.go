package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func getCommand(clientFn clientFunc) *cobra.Command {
	var encodedParent string
	r := &cobra.Command{
		Use:   "get KIND-OR-ENCODED-KEY [ID-OR-NAME]",
		Short: "Get an entity",
		Long: `Get an entity by one of these arguments
- get KIND ID-KEY
- get KIND NAME-KEY
- get ENCODED-KEY
		`,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := validateFirstArgAsKind(cmd, args); err != nil {
				return err
			}
			if len(args) < 1 {
				return errors.Errorf("get requires 1 or 2 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			key, err := buildKey(args, len(args) == 1, false, encodedParent)
			if err != nil {
				return err
			}

			client := clientFn()
			if d, err := client.Get(context.Background(), key); err != nil {
				return err
			} else {
				return formatData(d)
			}
		},
	}

	r.Flags().StringVar(&encodedParent, "encoded-parent", "", "Encoded parent key")
	return r
}
