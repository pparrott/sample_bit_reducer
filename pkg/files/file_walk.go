package files

import (
	"os"
	"path/filepath"
)

func GetWavFilePaths(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".wav" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
