package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/pparrott/sample_bit_reducer/pkg/audio"
	"github.com/pparrott/sample_bit_reducer/pkg/files"
)

const targetBitRate = 16

func main() {
	rootFolder := flag.String("path", "", "Path to root folder to search for wav files")
	flag.Parse()

	if *rootFolder == "" {
		log.Fatalf("Error: -path flag is required")
	}

	info, err := os.Stat(*rootFolder)
	if err != nil {
		log.Fatalf("Error accessing path %s: %v", *rootFolder, err)
	}
	if !info.IsDir() {
		log.Fatalf("Path %s is not a directory", *rootFolder)
	}

	pathsCh := make(chan string)
	filteredCh := make(chan string)

	go func() {
		if err := files.GetWavFilePaths(*rootFolder, pathsCh); err != nil {
			log.Fatal(err)
		}
		close(pathsCh)
	}()

	go func() {
		if err := audio.FilterBitRate(pathsCh, filteredCh, targetBitRate); err != nil {
			log.Fatal(err)
		}
		close(filteredCh)
	}()

	if err := audio.DownsampleFiles(filteredCh, targetBitRate, runtime.NumCPU()); err != nil {
		log.Fatal(err)
	}
}
