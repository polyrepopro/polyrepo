# Polyrepo CLI

## Commands

| Command                                                 | Description                                 |
| ------------------------------------------------------- | ------------------------------------------- |
| [polyrepo workspace init](#polyrepo-workspace-init)     | Initialize a new polyrepo workspace         |
| [polyrepo workspace sync](#polyrepo-workspace-sync)     | Sync the polyrepo workspace with the remote |
| [polyrepo workspace status](#polyrepo-workspace-status) | Show the status of the polyrepo workspace   |
| [polyrepo workspace list](#polyrepo-workspace-list)     | List the polyrepo workspaces                |
| [polyrepo workspace add](#polyrepo-workspace-add)       | Add a new polyrepo workspace                |
| [polyrepo workspace remove](#polyrepo-workspace-remove) | Remove a polyrepo workspace                 |
| [polyrepo workspace config](#polyrepo-workspace-config) | Configure the polyrepo workspace            |

### `polyrepo workspace init`

Initialize a new polyrepo workspace.

| Flag       | Default            | Required | Description                                       |
| ---------- | ------------------ | -------- | ------------------------------------------------- |
| -p, --path | `~/.polyrepo.yaml` | Yes      | The path to save the `.polyrepo.yaml` file.       |
| -u, --url  |                    | Yes      | A URL to download the `.polyrepo.yaml` file from. |
