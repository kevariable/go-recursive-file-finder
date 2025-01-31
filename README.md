# Go Recursive File Finder

This is a Go program that recursively searches for a file in a given directory and prints the found file path.

## Features
- Recursively searches through directories.
- Ignores `node_modules` and `vendor` directories by default.
- Uses Goroutines and channels for concurrency.
- Displays the time taken for the search.
- Gracefully stops when the file is found.

## Prerequisites
- Go 1.23.4 or later

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/kevariable/go-recursive-file-finder.git
   cd go-recursive-file-finder
   ```
2. Build the program:
   ```sh
   go build -o filefinder find.go
   ```

## Usage
Run the program with the following command:
```sh
./filefinder -path "/your/search/path" -filename "target_file.txt" -ignores "node_modules, vendor"
```

### Flags:
- `-path` (default: `.`) - The starting directory for the search.
- `-filename` - The filename to search for.
- `-ignores` (default: `node_modules, vendor`) - Comma-separated list of directories to ignore.

## Example
```sh
./filefinder -path "/home/user/projects" -filename "config.json"
```
If `config.json` is found, its full path is printed.

## How It Works
1. The program initializes and parses flags.
2. It executes the `ls` command to list directory contents.
3. It uses Goroutines and channels to efficiently traverse directories.
4. When the target filename is found, it:
   - Prints the full path.
   - Stops further execution.
5. The program displays the total execution time.

## License
This project is open-source and available under the MIT License.
