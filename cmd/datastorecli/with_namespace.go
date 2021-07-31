package main

import (
	"github.com/akm/datastorecli"
	"github.com/spf13/cobra"
)

type namespaceFunc func() datastorecli.Namespace

func withNamespace(fn func(namespaceFunc) *cobra.Command) func() *cobra.Command {
	return func() *cobra.Command {
		var namespace string
		r := fn(func() datastorecli.Namespace {
			return datastorecli.Namespace(namespace)
		})
		r.Flags().StringVar(&namespace, "namespace", "", "namespace")
		return r
	}
}
