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
	pushCommand.Flags().StringSliceP("workspace", "w", []string{}, "the name of the workspace(s) to commit in")
	pushCommand.Flags().StringSliceP("tag", "t", []string{}, "the tags to filter the repositories by")
	root.AddCommand(pushCommand)
}

var pushCommand = &cobra.Command{
	Use:   "push",
	Short: "push changes for each repository in the workspace",
	Long:  "push changes for each repository in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		setup, err := Setup("workspace.push", "", util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.push", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}

		workspaces, err := setup.Config.GetWorkspaces(util.GetArg[[]string](cmd, "workspace"))
		if err != nil {
			multilog.Fatal("workspace.push", "failed to get workspaces", map[string]interface{}{
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
					errs := repositories.Push(repositories.PushArgs{
						Workspace:  setup.Workspace,
						Repository: &repo,
					})
					if err != nil {
						multilog.Error(workspace.Name, fmt.Sprintf("failed to push in %s", repo.Name), map[string]interface{}{
							"workspace": setup.Workspace.Name,
							"path":      setup.Workspace.Path,
							"errors":    errs,
						})
					} else {
						multilog.Info(workspace.Name, fmt.Sprintf("pushed in %s", repo.Name), map[string]interface{}{
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
