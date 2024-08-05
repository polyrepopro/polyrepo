package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	syncCommand.Flags().StringP("name", "n", "", "The name of the workspace to update.")

	WorkspaceCommand.AddCommand(syncCommand)
}

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync a workspace.",
	Long:  "Sync a workspace by syncing all repositories.",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			multilog.Fatal("workspace.sync", "sync failed", map[string]interface{}{
				"error": err,
			})
		}

		msgs, err := workspaces.Sync(workspaces.SyncArgs{
			Name: name,
		})
		if err != nil {
			multilog.Fatal("workspace.sync", "sync failed", map[string]interface{}{
				"error": err,
			})
		}

		multilog.Info("workspace.sync", "synced", map[string]interface{}{
			"messages": msgs,
		})
	},
}
