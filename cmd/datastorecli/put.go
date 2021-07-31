package main

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func putCommand(clientFn clientFunc) *cobra.Command {
	var encodedParent string
	r := &cobra.Command{
		Use:   "put KIND-OR-ENCODED-KEY [ID-OR-NAME-OR-JSON-DATA] [JSON-DATA]",
		Short: "Pet an entity",
		Long: `Put an entity by one of these arguments
- put KIND ID-KEY JSON-DATA
- put KIND NAME-KEY JSON-DATA
- put ENCODED-KEY JSON-DATA
		`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.Errorf("put requires 2 or 3 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			key, err := buildKey(args, len(args) == 2, encodedParent)
			if err != nil {
				return err
			}

			var dataIndex int
			if len(args) == 2 {
				dataIndex = 1
			} else {
				dataIndex = 2
			}

			var src interface{}
			if err := json.Unmarshal([]byte(args[dataIndex]), &src); err != nil {
				return err
			}

			client := clientFn()
			if resKey, err := client.Put(context.Background(), key, src); err != nil {
				return err
			} else {
				return formatStringer(resKey)
			}
		},
	}

	r.Flags().StringVar(&encodedParent, "encoded-parent", "", "Encoded parent key")
	return r
}
