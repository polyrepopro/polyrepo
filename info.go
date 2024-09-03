package main

import (
	"fmt"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	infoCommand.Flags().StringSliceP("workspace", "w", []string{}, "the name of the workspace(s) to get info for")
	infoCommand.Flags().StringSliceP("tag", "t", []string{}, "the tags to filter the repositories by")
	root.AddCommand(infoCommand)
}

var infoCommand = &cobra.Command{
	Use:   "info",
	Short: "Get information about the current workspace.",
	Long:  "Get information about the current workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		setup, err := Setup("workspace.commit", "", util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.commit", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}

		workspaces, err := setup.Config.GetWorkspaces(util.GetArg[[]string](cmd, "workspace"))
		if err != nil {
			multilog.Fatal("workspace.commit", "failed to get workspaces", map[string]interface{}{
				"error": err,
			})
		}

		multilog.Info("config", "located", map[string]interface{}{
			"path": setup.Config.Path,
		})

		for _, workspace := range *workspaces {
			repos := workspace.GetRepositories(util.GetArg[[]string](cmd, "tag"))
			multilog.Info("workspace", workspace.Name, map[string]interface{}{
				"path":         workspace.Path,
				"repositories": len(*workspace.Repositories),
				"tags":         workspace.Tags,
			})

			for _, repo := range *repos {
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
