package main

import (
	"github.com/akm/datastorecli"
	"github.com/spf13/cobra"
)

type clientFunc func() (*datastorecli.Client, error)

func connectableCommandFunc(fn func(clientFn clientFunc) *cobra.Command) func() *cobra.Command {
	return func() *cobra.Command {
		var projectID string
		var namespace string
		r := fn(func() (*datastorecli.Client, error) {
			return newClient(projectID, namespace)
		})
		r.Flags().StringVar(&projectID, "project-id", "", "GCP Project ID")
		r.Flags().StringVar(&namespace, "namespace", "", "namespace")
		return r
	}
}