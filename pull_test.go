package main

import (
	"fmt"
	"testing"
)

func TestPull(t *testing.T) {
	setup, err := Setup("workspace.pull", "", "~/workspace/cmskit/workspace/.polyrepo.yaml")
	if err != nil {
		t.Fatalf("failed to setup: %v", err)
	}

	workspaces, err := setup.Config.GetWorkspaces(nil)
	if err != nil {
		t.Fatalf("failed to get workspaces: %v", err)
	}

	repos := (*workspaces)[0].GetRepositories(nil)

	for _, repo := range *repos {
		fmt.Printf("repo: %+v\n", repo)
	}

}
