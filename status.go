package main

import (
	"fmt"
	"path/filepath"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/mateothegreat/go-util/files"
	"github.com/polyrepopro/api/repositories"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

type StatusResult struct {
	Dirty bool
}

func init() {
	statusCommand.Flags().StringP("workspace", "w", "", "the name of the workspace to get the status of")
	root.AddCommand(statusCommand)
	statusCommand.Flags().StringSliceP("tags", "t", []string{}, "the tags to filter repositories by")
}

var statusCommand = &cobra.Command{
	Use:   "status",
	Short: "get the status of each repository in workspace(s)",
	Long:  "get the status of each repository in workspace(s)",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := Setup("workspace.status", util.GetArg[string](cmd, "workspace"), util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.status", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}
		for _, w := range *cfg.Config.Workspaces {
			dirty := 0
			total := 0
			repos := w.GetRepositories(util.GetArg[[]string](cmd, "tags"))
			if len(*repos) == 0 {
				multilog.Info("status", "no repositories found", map[string]interface{}{
					"workspace": w.Name,
					"tags":      util.GetArg[[]string](cmd, "tags"),
				})
				continue
			}
			for _, repo := range *repos {
				status, err := repositories.Status(files.ExpandPath(filepath.Join(w.Path, repo.Path)))
				if err != nil {
					multilog.Fatal("status", "failed to get status", map[string]interface{}{
						"error": err,
					})
				}
				if status.Dirty {
					dirty++
					multilog.Info(w.Name, repo.Name, map[string]interface{}{
						"path":    filepath.Join(w.Path, repo.Path),
						"name":    repo.Name,
						"dirty":   true,
						"message": "pending changes",
					})
				}
				total++
			}
			var message string
			if dirty == 0 {
				message = fmt.Sprintf("all %d repositories in %s are clean", total, w.Name)
			} else if dirty == total {
				message = fmt.Sprintf("all %d repositories in %s are dirty", total, w.Name)
			} else {
				message = fmt.Sprintf("%d/%d repositories in %s are dirty", dirty, total, w.Name)
			}
			multilog.Info(w.Name, message, map[string]interface{}{
				"workspace":    w.Name,
				"repositories": total,
				"dirty":        dirty,
			})
		}
	},
}
