package main

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/workspaces"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

func init() {
	pushCommand.Flags().StringP("workspace", "w", "", "isolate to a specific workspace")
	root.AddCommand(pushCommand)
}

var pushCommand = &cobra.Command{
	Use:   "push",
	Short: "push changes for each repository in the workspace",
	Long:  "push changes for each repository in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		setup, err := Setup("workspace.push", util.GetArg[string](cmd, "workspace"), util.GetArg[string](cmd, "config"))
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