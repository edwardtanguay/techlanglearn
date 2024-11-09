package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

)

/*
Get all files of a certain kind as a slice of strings

mdPathAndFileNames, _ := getFilesFromDirectory("../../static/data", "md")

- use relative path

- extension is mandatory
*/
func getFilesFromDirectory(dirPath, ext string) ([]string, error) {
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

/*
Get all lines from a file as a slice of strings

lines := getLinesFromFile("../../notes.txt")

- use relative path
*/
func getLinesFromFile(fileName string) []string {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	lines := strings.Split(contents, "\n")
	return lines
}

/*
Output information in console in a uniform way

devlog("no files are locked")

devlog(fmt.Sprintf("There are %d flashcards.", len(flashcards)))
*/
func devlog(line string) {
	fmt.Printf(">>> %s\n", line)
}

/*
Write text to a file

writeTextFile("../../src/data/test.txt", "testcontent")

- use relative path
*/
func writeTextFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("File could not be created: %s\n", err.Error())
	}
	_, err2 := file.WriteString(content)
	if err2 != nil {
		fmt.Printf("Could not write to file: %s\n", err2.Error())
	}
}
