package main

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/cli/workspace"
	"github.com/spf13/cobra"
)

func main() {
	multilog.RegisterLogger(multilog.LogMethod("console"), multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
		Level:  multilog.DEBUG,
		Format: multilog.FormatText,
	}))

	var root = &cobra.Command{
		Use:   "poly",
		Short: "Polyrepo CLI tool",
		Long:  "Polyrepo CLI tool",
	}

	root.AddCommand(workspace.WorkspaceCommand)

	root.Execute()
}
