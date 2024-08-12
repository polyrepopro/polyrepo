package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	pushCommand.Flags().StringP("workspace", "w", "", "The name of the workspace to push changes for.")
	pushCommand.MarkFlagRequired("workspace")

	WorkspaceCommand.AddCommand(pushCommand)
}

var pushCommand = &cobra.Command{
	Use:   "push",
	Short: "Push changes for each repository in the workspace.",
	Long:  "Push changes for each repository in the workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		workspaceName, err := cmd.Flags().GetString("workspace")
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get name", map[string]interface{}{
				"error": err,
			})
		}

		setup, err := Setup("workspace.push", workspaceName)
		if err != nil {
			multilog.Fatal("workspace.push", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}

		errs := workspaces.Push(workspaces.PushArgs{
			Workspace: setup.Workspace,
		})
		if len(errs) > 0 {
			multilog.Fatal("workspace.push", "push failed", map[string]interface{}{
				"workspace": setup.Workspace.Name,
				"path":      setup.Workspace.Path,
				"errors":    errs,
			})
		}

		multilog.Info("workspace.push", "pushed changes for all repositories", map[string]interface{}{
			"workspace":    setup.Workspace.Name,
			"path":         setup.Workspace.Path,
			"repositories": len(*setup.Workspace.Repositories),
		})
	},
}
