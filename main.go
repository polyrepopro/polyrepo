package main

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/cli/workspace"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "poly",
	Short: "Polyrepo CLI tool",
	Long:  "Polyrepo CLI tool",
}

func main() {
	multilog.RegisterLogger(multilog.LogMethod("console"), multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
		Level:  multilog.INFO,
		Format: multilog.FormatText,
	}))

	root.AddCommand(workspace.WorkspaceCommand)

	root.Execute()
}
