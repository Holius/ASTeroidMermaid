package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type GoMetadata struct {
	File      string   `json:"file"`
	Package   string   `json:"package"`
	Imports   []string `json:"imports"`
	Functions []string `json:"functions"`
}

func GoMetadataFromDirectory(directory string) ([]GoMetadata, error) {
	gmSlice := []GoMetadata{}

	filesProcessed := 0

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			if strings.HasSuffix(path, ".go") {
				gm, err := GoMetadataFromFile(path)
				if err != nil {
					return err
				}
				if gm == nil {
					return fmt.Errorf("recived null GoMetadata pointer when walking directory")
				}
				gmSlice = append(gmSlice, *gm)
				filesProcessed++
				if filesProcessed == 3 {
					//return fmt.Errorf("break")
				}

			}
		}
		return nil
	})

	return gmSlice, err
}
