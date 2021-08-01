package cli

import "github.com/spf13/cobra"

func GroupCommand(name string, cmds ...*cobra.Command) *cobra.Command {
	r := &cobra.Command{Use: name}
	r.AddCommand(cmds...)
	return r
}
