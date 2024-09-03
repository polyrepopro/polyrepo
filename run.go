package main

import (
	"context"
	"path/filepath"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/mateothegreat/go-util/files"
	"github.com/polyrepopro/api/commands"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	runCommand.Flags().StringP("workspace", "w", "", "the name of the workspace to run the commands for")
	root.AddCommand(runCommand)
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
			if repo.Runners != nil {
				for _, runner := range *repo.Runners {
					if runner.Watch {
						go commands.Watch(ctx, repo.Name, workspace.Path, runner)
					} else {
						for _, command := range runner.Commands {
							var base string
							if runner.Cwd != "" {
								base = files.ExpandPath(filepath.Join(workspace.Path, runner.Cwd))
							} else {
								base = files.ExpandPath(filepath.Join(workspace.Path, command.Cwd))
							}
							go commands.Run(ctx, repo.Name, command, base)
						}
					}
				}
			}
		}

		<-ctx.Done()
	},
}
