package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/workspaces"
	"github.com/polyrepopro/cli/util"
	"github.com/spf13/cobra"
)

func init() {
	commitCommand.Flags().StringP("workspace", "w", "", "The name of the workspace to update.")
	commitCommand.Flags().StringP("message", "m", "", "The message to commit with.")
	WorkspaceCommand.AddCommand(commitCommand)
}

var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit the changes for each repository in the workspace.",
	Long:  "Commit the changes for each repository in the workspace.",
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

		for _, repo := range result {
			for _, msg := range *repo.Messages {
				multilog.Info("workspace.commit", "committed changes", map[string]interface{}{
					"change": msg,
					"repo":   repo.Path,
				})
			}
		}

		multilog.Info("workspace.commit", "committed all changes", map[string]interface{}{
			"workspace":    setup.Workspace.Name,
			"path":         setup.Workspace.Path,
			"repositories": len(*setup.Workspace.Repositories),
		})
	},
}
