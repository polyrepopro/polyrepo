package workspace

import (
	"sync"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/workspaces"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	WorkspaceCommand.AddCommand(commitCommand)
}

var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "commit the changes for each repository in the workspace",
	Long:  "commit the changes for each repository in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		setup, err := Setup("workspace.commit", util.GetArg[string](cmd, "workspace"), util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.commit", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}

		result, errs := workspaces.Commit(workspaces.CommitArgs{
			Workspace: setup.Workspace,
			Message:   util.GetArg[string](cmd, "message"),
		})
		if len(errs) > 0 {
			multilog.Fatal("workspace.commit", "commit failed", map[string]interface{}{
				"workspace": setup.Workspace.Name,
				"path":      setup.Workspace.Path,
				"errors":    errs,
			})
		}

		var wg sync.WaitGroup
		for _, repo := range result {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, msg := range *repo.Messages {
					multilog.Info("workspace.commit", "committed changes", map[string]interface{}{
						"change": msg,
						"repo":   repo.Path,
					})
				}
			}()
		}
		wg.Wait()

		multilog.Info("workspace.commit", "committed all changes", map[string]interface{}{
			"workspace":    setup.Workspace.Name,
			"path":         setup.Workspace.Path,
			"repositories": len(*setup.Workspace.Repositories),
		})
	},
}
