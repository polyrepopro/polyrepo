package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/commands"
	"github.com/polyrepopro/api/config"
	"github.com/spf13/cobra"
)

func init() {
	runCommand.Flags().StringP("workspace", "w", "", "The name of the workspace to update.")
	WorkspaceCommand.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run the commands for each repository in the workspace.",
	Long:  "Run the commands for each repository in the workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.GetRelativeConfig()
		if err != nil {
			multilog.Fatal("workspace.run", "failed to get config", map[string]interface{}{
				"error": err,
			})
		}

		workspace, err := config.GetWorkspaceByWorkingDir()
		if err != nil {
			multilog.Fatal("workspace.run", "failed to get workspace", map[string]interface{}{
				"error": err,
			})
		}

		for _, repo := range *workspace.Repositories {
			if repo.Watches != nil {
				for _, watch := range *repo.Watches {
					commands.Watch(watch)
				}
			}
		}
	},
}
