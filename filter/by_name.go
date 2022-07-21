// Copyright (C) 2022, Anderson Lizarazo Tellez

package filter

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

// Name filter files that do not have any name of the filter's arguments.
// This requires at least an argument.
func Name(path string, args []string) bool {
	if len(args) < 1 {
		log.Print("Name: It requires at least an argument.")
		return false
	}

	path = filepath.Base(path)
	for _, name := range args {
		if path == name {
			return true
		}
	}

	return false
}

// Glob filter files that do not match the specified Unix Shell Glob.
// This filter takes exactly one argument.
func Glob(path string, args []string) bool {
	if len(args) != 1 {
		log.Print("Pattern: It takes exactly an argument.")
		return false
	}

	// If the glob contains an inclined bar, it means that you want to trap
	// directories; In the opposite case, you only want the name of the file,
	// so we only leave this to work.
	if !strings.Contains(args[0], "/") {
		path = filepath.Base(path)
	}

	matched, err := filepath.Match(args[0], path)
	if err != nil {
		log.Printf("Glob: Invalid Glob “%s”: %s", args[0], err.Error())
		return false
	}

	if matched {
		return true
	}

	return false
}

// Pattern filter files that do not match regular expression provided.
// This filter takes exactly one argument.
func Pattern(path string, args []string) bool {
	if len(args) != 1 {
		log.Print("Pattern: It takes exactly an argument.")
		return false
	}

	matched, err := regexp.MatchString(args[0], path)
	if err != nil {
		log.Println("Pattern: Invalid pattern:", err)
		return false
	}

	return matched
}
