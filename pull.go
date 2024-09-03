package main

import (
	"fmt"
	"sync"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/repositories"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	pullCommand.Flags().StringSliceP("workspace", "w", []string{}, "the name of the workspace(s) to commit in")
	pullCommand.Flags().StringSliceP("tag", "t", []string{}, "the tags to filter the repositories by")
	root.AddCommand(pullCommand)
}

var pullCommand = &cobra.Command{
	Use:   "pull",
	Short: "pull the latest changes for each repository in the workspace",
	Long:  "pull the latest changes for each repository in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		setup, err := Setup("workspace.pull", "", util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.pull", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}
		workspaces, err := setup.Config.GetWorkspaces(util.GetArg[[]string](cmd, "workspace"))
		if err != nil {
			multilog.Fatal("workspace.pull", "failed to get workspaces", map[string]interface{}{
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
					err := repositories.Pull(repositories.PullArgs{
						Workspace:  setup.Workspace,
						Repository: &repo,
					})
					if err != nil {
						multilog.Error(workspace.Name, fmt.Sprintf("failed to pull in %s", repo.Name), map[string]interface{}{
							"workspace": setup.Workspace.Name,
							"path":      setup.Workspace.Path,
							"errors":    err,
						})
					} else {
						multilog.Info(workspace.Name, fmt.Sprintf("pulled in %s", repo.Name), map[string]interface{}{
							"workspace": setup.Workspace.Name,
							"path":      setup.Workspace.Path,
						})
					}
				}()
			}
		}
		wg.Wait()
	},
}
