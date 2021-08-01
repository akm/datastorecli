package cli

import (
	"context"
	"encoding/json"

	"github.com/akm/datastorecli/formatters"
	"github.com/akm/datastorecli/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func PutCommand(clientFn clientFunc) *cobra.Command {
	var encodedParent string
	var incompleteKey bool
	// var formatterName string
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
			client := clientFn()
			key, err := client.Namespace.BuildKey(args, len(args) == 2, incompleteKey, encodedParent)
			if err != nil {
				return err
			}

			var dataIndex int
			if len(args) == 2 {
				dataIndex = 1
			} else {
				dataIndex = 2
			}

			src := models.AnyEntity{}
			if err := json.Unmarshal([]byte(args[dataIndex]), &src); err != nil {
				return err
			}

			if resKey, err := client.Put(context.Background(), key, src); err != nil {
				return err
			} else {
				return formatters.FormatStringer("strings", resKey)
			}
		},
	}

	r.Flags().StringVar(&encodedParent, "encoded-parent", "", "Encoded parent key")
	r.Flags().BoolVar(&incompleteKey, "incomplete-key", false, "Incomplete key")
	// r.Flags().StringVar(&formatterName, "format", "string", "options: string")
	return r
}

var Put = connectable(PutCommand)
