package cli

import (
	"github.com/akm/datastorecli/models"
	"github.com/spf13/cobra"
)

type clientFunc func() *models.Client

func connectable(fn func(clientFn clientFunc) *cobra.Command) func() *cobra.Command {
	return func() *cobra.Command {
		var projectID string
		var namespace string
		r := fn(func() *models.Client {
			return models.NewClient(projectID, namespace)
		})
		r.Flags().StringVar(&projectID, "project-id", "", "GCP Project ID")
		r.Flags().StringVar(&namespace, "namespace", "", "namespace")
		return r
	}
}
