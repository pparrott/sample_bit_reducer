package audio

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

var infoLog = log.New(os.Stdout, "INFO: ", log.LstdFlags)
var wg sync.WaitGroup

func DownsampleFiles(filesCh <-chan string, targetBitRate uint16, numWorkers int) error {
	errCh := make(chan error, numWorkers)

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range filesCh {
				data, err := os.ReadFile(file)
				if err != nil {
					errCh <- err
					return
				}
				if err := downsampleFromBytes(data, file, targetBitRate); err != nil {
					errCh <- err
					return
				}
				infoLog.Printf("Successfully downsampled: %s", file)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		return err
	}
	return nil
}

func downsampleFromBytes(data []byte, file_path string, targetBitRate uint16) error {
	decoder := wav.NewDecoder(bytes.NewReader(data))
	buf, err := decoder.FullPCMBuffer()

	if err != nil {
		return err
	}

	buf = convertToRequiredBitDepth(buf, decoder.BitDepth, targetBitRate)

	outFile, err := os.Create(file_path)
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
		buf.Data[i] = ditherSample(buf.Data[i], shift) >> shift
	}

	return buf
}

func ditherSample(sample int, bitShift uint) int {
	ditherAmplitude := 1 << bitShift
	ditherNoise := rand.Intn(ditherAmplitude) - (ditherAmplitude / 2)
	return sample + ditherNoise
}
