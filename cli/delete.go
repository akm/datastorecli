package cli

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func DeleteCommand(clientFn ClientFunc) *cobra.Command {
	var encodedParent string
	r := &cobra.Command{
		Use:   "delete KIND-OR-ENCODED-KEY [ID-OR-NAME]",
		Short: "Delete an entity",
		Long: `Delete an entity by one of these arguments
- delete KIND ID-KEY
- delete KIND NAME-KEY
- delete ENCODED-KEY
		`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.Errorf("get requires 1 or 2 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client := clientFn()
			key, err := client.Namespace.BuildKey(args, len(args) == 1, false, encodedParent)
			if err != nil {
				return err
			}

			if err := client.Delete(context.Background(), key); err != nil {
				return err
			} else {
				return nil
			}
		},
	}

	r.Flags().StringVar(&encodedParent, "encoded-parent", "", "Encoded parent key")
	return r
}
