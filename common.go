package main

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/config"
)

type SetupResult struct {
	Config    *config.Config
	Workspace *config.Workspace
}

func Setup(group string, workspaceName string, configPath string) (SetupResult, error) {
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		multilog.Fatal(group, "failed to get config", map[string]interface{}{
			"error": err,
		})
	}

	if workspaceName == "" && cfg.Default != "" {
		workspaceName = cfg.Default
	} else {
		if len(*cfg.Workspaces) > 0 {
			workspaceName = (*cfg.Workspaces)[0].Name
		} else {
			multilog.Fatal(group, "no workspaces found in config", map[string]interface{}{
				"error": err,
			})
		}
	}

	workspace, err := cfg.GetWorkspace(workspaceName)
	if err != nil {
		multilog.Fatal(group, "failed to get workspace", map[string]interface{}{
			"name":  workspaceName,
			"error": err,
		})
	}

	return SetupResult{
		Config:    cfg,
		Workspace: workspace,
	}, nil
}
