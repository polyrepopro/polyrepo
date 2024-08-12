package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	commitCommand.Flags().StringP("workspace", "w", "", "The name of the workspace to update.")
	commitCommand.MarkFlagRequired("workspace")

	WorkspaceCommand.AddCommand(commitCommand)
}

var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit the changes for each repository in the workspace.",
	Long:  "Commit the changes for each repository in the workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		workspaceName, err := cmd.Flags().GetString("workspace")
		if err != nil {
			multilog.Fatal("workspace.commit", "failed to get workspace", map[string]interface{}{
				"error": err,
			})
		}

		setup, err := Setup("workspace.commit", workspaceName)
		if err != nil {
			multilog.Fatal("workspace.commit", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}

		msgs, errs := workspaces.Commit(workspaces.CommitArgs{
			Workspace: setup.Workspace,
		})
		if len(errs) > 0 {
			multilog.Fatal("workspace.commit", "commit failed", map[string]interface{}{
				"workspace": setup.Workspace.Name,
				"path":      setup.Workspace.Path,
				"errors":    errs,
			})
		}

		multilog.Info("workspace.commit", "committed", map[string]interface{}{
			"workspace":    setup.Workspace.Name,
			"path":         setup.Workspace.Path,
			"repositories": len(*setup.Workspace.Repositories),
			"messages":     msgs,
		})
	},
}
