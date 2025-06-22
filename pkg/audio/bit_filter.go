package audio

import (
	"os"

	"github.com/go-audio/wav"
)

type AudioFile struct {
	Path string
	Data []byte
}

func FilterBitRate(inputPaths <-chan string, outputFiles chan<- AudioFile, targetBitRate uint16) error {
	defer close(outputFiles)

	for path := range inputPaths {
		file, err := os.Open(path)

		if err != nil {
		}

		decoder := wav.NewDecoder(file)
		if !decoder.IsValidFile() {
			infoLog.Printf("Invalid WAV file skipped: %s", path)
			file.Close()
			continue
		}
		if decoder.BitDepth > targetBitRate {
			data, err := os.ReadFile(path)
			if err != nil {
				infoLog.Printf("Error reading file %s: %v", path, err)
				file.Close()
				continue
			}
			outputFiles <- AudioFile{
				Path: path,
				Data: data,
			}
		}

		file.Close()
	}
	return nil
}
