package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	cmds := []*cobra.Command{execute()}

	root := &cobra.Command{}

	for _, cmd := range cmds {
		root.AddCommand(cmd)
	}

	return root
}
