# Polyrepo CLI

## Commands

| Command                                                 | Description                                     |
| ------------------------------------------------------- | ----------------------------------------------- |
| [polyrepo workspace init](#polyrepo-workspace-init)     | Initialize a new polyrepo workspace.            |
| [polyrepo workspace info](#polyrepo-workspace-info)     | Show the information of the polyrepo workspace. |
| [polyrepo workspace sync](#polyrepo-workspace-sync)     | Sync the polyrepo workspace with the remote.    |
| [polyrepo workspace status](#polyrepo-workspace-status) | Show the status of the polyrepo workspace.      |

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
