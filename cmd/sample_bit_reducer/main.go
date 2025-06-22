package main

import (
	"flag"
	"log"
	"os"

	"github.com/pparrott/sample_bit_reader/pkg/audio"
	"github.com/pparrott/sample_bit_reader/pkg/files"
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

	file_paths, err := files.GetWavFilePaths(*rootFolder)

	if err != nil {
		log.Fatal(err)
	}

	filtered_files, err := audio.FilterBitRate(file_paths, targetBitRate)

	if err != nil {
		log.Fatal(err)
	}

	if err := audio.DownsampleFiles(filtered_files, targetBitRate); err != nil {
		log.Fatal(err)
	}
}
