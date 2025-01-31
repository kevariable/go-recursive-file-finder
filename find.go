package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
	"sync"
	"time"
)

func main() {
	ch := make(chan string, 5)
	done := make(chan struct{})
	var wg sync.WaitGroup

	var path = flag.String("path", ".", "path of ls command")
	var search = flag.String("filename", "", "Finding filename path")
	var ignores = flag.String("ignores", "node_modules, vendor", "Ignore paths")

	flag.Parse()

	var ignorePaths = strings.Split(*ignores, ", ")

	var output1, _ = exec.Command(
		"ls",
		*path,
	).Output()

	var response = strings.Split(strings.TrimSpace(string(output1)), "\n")

	startTime := time.Now()

	// Goroutine to close channel after all workers are done
	go func() {
		wg.Wait()
		fmt.Println("All workers completed")
		close(ch)
	}()

	wg.Add(1)
	go recursiveDir(
		response,
		*path,
		ch,
		&wg,
		done, // Pass done channel to workers,
		ignorePaths,
	)

	//fmt.Printf("Command: %s\n", file)
	osName := runtime.GOOS

	switch osName {
	case "darwin":
		fmt.Printf("Command + ^C to stop searching.\n")
		break
	default:
		fmt.Printf("Ctrl + ^C to stop searching.\n")
		break
	}

	// Main loop for receiving files
	for file := range ch {
		if strings.Contains(file, *search) {
			fmt.Printf("Filename found: %s\n", file)
			// close(done)
			// break
		}
	}

	wg.Wait() // Wait for all goroutines to finish
	endTime := time.Since(startTime)
	fmt.Printf("time took is %.2f\n", endTime.Seconds())
}

func recursiveDir(files []string, path string, ch chan<- string, wg *sync.WaitGroup, done <-chan struct{}, ignores []string) {
	defer wg.Done()

	for _, fileOrDir := range files {
		// Check if we should stop
		select {
		case <-done:
			return
		default:
		}

		if slices.Contains(ignores, fileOrDir) {
			continue
		}

		fullPath := fmt.Sprintf("%s/%s", path, fileOrDir)

		ok, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		if ok.Mode().IsDir() {
			var output, _ = exec.Command(
				"ls",
				fullPath,
			).Output()

			var splitEachFileOrDirectory = strings.Split(strings.TrimSpace(string(output)), "\n")

			wg.Add(1)
			go recursiveDir(
				splitEachFileOrDirectory,
				fullPath,
				ch,
				wg,
				done,
				ignores,
			)
		} else {
			select {
			case <-done:
				return
			case ch <- fullPath:
			}
		}
	}
}
