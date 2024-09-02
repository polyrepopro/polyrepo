package main

import (
	"fmt"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/config"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	infoCommand.Flags().StringP("path", "p", "", "The path to the workspace.")
	root.AddCommand(infoCommand)
}

var infoCommand = &cobra.Command{
	Use:   "info",
	Short: "Get information about the current workspace.",
	Long:  "Get information about the current workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfig(util.GetArg[string](cmd, "path"))
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get config", map[string]interface{}{
				"error": err,
			})
		}

		multilog.Info("config", "located", map[string]interface{}{
			"path": cfg.Path,
		})

		for _, workspace := range *cfg.Workspaces {
			multilog.Info("workspace", "workspace", map[string]interface{}{
				"name": workspace.Name,
				"path": workspace.Path,
			})

			for _, repo := range *workspace.Repositories {
				multilog.Info("workspace.repositories", fmt.Sprintf("%s:%s", workspace.Name, repo.Name), map[string]interface{}{
					"url":    repo.URL,
					"branch": repo.Branch,
					"path":   repo.Path,
					"origin": repo.Origin,
				})
			}
		}
	},
}
