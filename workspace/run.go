package workspace

import (
	"context"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/commands"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	runCommand.Flags().StringP("workspace", "w", "", "the name of the workspace to run the commands for")
	WorkspaceCommand.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "run the commands for each repository in the workspace",
	Long:  "run the commands for each repository in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := Setup("workspace.run", util.GetArg[string](cmd, "workspace"), util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.run", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}

		workspace, err := cfg.Config.GetWorkspace(util.GetArg[string](cmd, "workspace"))
		if err != nil {
			multilog.Fatal("workspace.run", "failed to get workspace", map[string]interface{}{
				"error": err,
			})
		}

		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		for _, repo := range *workspace.Repositories {
			if repo.Watches != nil {
				for _, watch := range *repo.Watches {
					go commands.Watch(ctx, workspace.Path, watch)
				}
			}
		}

		<-ctx.Done()
	},
}
