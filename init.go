package main

import (
	"log"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(initCommand)
	initCommand.Flags().StringP("path", "p", "", "the path to save the polyrepo config to")
	initCommand.MarkFlagRequired("path")

	initCommand.Flags().StringP("url", "u", "", "the URL to the polyrepo config to source from")
	initCommand.Flags().StringSliceP("tags", "t", []string{}, "use tags to filter repositories")
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new polyrepo config.",
	Long:  "Initialize a new polyrepo config.",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf("Failed to get path: %v", err)
		}

		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalf("Failed to get url: %v", err)
		}

		cfg, err := workspaces.Init(workspaces.InitArgs{
			Path: path,
			URL:  url,
		})
		if err != nil {
			log.Fatalf("Failed to init workspace: %v", err)
		}

		multilog.Info("init", "workspace initialized", map[string]interface{}{
			"path": cfg.Path,
		})
	},
}
