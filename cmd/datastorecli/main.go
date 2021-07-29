package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
		validateArgs := func(cmd *cobra.Command, args []string) error {
			if err := validateFirstArgAsKind(cmd, args); err != nil {
				return err
			}
			if len(args) < 2 {
				return errors.Errorf("entity name is required")
			}
			return nil
		}

		r := &cobra.Command{
			Use:  "get KIND NAME",
			Args: validateArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				client, err := newClient(args)
				if err != nil {
					return err
				}
				if d, err := client.Get(context.Background(), args[1]); err != nil {
					return err
				} else {
					return formatData(d)
				}
			},
		}
		return r
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
