// Copyright (C) 2022, Anderson Lizarazo Tellez

package filter

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

// Name filter files that do not have any name of the filter's arguments.
// For a route to pass the filter, it must end at least one of the names
// provided, for example:
//  Name(path, []string{"Img/Go"})
// It will return true if `path` is one of:
//  - /home/gopher/Img/Go
//  - /Img/Go
//  - Img/Go
// But, will be false if it is one of:
//  - MyImg/Go
//  - Go
//  - Img
//
// This requires at least an argument.
func Name(path string, args []string) bool {
	if len(args) < 1 {
		log.Print("Name: It requires at least an argument.")
		return false
	}

	path = filepath.Clean(path)
VerifyPaths:
	for _, name := range args {
		// Separate each name from the route.
		path := strings.Split(path, "/")
		name := strings.Split(name, "/")

		// The number of names to compare must be the same.
		if len(path) > len(name) {
			// Take the last `len(name)` elements.
			path = path[len(path)-len(name):]
		}

		// Verify that each name matches.
		for i := 0; i < len(name); i++ {
			if path[i] != name[i] {
				// If it does not match, try with the following.
				continue VerifyPaths
			}
		}
		return true
	}

	return false
}

// Glob filter files that do not match the specified Unix Shell Glob.
// In the current implementation, the `*` character also coincides with `.`
// (dot). If the Glob provided does not contain the `/` character, only the
// last element of the route provided will be evaluated. In the opposite case,
// the entire route will be evaluated.
//
// This filter takes exactly one argument.
func Glob(path string, args []string) bool {
	if len(args) != 1 {
		log.Print("Glob: It takes exactly an argument.")
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
		log.Println("Pattern: Invalid Pattern:", err)
		return false
	}

	return matched
}
