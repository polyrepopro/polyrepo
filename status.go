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
				// Use enhanced status with remote information
				var status repositories.StatusResult
				remoteName := repo.Origin
				if remoteName == "" {
					// Fall back to detecting default remote
					remoteName = "origin" // Default assumption
				}

				// Try enhanced status first, fall back to basic status if it fails
				status = repositories.StatusWithRemote(files.ExpandPath(filepath.Join(w.Path, repo.Path)), remoteName)
				if status.Code == repositories.StatusError {
					// Fall back to basic status if remote status fails
					status = repositories.Status(files.ExpandPath(filepath.Join(w.Path, repo.Path)))
				}

				// Count dirty repositories (including unpushed/unpulled)
				isDirty := status.Code == repositories.StatusDirty ||
					status.Code == repositories.StatusUnpushed ||
					status.Code == repositories.StatusUnpulled

				if isDirty {
					totalDirty++
					if all || dirty {
						var message string
						var coloredMessage string

						if verbose {
							message = status.Message
						} else {
							switch status.Code {
							case repositories.StatusDirty:
								message = "pending changes"
							case repositories.StatusUnpushed:
								message = fmt.Sprintf("↑%d", status.AheadCount)
							case repositories.StatusUnpulled:
								message = fmt.Sprintf("↓%d", status.BehindCount)
							default:
								message = status.Message
							}
						}

						// Color coding based on status
						switch status.Code {
						case repositories.StatusDirty:
							coloredMessage = color.RedString(message)
						case repositories.StatusUnpushed:
							coloredMessage = color.YellowString(message)
						case repositories.StatusUnpulled:
							coloredMessage = color.BlueString(message)
						default:
							coloredMessage = color.RedString(message)
						}

						logData := map[string]interface{}{
							"path":    filepath.Join(w.Path, repo.Path),
							"name":    repo.Name,
							"dirty":   true,
							"message": coloredMessage,
						}

						// Add remote status details if available
						if status.NeedsPush {
							logData["needs_push"] = true
							logData["ahead"] = status.AheadCount
						}
						if status.NeedsPull {
							logData["needs_pull"] = true
							logData["behind"] = status.BehindCount
						}

						multilog.Info(w.Name, repo.Name, logData)
					}
				} else {
					if all || clean {
						message := "clean"
						if status.Code == repositories.StatusClean {
							message = "clean and up to date"
						}
						if verbose {
							message = status.Message
						}

						multilog.Info(w.Name, repo.Name, map[string]interface{}{
							"path":    filepath.Join(w.Path, repo.Path),
							"name":    repo.Name,
							"dirty":   false,
							"message": color.GreenString(message),
						})
					}
				}
				total++
			}

			var message string
			cleanCount := total - totalDirty

			switch {
			case totalDirty == 0:
				message = fmt.Sprintf("all %d repositories in %s are clean and up to date", total, w.Name)
			case totalDirty == total:
				message = fmt.Sprintf("all %d repositories in %s need attention", total, w.Name)
			default:
				message = fmt.Sprintf("%d/%d repositories in %s need attention", totalDirty, total, w.Name)
			}

			multilog.Info(w.Name, message, map[string]interface{}{
				"workspace": w.Name,
				"clean":     color.GreenString(fmt.Sprintf("%d", cleanCount)),
				"dirty":     color.YellowString(fmt.Sprintf("%d", totalDirty)),
				"total":     total,
			})
		}
	},
}
