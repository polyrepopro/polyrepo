package main

import (
	"os"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "polyrepo",
	Short: "Manage polyrepos and workspaces like a boss 🚀.",
	Long:  "Manage polyrepos and workspaces like a boss 🚀.",
}

func main() {
	root.PersistentFlags().StringP("config", "c", "", "the path to the polyrepo config file")
	root.PersistentFlags().StringP("workspace", "w", "", "the name of the workspace to use")
	root.PersistentFlags().BoolP("verbose", "v", false, "output detailed logs")
	root.ParseFlags(os.Args)

	logLevel := multilog.INFO

	verbose, _ := root.PersistentFlags().GetBool("verbose")

	if verbose {
		logLevel = multilog.TRACE
	}

	multilog.RegisterLogger(multilog.LogMethod("console"), multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
		Level:  logLevel,
		Format: multilog.FormatText,
	}))

	root.Execute()
}
