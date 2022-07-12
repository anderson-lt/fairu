package filter

import (
	"log"
	"path/filepath"
	"strings"
)

// Name filter files that do not have any name of the filter's arguments.
func Name(path string, args []string) bool {
	path = filepath.Base(path)
	for _, name := range args {
		if path == name {
			return true
		}
	}

	return false
}

// Glob Filter files that do not match the specified Unix Shell Glob.
func Glob(path string, args []string) bool {
	for _, glob := range args {
		if !strings.Contains("/", glob) {
			path = filepath.Base(path)
		}
		matched, err := filepath.Match(glob, path)
		if err != nil {
			log.Printf("Glob: invalid glob “%s”: %s", glob, err.Error())
		}
		if matched {
			return true
		}
	}

	return false
}

// Pattern filter files that do not match regular expression provided.
func Pattern(string, []string) bool {
	log.Println("Pattern: NOT IMPLEMENTED!")
	return false
}
