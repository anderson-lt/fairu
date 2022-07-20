// Copyright (C) 2022, Anderson Lizarazo Tellez

package filter

import "log"

// Type filter files that are not of the specified type.
func Type(string, []string) bool {
	log.Println("Type: NOT IMPLEMENTED!")
	return false
}

// Identifier filter files that do not start with the magic number provided.
func Identifier(path string, args []string) bool {
	log.Println("Identifier: NOT IMPLEMENTED!")
	return false
}
