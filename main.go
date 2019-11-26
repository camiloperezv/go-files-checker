package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

var allowedTypes = []string{"png", "doc", "docx", "jpeg", "jpg", "pdf"}
var allowedMimeTypes = []string{"application/vnd.openxmlformats-officedocument.wordprocessingml.document", "application/pdf", "image/jpeg", "image/png"}

func isAllowedByArray(allowedArray []string, fileType string) bool {
	for _, allowed := range allowedArray {
		if fileType == allowed {
			return true
		}
	}
	return false
}

func main() {
	var files []string
	err := filepath.Walk("files/", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	_, files = files[0], files[1:]
	for _, file := range files {
		fileBuffer, _ := ioutil.ReadFile(file)
		kind, _ := filetype.Match(fileBuffer)
		if kind != filetype.Unknown {
			if isAllowedByArray(allowedTypes, kind.Extension) {
				if isAllowedByArray(allowedMimeTypes, kind.MIME.Value) {
					fmt.Printf("We can trust this file: %s\n Extension: %s\n Mimetype: %s\n---------------------------\n", file, kind.Extension, kind.MIME.Value)
					continue
				}
			}
		}
		fmt.Printf("Sorry. Unable to upload: %s Unknown %s %s\n---------------------------\n", file, kind.Extension, kind.MIME.Value)
	}
}
