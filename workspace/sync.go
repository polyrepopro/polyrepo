package workspace

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	syncCommand.Flags().StringP("name", "n", "", "The name of the workspace.")
	syncCommand.MarkFlagRequired("name")

	WorkspaceCommand.AddCommand(syncCommand)
}

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync the workspace.",
	Long:  "Sync the workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Failed to get name: %v", err)
		}

		log.Printf("Sync workspace %s", name)
	},
}
