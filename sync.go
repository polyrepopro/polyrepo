package main

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	syncCommand.Flags().StringP("workspace", "w", "", "the name of the workspace to sync")
	root.AddCommand(syncCommand)
}

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync a workspace by syncing all repositories",
	Long:  "Sync a workspace by syncing all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		workspaceName, err := cmd.Flags().GetString("workspace")
		if err != nil {
			multilog.Fatal("workspace.sync", "sync failed", map[string]interface{}{
				"error": err,
			})
		}

		msgs, err := workspaces.Sync(workspaces.SyncArgs{
			Name: workspaceName,
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
