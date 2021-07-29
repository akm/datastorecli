package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/akm/datastorecli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	projectID string
	namespace string
)

func main() {
	rootCmd := &cobra.Command{
		Use: "datastorecli",
	}

	rootCmd.PersistentFlags().StringVar(&projectID, "project-id", "", "GCP Project ID")
	rootCmd.PersistentFlags().StringVar(&projectID, "namespace", "", "namespace")

	rootCmd.AddCommand((func() *cobra.Command {
		var offset int
		var limit int
		var keysOnly bool
		r := &cobra.Command{
			Use:  "query KIND",
			Args: validateFirstArgAsKind,
			RunE: func(cmd *cobra.Command, args []string) error {
				client, err := newClient(args)
				if err != nil {
					return err
				}
				if keysOnly {
					if d, err := client.QueryKeys(context.Background(), offset, limit); err != nil {
						return err
					} else {
						return formatStrings(d)
					}
				} else {
					if d, err := client.QueryData(context.Background(), offset, limit); err != nil {
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
	})())

	rootCmd.AddCommand((func() *cobra.Command {
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
					if key, err = datastore.DecodeKey(args[0]); err != nil {
						return errors.Wrapf(err, "Failed to decode %s", args[0])
					}
				} else {
					kind := args[0]

					var parentKey *datastore.Key
					if encodedParent != "" {
						var err error
						if parentKey, err = datastore.DecodeKey(encodedParent); err != nil {
							return errors.Wrapf(err, "Failed to decode %s", encodedParent)
						}
					} else {
						parentKey = nil
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

				client, err := newClient(args)
				if err != nil {
					return err
				}
				if d, err := client.Get(context.Background(), key); err != nil {
					return err
				} else {
					return formatData(d)
				}
			},
		}
		r.Flags().StringVar(&encodedParent, "encoded-parent", "", "Encoded parent key")
		return r
	})())

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

					var parentKey *datastore.Key
					if encodedParent != "" {
						var err error
						if parentKey, err = datastore.DecodeKey(encodedParent); err != nil {
							return errors.Wrapf(err, "Failed to decode %s", encodedParent)
						}
					} else {
						parentKey = nil
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
					encoded := args[0]
					key, err := datastore.DecodeKey(encoded)
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

func newClient(args []string) (*datastorecli.Client, error) {
	kind := args[0]
	return datastorecli.NewClient(projectID, namespace, kind), nil
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
