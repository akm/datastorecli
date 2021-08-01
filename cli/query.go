package cli

import (
	"context"

	"github.com/akm/datastorecli/formatters"
	"github.com/spf13/cobra"
)

func QueryCommand(clientFn clientFunc) *cobra.Command {
	var offset int
	var limit int
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
				return formatters.NewDefaultWriter().FormatArray(d)
			}
		},
	}
	r.Flags().IntVar(&offset, "offset", 0, "offset")
	r.Flags().IntVar(&limit, "limit", 10, "limit")
	return r
}

var Query = connectable(QueryCommand)
