package audio

import (
	"bytes"
	"log"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

var infoLog = log.New(os.Stdout, "INFO: ", log.LstdFlags)

func DownsampleFiles(files []AudioFile, targetBitRate uint16) error {
	for _, file := range files {
		if err := downsampleFromBytes(file, targetBitRate); err != nil {
			return err
		}
		infoLog.Printf("Successfully downsampled: %s", file.Path)
	}
	return nil
}

func downsampleFromBytes(file AudioFile, targetBitRate uint16) error {
	decoder := wav.NewDecoder(bytes.NewReader(file.Data))
	buf, err := decoder.FullPCMBuffer()

	if err != nil {
		return err
	}

	buf = convertToRequiredBitDepth(buf, decoder.BitDepth, targetBitRate)

	outFile, err := os.Create(file.Path)
	if err != nil {
		return err
	}
	defer outFile.Close()

	encoder := wav.NewEncoder(outFile, int(decoder.SampleRate), int(targetBitRate), int(decoder.NumChans), 1)
	if err := encoder.Write(buf); err != nil {
		return err
	}

	return encoder.Close()
}

func convertToRequiredBitDepth(buf *audio.IntBuffer, originalBitRate uint16, targetBitRate uint16) *audio.IntBuffer {
	shift := uint(originalBitRate - targetBitRate)

	for i := range buf.Data {
		buf.Data[i] = buf.Data[i] >> shift
	}

	return buf
}
