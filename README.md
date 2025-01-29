# Go Recursive File Finder

This is a Go program that recursively searches for a file in a given directory and copies the found file path to the clipboard.

## Features
- Recursively searches through directories.
- Ignores `node_modules` and `vendor` directories.
- Uses Go routines and channels for concurrency.
- Copies the found file path to the clipboard.
- Gracefully stops when the file is found.

## Prerequisites
- Go 1.18 or later
- Install the required package:
  ```sh
  go get golang.design/x/clipboard
  ```

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/kevariable/go-recursive-file-finder.git
   cd go-recursive-file-finder
   ```
2. Build the program:
   ```sh
   go build -o filefinder main.go
   ```

## Usage
Run the program with the following command:
```sh
./filefinder -path "/your/search/path" -filename "target_file.txt"
```

### Flags:
- `-path` (default: `.`) - The starting directory for the search.
- `-filename` - The filename to search for.

## Example
```sh
./filefinder -path "/home/user/projects" -filename "config.json"
```
If `config.json` is found, its full path is printed and copied to the clipboard.

## How It Works
1. The program initializes the clipboard.
2. It starts searching from the provided directory (`-path`).
3. It uses goroutines and channels to efficiently traverse directories.
4. When the target filename is found, it:
    - Prints the full path.
    - Copies it to the clipboard.
    - Stops further execution.

## License
This project is open-source and available under the MIT License.

