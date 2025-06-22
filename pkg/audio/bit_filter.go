package audio

import (
	"os"

	"github.com/go-audio/wav"
)

func FilterBitRate(inputPaths <-chan string, outputFiles chan<- string, targetBitRate uint16) error {

	for path := range inputPaths {
		file, err := os.Open(path)

		if err != nil {
			infoLog.Printf("Failed to open file: %s", path)
			continue
		}

		decoder := wav.NewDecoder(file)
		if !decoder.IsValidFile() {
			infoLog.Printf("Invalid WAV file skipped: %s", path)
			file.Close()
			continue
		}
		if decoder.BitDepth > targetBitRate {
			outputFiles <- path
		}

		file.Close()
	}
	return nil
}
