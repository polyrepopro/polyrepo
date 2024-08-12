package workspace

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/api/config"
)

type SetupResult struct {
	Config    *config.Config
	Workspace *config.Workspace
}

func Setup(group string, name string) (SetupResult, error) {
	cfg, err := config.GetRelativeConfig()
	if err != nil {
		multilog.Fatal(group, "failed to get config", map[string]interface{}{
			"error": err,
		})
	}

	workspace, err := cfg.GetWorkspace(name)
	if err != nil {
		multilog.Fatal(group, "failed to get workspace", map[string]interface{}{
			"name":  name,
			"error": err,
		})
	}

	return SetupResult{
		Config:    cfg,
		Workspace: workspace,
	}, nil
}
