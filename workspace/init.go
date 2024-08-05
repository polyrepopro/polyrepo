package workspace

import (
	"log"

	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	initCommand.Flags().StringP("path", "p", "", "The path to the workspace.")
	initCommand.MarkFlagRequired("path")

	initCommand.Flags().StringP("url", "u", "", "The URL to the workspace.")

	WorkspaceCommand.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new workspace.",
	Long:  "Apply tags to subscriptions.",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf("Failed to get path: %v", err)
		}

		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalf("Failed to get url: %v", err)
		}

		err = workspaces.Init(workspaces.InitArgs{
			Path: path,
			URL:  url,
		})
		if err != nil {
			log.Fatalf("Failed to init workspace: %v", err)
		}
	},
}
