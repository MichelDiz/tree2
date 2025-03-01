# Simple Tree CLI

Tree2 CLI is a command-line tool that displays the structure of directories and files in a tree format. It allows filtering by file extension, exporting the tree to a file, and optionally displaying additional file details such as permissions, modification time, and size.

## Project Intent
The idea behind this little project was to create a customizable tool that replicates the functionality of the `tree` command while allowing more control over its features. Additionally, I wanted a solution that could be easily built for Windows whenever needed.

In fact there is nothing especial about this. Just for own use.


## Features
- Displays directory and file structure in a tree format.
- Optionally shows file permissions, modification time, and size.
- Supports filtering by file extension.
- Outputs to a specified file if needed.
- Allows ignoring specific directories or files.
- Supports JSON export.

## Installation
### Using Makefile
To build and install the program, use the following commands:
```sh
make build    # Compile the binary
make install  # Install the binary globally (requires sudo)
```

Alternatively, you can compile manually using:
```sh
go build -o tree2 main.go
```

## Usage
Run the command to display the directory tree:
```sh
tree2
```

### Available Flags
| Flag | Description |
|------|-------------|
| `-d, --directory` | Specify the root directory (default: `.`) |
| `-e, --extension` | Filter files by extension |
| `-o, --output` | Save the output to a file |
| `-l, --depth` | Set the depth limit for display |
| `-i, --ignore` | Ignore specific files/directories |
| `-j, --json` | Export the tree structure in JSON format |
| `--permissions` | Show file permissions |
| `--modtime` | Show file modification time |
| `--size` | Show file size |

### Example Usage
Show the directory tree with all details:
```sh
tree2 --permissions --modtime --size
```

Filter files by `.go` extension:
```sh
tree2 -e .go
```

Save the output to `tree.txt`:
```sh
tree2 -o tree.txt
```

Ignore `node_modules` and `vendor` directories:
```sh
tree2 -i node_modules,vendor
```

## Uninstallation
To remove the installed binary:
```sh
sudo rm /usr/local/bin/tree2
```

## License
This project is open-source and available under the MIT License.

