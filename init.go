package main

import (
	"log"

	"github.com/polyrepopro/api/workspaces"
	"github.com/spf13/cobra"
)

func init() {
	initCommand.Flags().StringP("path", "p", "", "The path to the workspace.")
	initCommand.MarkFlagRequired("path")

	initCommand.Flags().StringP("url", "u", "", "The URL to the workspace.")

	root.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new polyrepo config.",
	Long:  "Initialize a new polyrepo config.",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf("Failed to get path: %v", err)
		}

		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalf("Failed to get url: %v", err)
		}

		cfg, err := workspaces.Init(workspaces.InitArgs{
			Path: path,
			URL:  url,
		})
		if err != nil {
			log.Fatalf("Failed to init workspace: %v", err)
		}

		log.Printf("Initialized workspace: %v", cfg.Path)
	},
}
