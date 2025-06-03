package main

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
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
	statusCommand.Flags().BoolP("all", "a", false, "show all repositories, even if they are clean")
	statusCommand.Flags().BoolP("clean", "", false, "show only clean repositories")
	statusCommand.Flags().BoolP("dirty", "", false, "show only dirty repositories")
}

var statusCommand = &cobra.Command{
	Use:   "status",
	Short: "get the status of each repository in workspace(s)",
	Long:  "get the status of each repository in workspace(s)",
	Run: func(cmd *cobra.Command, args []string) {
		all := util.GetArg[bool](cmd, "all")
		clean := util.GetArg[bool](cmd, "clean")
		dirty := util.GetArg[bool](cmd, "dirty")
		verbose := util.GetArg[bool](cmd, "verbose")
		cfg, err := Setup("workspace.status", util.GetArg[string](cmd, "workspace"), util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.status", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}

		multilog.Info("locating config(s)", "found", map[string]interface{}{
			"paths": []string{cfg.Config.Path},
		})

		for _, w := range *cfg.Config.Workspaces {
			totalDirty := 0
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
				status := repositories.Status(files.ExpandPath(filepath.Join(w.Path, repo.Path)))

				if status.Code == repositories.StatusDirty {
					totalDirty++
					if all || dirty {
						if verbose {
							multilog.Info(w.Name, repo.Name, map[string]interface{}{
								"path":    filepath.Join(w.Path, repo.Path),
								"name":    repo.Name,
								"dirty":   true,
								"message": color.RedString(status.Message),
							})
						} else {
							multilog.Info(w.Name, repo.Name, map[string]interface{}{
								"path":    filepath.Join(w.Path, repo.Path),
								"name":    repo.Name,
								"dirty":   true,
								"message": color.RedString("pending changes"),
							})
						}
					}
				} else {
					if all || clean {
						message := color.GreenString("clean")
						if verbose {
							message = color.GreenString(status.Message)
						}
						multilog.Info(w.Name, repo.Name, map[string]interface{}{
							"path":    filepath.Join(w.Path, repo.Path),
							"name":    repo.Name,
							"dirty":   false,
							"message": message,
						})
					}
				}
				total++
			}

			var message string
			if totalDirty == 0 {
				message = fmt.Sprintf("all %d repositories in %s are clean", total, w.Name)
			} else if totalDirty == total {
				message = fmt.Sprintf("all %d repositories in %s are dirty", total, w.Name)
			} else {
				message = fmt.Sprintf("%d/%d repositories in %s are dirty", totalDirty, total, w.Name)
			}
			multilog.Info(w.Name, message, map[string]interface{}{
				"workspace": w.Name,
				"clean":     color.GreenString(fmt.Sprintf("%d", total-totalDirty)),
				"dirty":     color.RedString(fmt.Sprintf("%d", totalDirty)),
			})
		}
	},
}
