// Copyright (C) 2022, Anderson Lizarazo Tellez

package action

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Print print the arguments provided in the standard output and format them
// by separating those with a space and adding a line break at the end. You
// can embed environment variables using shell syntax.
func Print(path string, args []string) string {
	text := make([]any, len(args))
	for index, arg := range args {
		text[index] = expandEnviron(path, arg)
	}

	fmt.Println(text...)

	// Avoid modification of the route.
	return path
}

// Write prints the arguments provided in the standard output, this does not
// print spaces between the arguments, nor does it add a line break at the
// end. You can embed environment variables using shell syntax.
func Write(path string, args []string) string {
	text := make([]any, len(args))
	for index, arg := range args {
		text[index] = expandEnviron(path, arg)
	}

	fmt.Print(text...)

	// Avoid modification of the route.
	return path
}

// Report is similar to Print, but, write on the standard error output.
func Report(path string, args []string) string {
	text := make([]any, len(args))
	for index, arg := range args {
		text[index] = expandEnviron(path, arg)
	}

	fmt.Fprintln(os.Stderr, text...)

	// Avoid modification of the route.
	return path
}

// Error is similar to Write, but, write on the standard error output.
func Error(path string, args []string) string {
	text := make([]any, len(args))
	for index, arg := range args {
		text[index] = expandEnviron(path, arg)
	}

	fmt.Fprintln(os.Stderr, text...)

	// Avoid modification of the route.
	return path
}

func expandEnviron(path, text string) string {
	// Get absolute path.
	absPath, err := filepath.Abs(path)
	if err != nil {
		absPath = path

		log.Println("Logging: Expand: Error to get the absolute path:", err)
	}

	// Set dynamic variables.
	os.Setenv("Path", absPath)
	os.Setenv("BaseName", filepath.Base(path))
	os.Setenv("Extension", strings.ToUpper(strings.TrimPrefix(filepath.Ext(path), ".")))
	os.Setenv("ShellPath", toShellPath(path))
	os.Setenv("ShortPath", toShortPath(path))

	return os.ExpandEnv(text)
}

func toShellPath(path string) string {
	// Clean path and get absolute version.
	absPath, err := filepath.Abs(path)
	if err != nil {
		// Fallback absolute version (clean).
		absPath = filepath.Clean(path)
	}
	path = absPath

	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}
	home = filepath.Clean(home)

	if strings.HasPrefix(path, home) {
		// Replace first occurrence.
		path = strings.Replace(path, home, "~", 1)
	}

	return path
}

func toShortPath(path string) string {
	path = toShellPath(path)

	dir, base := filepath.Split(path)
	var newPath []string
	for _, d := range strings.Split(dir, "/") {
		if len(d) == 0 {
			continue
		}
		newPath = append(newPath, string(d[0]))
	}
	newPath = append(newPath, base)

	pth := filepath.Join(newPath...)
	if path[0] == '/' {
		return "/" + pth
	}

	return pth
}
