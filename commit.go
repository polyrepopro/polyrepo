package main

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/repositories"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	commitCommand.Flags().StringSliceP("workspace", "w", []string{}, "the name of the workspace(s) to commit in")
	commitCommand.Flags().StringSliceP("tag", "t", []string{}, "the tags to filter the repositories by")
	commitCommand.Flags().StringP("message", "m", "", "the message to commit with")
	commitCommand.MarkFlagRequired("message")
	root.AddCommand(commitCommand)
}

var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit the changes for each repository in the workspace.",
	Long:  "Commit the changes for each repository in the workspace.",
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

		wg := sync.WaitGroup{}
		for _, workspace := range *workspaces {
			repos := workspace.GetRepositories(util.GetArg[[]string](cmd, "tag"))
			for _, repo := range *repos {
				wg.Add(1)
				go func() {
					defer wg.Done()
					result, errs := repositories.Commit(repositories.CommitArgs{
						Workspace:  setup.Workspace,
						Repository: &repo,
						Message:    util.GetArg[string](cmd, "message"),
					})
					if err != nil {
						multilog.Error(workspace.Name, fmt.Sprintf("failed to commit in %s", repo.Name), map[string]interface{}{
							"workspace": setup.Workspace.Name,
							"path":      setup.Workspace.Path,
							"errors":    errs,
						})
					}
					if result != nil && len(*result.Messages) > 0 {
						for _, msg := range *result.Messages {
							multilog.Info(workspace.Name, repo.Name, map[string]interface{}{
								"change": color.RedString(msg),
							})
						}
					} else {
						multilog.Info(workspace.Name, repo.Name, map[string]interface{}{
							"status": color.GreenString("clean"),
						})
					}
				}()
			}
		}
		wg.Wait()
	},
}
