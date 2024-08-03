package workspace

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	initCommand.Flags().StringP("name", "n", "", "The name of the workspace.")
	initCommand.MarkFlagRequired("name")

	WorkspaceCommand.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new workspace.",
	Long:  "Apply tags to subscriptions.",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Failed to get name: %v", err)
		}

		log.Printf("Init workspace %s", name)
	},
}
