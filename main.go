package main

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/polyrepo/workspace"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "polyrepo",
	Short: "Manage polyrepos and workspaces like a boss 🚀.",
	Long:  "Manage polyrepos and workspaces like a boss 🚀.",
}

func main() {
	root.AddCommand(workspace.WorkspaceCommand)
	root.PersistentFlags().StringP("config", "c", "", "the path to the polyrepo config file")
	root.PersistentFlags().BoolP("verbose", "v", false, "output detailed logs")

	logLevel := multilog.INFO

	if verbose, _ := root.PersistentFlags().GetBool("verbose"); verbose {
		logLevel = multilog.DEBUG
	}

	multilog.RegisterLogger(multilog.LogMethod("console"), multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
		Level:  logLevel,
		Format: multilog.FormatText,
	}))

	root.Execute()
}
