package cli

import (
	"context"

	"github.com/spf13/cobra"
)

func QueryCommand(clientFn ClientFunc) *cobra.Command {
	var offset int
	var limit int
	var keysOnly bool
	r := &cobra.Command{
		Use:  "query KIND",
		Args: validateFirstArgAsKind,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := clientFn()
			ctx := context.Background()
			kind := args[0]
			if keysOnly {
				if d, err := client.QueryKeys(ctx, kind, offset, limit); err != nil {
					return err
				} else {
					return formatStrings(d)
				}
			} else {
				if d, err := client.QueryData(ctx, kind, offset, limit); err != nil {
					return err
				} else {
					return formatArray(d)
				}
			}
		},
	}
	r.Flags().IntVar(&offset, "offset", 0, "offset")
	r.Flags().IntVar(&limit, "limit", 10, "limit")
	r.Flags().BoolVar(&keysOnly, "keys-only", false, "KeysOnly")
	return r
}
