package main

import (
	"os"
	"path/filepath"
	"strings"
)

func getMdFiles(dirPath, ext string) ([]string, error) {
	var fileList []string

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func getLinesFromFile(fileName string) []string {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	lines := strings.Split(contents, "\n")
	return lines
}
