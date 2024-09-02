package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/config"
	"github.com/polyrepopro/api/workspaces"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	pullCommand.Flags().StringP("workspace", "w", "", "isolate to a specific workspace")
	WorkspaceCommand.AddCommand(pullCommand)
}

var pullCommand = &cobra.Command{
	Use:   "pull",
	Short: "pull the latest changes for each repository in the workspace",
	Long:  "pull the latest changes for each repository in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfig(util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.pull", "failed to get config", map[string]interface{}{
				"error": err,
			})
		}

		workspaceName := util.GetArg[string](cmd, "workspace")
		if workspaceName == "" {
			for _, workspace := range *cfg.Workspaces {
				workspaces.Pull(workspaces.PullArgs{
					Workspace: &workspace,
				})

				multilog.Info("workspace.pull", "pull completed", map[string]interface{}{
					"workspace":    workspace.Name,
					"path":         workspace.Path,
					"repositories": len(*workspace.Repositories),
				})
			}
		} else {
			workspace, err := cfg.GetWorkspace(workspaceName)
			if err != nil {
				multilog.Fatal("workspace.pull", "failed to get workspace", map[string]interface{}{
					"config":    cfg.Path,
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
		}
	},
}
