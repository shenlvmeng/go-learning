package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose messages")
var sema = make(chan struct{}, 20)

type SizeInfo struct {
	index int
	size  int64
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan SizeInfo)
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	nFiles := make([]int64, len(roots))
	nBytes := make([]int64, len(roots))
	var tasks sync.WaitGroup

	for i, root := range roots {
		go walkDir(root, fileSizes, &tasks, i)
	}
	go func() {
		tasks.Wait()
		close(fileSizes)
	}()
loop:
	for {
		select {
		case sizeInfo, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nFiles[fileSizes.index]++
			nBytes[fileSizes.index] += size
		case <-tick:
			printDiskUsage(roots, nFiles, nBytes)
		}
	}
	printDiskUsage(roots, nFiles, nBytes)
}

func walkDir(dir string, fileSizes chan<- SizeInfo, tasks *sync.WaitGroup, index int) {
	defer tasks.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			tasks.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSizes)
		} else {
			fileSizes <- SizeInfo{index, entry.Size()}
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(roots []string, files, bytes []int64) {
	for i, r := range roots {
		fmt.Printf("%s: %d files %.1f GB\n", r, files[i], float64(nBytes[i])/1e9)
	}
}
