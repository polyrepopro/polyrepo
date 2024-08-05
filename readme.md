# Polyrepo CLI

Polyrepo is a tool for managing polyrepo workspaces like a boss.

Polyrepo CLI is a command-line interface for the Polyrepo project. It allows you to manage your polyrepo workspace and its repositories.

## Installation

```bash
go install github.com/polyrepo/cli
```

## Commands

| Command                                                 | Description                                                   |
| ------------------------------------------------------- | ------------------------------------------------------------- |
| [polyrepo help](#polyrepo-help)                         | Show the help for the polyrepo CLI.                           |
| [polyrepo version](#polyrepo-version)                   | Show the version of the polyrepo CLI.                         |
| [polyrepo workspace init](#polyrepo-workspace-init)     | Initialize a new polyrepo workspace.                          |
| [polyrepo workspace info](#polyrepo-workspace-info)     | Show the information of the polyrepo workspace.               |
| [polyrepo workspace sync](#polyrepo-workspace-sync)     | Sync the polyrepo workspace with the remote.                  |
| [polyrepo workspace status](#polyrepo-workspace-status) | Show the status of the polyrepo workspace.                    |
| [polyrepo repo add](#polyrepo-repo-add)                 | Add a repository to the polyrepo workspace.                   |
| [polyrepo repo remove](#polyrepo-repo-remove)           | Remove a repository from the polyrepo workspace.              |
| [polyrepo repo sync](#polyrepo-repo-sync)               | Sync a repo with the remote.                                  |
| [polyrepo repo track](#polyrepo-repo-track)             | Adds the current working directory to the polyrepo workspace. |

### `polyrepo help`

Show the help for the polyrepo CLI.

### `polyrepo version`

Show the version of the polyrepo CLI.

### `polyrepo workspace init`

Initialize a polyrepo workspace configuration.

> This will overwrite the `.polyrepo.yaml` file at the given path if it exists.

| Flag       | Default            | Required | Description                                       |
| ---------- | ------------------ | -------- | ------------------------------------------------- |
| -p, --path | `~/.polyrepo.yaml` | Yes      | The path to save the `.polyrepo.yaml` file.       |
| -u, --url  |                    | Yes      | A URL to download the `.polyrepo.yaml` file from. |

### `polyrepo workspace info`

Show the information of the polyrepo workspace.

### `polyrepo workspace sync`

Sync the polyrepo workspace with the remote.

### `polyrepo workspace status`

Show the status of the polyrepo workspace.

### `polyrepo repo add`

Add a repository to the polyrepo workspace.

### `polyrepo repo remove`

Remove a repository from the polyrepo workspace.

### `polyrepo repo sync`

Sync a repo with the remote.

### `polyrepo repo track`

Adds the current working directory to the polyrepo workspace.
