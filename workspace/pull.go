package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/config"
	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	pullCommand.Flags().StringP("workspace", "w", "", "The name of the workspace to get the latest changes for.")
	pullCommand.MarkFlagRequired("workspace")

	WorkspaceCommand.AddCommand(pullCommand)
}

var pullCommand = &cobra.Command{
	Use:   "pull",
	Short: "Pull the latest changes for each repository in the workspace.",
	Long:  "Pull the latest changes for each repository in the workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		workspaceName, err := cmd.Flags().GetString("workspace")
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get name", map[string]interface{}{
				"error": err,
			})
		}

		cfg, err := config.GetRelativeConfig()
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get config", map[string]interface{}{
				"error": err,
			})
		}

		workspace, err := cfg.GetWorkspace(workspaceName)
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get workspace", map[string]interface{}{
				"workspace": workspace,
				"error":     err,
			})
		}

		errs := workspaces.Pull(workspaces.PullArgs{
			Workspace: workspace,
		})
		if len(errs) > 0 {
			multilog.Fatal("workspace.pull", "pull failed", map[string]interface{}{
				"workspace": workspace.Name,
				"path":      workspace.Path,
				"errors":    errs,
			})
		}

		multilog.Info("workspace.pull", "pulled", map[string]interface{}{
			"workspace":    workspace.Name,
			"path":         workspace.Path,
			"repositories": len(*workspace.Repositories),
		})
	},
}
