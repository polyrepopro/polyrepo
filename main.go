package main

import (
	"github.com/polyrepopro/cli/workspace"
	"github.com/spf13/cobra"
)

func main() {
	var root = &cobra.Command{
		Use:   "poly",
		Short: "Polyrepo CLI tool",
		Long:  "Polyrepo CLI tool",
	}

	root.AddCommand(workspace.WorkspaceCommand)

	root.Execute()
}
