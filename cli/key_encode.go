package cli

import (
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/akm/datastorecli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func KeyEncodeCommand(fn NamespaceFunc) *cobra.Command {
	var id int64
	var name string
	var encodedParent string
	r := &cobra.Command{
		Use:  "encode KIND",
		Args: validateFirstArgAsKind,
		RunE: func(cmd *cobra.Command, args []string) error {
			kind := args[0]

			parentKey, err := datastorecli.DecodeKey(encodedParent)
			if err != nil {
				return err
			}

			var key *datastore.Key
			if id != 0 {
				key = datastore.IDKey(kind, id, parentKey)
			} else if name != "" {
				key = datastore.NameKey(kind, name, parentKey)
			} else {
				return errors.Errorf("key encode requires id or name")
			}
			namespace := fn()
			namespace.PrepareKey(key)

			fmt.Fprint(os.Stdout, key.Encode())
			return nil
		},
	}
	r.Flags().Int64Var(&id, "id", int64(0), "id")
	r.Flags().StringVar(&name, "name", "", "name")
	r.Flags().StringVar(&encodedParent, "encoded-parent", "", "Encoded parent key")

	return r
}

var KeyEncode = WithNamespace(KeyEncodeCommand)
