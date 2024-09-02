package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/config"
	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	switchCommand.Flags().StringP("branch", "b", "", "The branch to switch to.")
	switchCommand.MarkFlagRequired("branch")

	WorkspaceCommand.AddCommand(switchCommand)
}

var switchCommand = &cobra.Command{
	Use:   "switch",
	Short: "switch the branch for each repository in the workspace",
	Long:  "switch the branch for each repository in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		workspaceName, err := cmd.Flags().GetString("workspace")
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get workspace", map[string]interface{}{
				"error": err,
			})
		}

		branch, err := cmd.Flags().GetString("branch")
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get branch", map[string]interface{}{
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
			multilog.Fatal("workspace.switch", "failed to get workspace from config", map[string]interface{}{
				"workspace": workspaceName,
				"error":     err,
			})
		}

		errs := workspaces.Switch(workspaces.SwitchArgs{
			Workspace: workspace,
			Branch:    branch,
		})
		if len(errs) > 0 {
			multilog.Fatal("workspace.switch", "switch failed", map[string]interface{}{
				"workspace": workspace.Name,
				"path":      workspace.Path,
				"branch":    branch,
				"errors":    errs,
			})
		}

		multilog.Info("workspace.switch", "switched", map[string]interface{}{
			"branch":       branch,
			"repositories": len(*workspace.Repositories),
		})
	},
}
