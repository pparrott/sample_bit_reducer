package audio

import (
	"bytes"
	"os"

	"github.com/go-audio/wav"
)

type AudioFile struct {
	Path string
	Data []byte
}

func FilterBitRate(paths []string, targetBitRate uint16) ([]AudioFile, error) {
	var filtered_files []AudioFile
	for _, path := range paths {
		data, err := os.ReadFile(path)

		if err != nil {
			return nil, err
		}

		decoder := wav.NewDecoder(bytes.NewReader(data))
		if !decoder.IsValidFile() {
			continue
		}
		if decoder.BitDepth > targetBitRate {
			filtered_files = append(filtered_files, AudioFile{
				Path: path,
				Data: data,
			})
		}
	}
	return filtered_files, nil
}
