package cli

import (
	"github.com/akm/datastorecli"
	"github.com/spf13/cobra"
)

type NamespaceFunc func() datastorecli.Namespace

func WithNamespace(fn func(NamespaceFunc) *cobra.Command) func() *cobra.Command {
	return func() *cobra.Command {
		var namespace string
		r := fn(func() datastorecli.Namespace {
			return datastorecli.Namespace(namespace)
		})
		r.Flags().StringVar(&namespace, "namespace", "", "namespace")
		return r
	}
}
