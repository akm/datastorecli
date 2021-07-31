package main

import (
	"context"
	"encoding/json"
	"regexp"
	"strconv"

	"cloud.google.com/go/datastore"
	"github.com/akm/datastorecli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func putCommand(clientFn clientFunc) *cobra.Command {
	numberOnly := regexp.MustCompile(`\A\d+\z`)

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
			var key *datastore.Key

			if len(args) == 2 {
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
