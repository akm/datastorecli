package cli

import (
	"context"

	"github.com/akm/datastorecli/formatters"
	"github.com/spf13/cobra"
)

func QueryKeysOnlyCommand(clientFn clientFunc) *cobra.Command {
	var offset int
	var limit int
	var formatterName string
	r := &cobra.Command{
		Use:  "query KIND",
		Args: validateFirstArgAsKind,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := clientFn()
			ctx := context.Background()
			kind := args[0]
			if d, err := client.QueryKeys(ctx, kind, offset, limit); err != nil {
				return err
			} else {
				return formatters.FormatStrings(formatterName, d)
			}
		},
	}
	r.Flags().IntVar(&offset, "offset", 0, "offset")
	r.Flags().IntVar(&limit, "limit", 10, "limit")
	r.Flags().StringVar(&formatterName, "formatter", "strings", "options: strings, json, pretty-json")
	return r
}

var QueryKeysOnly = connectable(QueryKeysOnlyCommand)
