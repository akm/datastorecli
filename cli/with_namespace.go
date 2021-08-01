package cli

import (
	"github.com/akm/datastorecli/models"
	"github.com/spf13/cobra"
)

type namespaceFunc func() models.Namespace

func withNamespace(fn func(namespaceFunc) *cobra.Command) func() *cobra.Command {
	return func() *cobra.Command {
		var namespace string
		r := fn(func() models.Namespace {
			return models.Namespace(namespace)
		})
		r.Flags().StringVar(&namespace, "namespace", "", "namespace")
		return r
	}
}
