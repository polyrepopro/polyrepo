package workspace

import (
	"log"

	"github.com/spf13/cobra"
)

var doctorCommand = &cobra.Command{
	Use:   "doctor",
	Short: "Check the status and repair the workspace.",
	Long:  "Check the status and repair the workspace.",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Failed to get name: %v", err)
		}

		log.Printf("Doctor workspace %s", name)
	},
}

func init() {
	doctorCommand.Flags().StringP("name", "n", "", "The name of the workspace.")
	doctorCommand.MarkFlagRequired("name")

	WorkspaceCommand.AddCommand(doctorCommand)
}
