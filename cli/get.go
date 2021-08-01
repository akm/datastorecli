package cli

import (
	"context"

	"github.com/akm/datastorecli/formatters"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func GetCommand(clientFn clientFunc) *cobra.Command {
	var encodedParent string
	var formatterName string
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
			client := clientFn()
			key, err := client.Namespace.BuildKey(args, len(args) == 1, false, encodedParent)
			if err != nil {
				return err
			}

			if d, err := client.Get(context.Background(), key); err != nil {
				return err
			} else {
				return formatters.NewDefaultWriter().FormatData(d)
			}
		},
	}

	r.Flags().StringVar(&encodedParent, "encoded-parent", "", "Encoded parent key")
	r.Flags().StringVar(&formatterName, "formatter", "json", "options: json or pretty-json")
	return r
}

var Get = connectable(GetCommand)
