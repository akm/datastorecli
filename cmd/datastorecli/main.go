package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/akm/datastorecli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "datastorecli",
	}

	rootCmd.AddCommand(connectableCommandFunc(queryCommand)())

	rootCmd.AddCommand(connectableCommandFunc(getCommand)())

	rootCmd.AddCommand((func() *cobra.Command {
		keyCommand := &cobra.Command{
			Use: "key",
		}
		keyCommand.AddCommand((func() *cobra.Command {
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

					fmt.Fprint(os.Stdout, key.Encode())
					return nil
				},
			}
			r.Flags().Int64Var(&id, "id", int64(0), "id")
			r.Flags().StringVar(&name, "name", "", "name")
			r.Flags().StringVar(&encodedParent, "encoded-parent", "", "Encoded parent key")

			return r
		})())

		keyCommand.AddCommand((func() *cobra.Command {
			validateArgs := func(cmd *cobra.Command, args []string) error {
				if len(args) < 1 {
					return errors.Errorf("encoded-key is required")
				}
				return nil
			}

			r := &cobra.Command{
				Use:  "decode ENCODED-KEY",
				Args: validateArgs,
				RunE: func(cmd *cobra.Command, args []string) error {
					key, err := datastorecli.DecodeKey(args[0])
					if err != nil {
						return err
					}
					fmt.Fprint(os.Stdout, key.String())
					return nil
				},
			}
			return r
		})())

		return keyCommand
	})())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func validateFirstArgAsKind(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.Errorf("kind is required")
	}
	return nil
}

func newClient(projectID, namespace string) (*datastorecli.Client, error) {
	return datastorecli.NewClient(projectID, namespace), nil
}

func formatData(d interface{}) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(b); err != nil {
		return err
	}
	return nil
}

func formatArray(d *[]interface{}) error {
	for _, i := range *d {
		if err := formatData(i); err != nil {
			return err
		}
	}
	return nil
}

func formatStrings(d *[]string) error {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(*d, "\n"))
	return nil
}
