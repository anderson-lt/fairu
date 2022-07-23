// Copyright (C) 2022, Anderson Lizarazo Tellez

package filter

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Type filter files that are not of the specified type.
// This accept MIME Types.
func Type(path string, args []string) bool {
	// Get content type of the file.
	file, err := os.Open(path)
	if err != nil {
		log.Println("Type: Error reading the file:", err)
	}
	defer file.Close()

	fileHeader := make([]byte, 512)
	_, err = io.ReadFull(file, fileHeader)
	if err != nil {
		log.Println("Type:", err)
	}

	contentType := http.DetectContentType(fileHeader)

	for _, arg := range args {
		if arg == contentType {
			return true
		}
	}

	return false
}

// Identifier filter files that do not start with the magic number provided.
func Identifier(path string, args []string) bool {
	log.Println("Identifier: NOT IMPLEMENTED!")
	return false
}
