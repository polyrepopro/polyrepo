package workspace

import (
	"path/filepath"
	"sync"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/mateothegreat/go-util/files"
	"github.com/polyrepopro/api/repositories"
	"github.com/polyrepopro/polyrepo/util"
	"github.com/spf13/cobra"
)

type StatusResult struct {
	Dirty bool
}

func init() {
	statusCommand.Flags().StringP("workspace", "w", "", "the name of the workspace to get the status of")
	WorkspaceCommand.AddCommand(statusCommand)
}

var statusCommand = &cobra.Command{
	Use:   "status",
	Short: "get the status of each repository in workspace(s)",
	Long:  "get the status of each repository in workspace(s)",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := Setup("workspace.status", util.GetArg[string](cmd, "workspace"), util.GetArg[string](cmd, "config"))
		if err != nil {
			multilog.Fatal("workspace.status", "failed to setup", map[string]interface{}{
				"error": err,
			})
		}

		dirty := 0
		total := 0
		wg := sync.WaitGroup{}
		for _, w := range *cfg.Config.Workspaces {
			for _, repo := range *w.Repositories {
				wg.Add(1)
				go func() {
					defer wg.Done()
					status, err := repositories.Status(files.ExpandPath(filepath.Join(w.Path, repo.Path)))
					if err != nil {
						multilog.Fatal("workspace.status", "failed to get status", map[string]interface{}{
							"error": err,
						})
					}
					if status.Dirty {
						dirty++
						multilog.Info("workspace.status", repo.Name, map[string]interface{}{
							"path":    filepath.Join(w.Path, repo.Path),
							"name":    repo.Name,
							"dirty":   true,
							"message": "pending changes",
						})
					}
					total++
				}()
			}
		}
		wg.Wait()
		multilog.Info("workspace.status", "summary", map[string]interface{}{
			"workspaces":   len(*cfg.Config.Workspaces),
			"repositories": total,
			"dirty":        dirty,
		})
	},
}
