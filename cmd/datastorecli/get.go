package main

import (
	"context"
	"regexp"
	"strconv"

	"cloud.google.com/go/datastore"
	"github.com/akm/datastorecli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func getCommand(clientFn clientFunc) *cobra.Command {
	numberOnly := regexp.MustCompile(`\A\d+\z`)

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
			var key *datastore.Key
			if len(args) == 1 {
				var err error
				key, err = datastorecli.DecodeKey(args[0])
				if err != nil {
					return err
				}
			} else {
				kind := args[0]

				parentKey, err := datastorecli.DecodeKey(encodedParent)
				if err != nil {
					return err
				}

				if numberOnly.MatchString(args[1]) {
					id, err := strconv.ParseInt(args[1], 10, 64)
					if err != nil {
						return err
					}
					key = datastore.IDKey(kind, id, parentKey)
				} else {
					key = datastore.NameKey(kind, args[1], parentKey)
				}
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
