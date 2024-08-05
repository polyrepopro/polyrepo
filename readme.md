# Polyrepo CLI

Polyrepo is a tool for managing polyrepo workspaces like a boss.

Polyrepo CLI is a command-line interface for the Polyrepo project. It allows you to manage your polyrepo workspace and its repositories.

## Installation

```bash
go install github.com/polyrepo/cli
```

## Configuration

Polyrepo workspaces are configured in a `.polyrepo.yaml` file which can be created by running `polyrepo workspace init`.

An example `.polyrepo.yaml` file looks like this:

```yaml
workspaces:
  - name: dev
    path: ~/workspace/polyrepo-dev
    repositories:
      - url: git@github.com:polyrepopro/api.git
        branch: main
        path: pkg/api
      - url: git@github.com:polyrepopro/cli.git
        branch: main
        path: pkg/cli
```

## Commands

| Command                                                 | Description                                                   |
| ------------------------------------------------------- | ------------------------------------------------------------- |
| [polyrepo help](#polyrepo-help)                         | Show the help for the polyrepo CLI.                           |
| [polyrepo version](#polyrepo-version)                   | Show the version of the polyrepo CLI.                         |
| [polyrepo init](#polyrepo-init)                         | Initialize a new global `.polyrepo.yaml` configuration file.  |
| [polyrepo workspace status](#polyrepo-workspace-status) | Show the status of the polyrepo workspace.                    |
| [polyrepo workspace sync](#polyrepo-workspace-sync)     | Sync workspace with the remotes.                              |
| [polyrepo workspace switch](#polyrepo-workspace-switch) | Switch the branch of repositories in a workspace.             |
| [polyrepo repo add](#polyrepo-repo-add)                 | Add a repository to the polyrepo workspace.                   |
| [polyrepo repo remove](#polyrepo-repo-remove)           | Remove a repository from the polyrepo workspace.              |
| [polyrepo repo sync](#polyrepo-repo-sync)               | Sync a repo with the remote.                                  |
| [polyrepo repo track](#polyrepo-repo-track)             | Adds the current working directory to the polyrepo workspace. |

### `polyrepo help`

Show the help for the polyrepo CLI.

### `polyrepo version`

Show the version of the polyrepo CLI.

### `polyrepo init`

Initialize a new global `.polyrepo.yaml` configuration file.

> This will overwrite the `.polyrepo.yaml` file at the given path if it exists.

| Flag       | Default            | Required | Description                                       |
| ---------- | ------------------ | -------- | ------------------------------------------------- |
| -p, --path | `~/.polyrepo.yaml` | **Yes**  | The path to save the `.polyrepo.yaml` file.       |
| -u, --url  |                    | No       | A URL to download the `.polyrepo.yaml` file from. |

### `polyrepo workspace status`

Show the status of the polyrepo workspace.

### `polyrepo workspace sync`

Sync workspace with the remotes.

This command syncs the workspace by ensuring that each repository exists locally.

| Flag       | Default | Required | Description                        |
| ---------- | ------- | -------- | ---------------------------------- |
| -n, --name |         | **Yes**  | The name of the workspace to sync. |

### `polyrepo workspace switch`

Switch the branch of repositories in a workspace.

| Flag         | Default | Required | Description                               |
| ------------ | ------- | -------- | ----------------------------------------- |
| -n, --name   |         | **Yes**  | The name of the workspace to switch.      |
| -b, --branch |         | **Yes**  | The branch to switch the repositories to. |

### `polyrepo repo add`

Add a repository to the polyrepo workspace.

### `polyrepo repo remove`

Remove a repository from the polyrepo workspace.

### `polyrepo repo sync`

Sync a repo with the remote.

### `polyrepo repo track`

Adds the current working directory to the polyrepo workspace.
