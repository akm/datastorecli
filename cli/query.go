package cli

import (
	"context"

	"github.com/akm/datastorecli/formatters"
	"github.com/spf13/cobra"
)

func QueryCommand(clientFn clientFunc) *cobra.Command {
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
			if d, err := client.QueryData(ctx, kind, offset, limit); err != nil {
				return err
			} else {
				return formatters.FormatArray(formatterName, d)
			}
		},
	}
	r.Flags().IntVar(&offset, "offset", 0, "offset")
	r.Flags().IntVar(&limit, "limit", 10, "limit")
	r.Flags().StringVar(&formatterName, "format", "ndjson", "options: ndjson, json or pretty-json")
	return r
}

var Query = connectable(QueryCommand)
