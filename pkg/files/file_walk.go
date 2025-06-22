package files

import (
	"os"
	"path/filepath"
)

func GetWavFilePaths(root string, out chan<- string) error {

	defer close(out)

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".wav" {
			out <- path
		}
		return nil
	})
	return err
}
