package workspace

import "github.com/spf13/cobra"

var WorkspaceCommand = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces.",
	Long:  "Manage workspaces.",
	Args:  cobra.ExactArgs(1),
}
