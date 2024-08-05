package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/config"
	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	switchCommand.Flags().StringP("name", "n", "", "The name of the workspace to update.")
	switchCommand.MarkFlagRequired("name")

	switchCommand.Flags().StringP("branch", "b", "", "The branch to switch to.")
	switchCommand.MarkFlagRequired("branch")

	WorkspaceCommand.AddCommand(switchCommand)
}

var switchCommand = &cobra.Command{
	Use:   "switch",
	Short: "Switch the branch for each repository in the workspace.",
	Long:  "Switch the branch for each repository in the workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get name", map[string]interface{}{
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

		workspace, err := cfg.GetWorkspace(name)
		if err != nil {
			multilog.Fatal("workspace.switch", "failed to get workspace", map[string]interface{}{
				"name":  name,
				"error": err,
			})
		}

		errs := workspaces.Switch(workspaces.SwitchArgs{
			Workspace: workspace,
			Branch:    branch,
		})
		if len(errs) > 0 {
			multilog.Fatal("workspace.switch", "switch failed", map[string]interface{}{
				"errors": errs,
			})
		}

		multilog.Info("workspace.switch", "switched", map[string]interface{}{
			"branch":       branch,
			"repositories": len(*workspace.Repositories),
		})
	},
}
