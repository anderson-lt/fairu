// Copyright (C) 2022, Anderson Lizarazo Tellez

package filter

import (
	"bytes"
	"log"
	"testing"
)

func init() {
	// Change the logger output to avoid staining the tests.
	log.SetOutput(new(bytes.Buffer))
}

func TestName(t *testing.T) {
	// Data for testing.
	filenames := [...]string{
		// Correct results (with Image.png).
		"Image.png",
		"../Image.png",
		"/home/user/Image.png",
		// Incorrect results.
		"MyImage.png",
		"Image1.png",
		"Image2.png",
	}

	// Apply operations.
	var results []string
	for _, fn := range filenames {
		if Name(fn, []string{"Image.png"}) {
			results = append(results, fn)
		}
	}

	// Verify results.
	// The length must be 3.
	if len(results) != 3 {
		t.Log("Results:", results)
		t.Fatal("The results length must be 3, but I got:", len(results))
	}

	// Verify contents.
	for i := 0; i < 3; i++ {
		if results[i] != filenames[i] {
			t.Logf("Correct (%d): %s", i, filenames[i])
			t.Logf("I got (Result %d): %s", i, results[i])
			t.Error("The result is incorrect.")
		}
	}

	// Test names as path.
	if Name("/gopher/blue/and/simple/", []string{"lue/and/simple"}) {
		t.Error("This should not accept the path.")
	}

	if !Name("gopher/blue/and/simple", []string{"and/simple"}) {
		t.Error("This should accept the path.")
	}
}

func TestNameMultiple(t *testing.T) {
	valid := map[string][]string{
		"go":           {"go", "go", "go"},
		"/home/gopher": {"house", "go", "gopher"},
		"../blue":      {"red", "yellow", "blue"},
	}

	invalid := map[string][]string{
		"other":  {"another", "file"},
		"/dir":   {"gopher", "blue"},
		"./code": {"red", "blue", "green"},
	}

	// Test valid.
	for path, names := range valid {
		if !Name(path, names) {
			t.Log("Path:", path)
			t.Log("Names:", names)
			t.Error("The path must be accepted")
		}
	}

	// Test invalid.
	for path, names := range invalid {
		if Name(path, names) {
			t.Log("Path:", path)
			t.Log("Names:", names)
			t.Error("The path must be not accepted")
		}
	}

	// Test invalid argument length.
	if Name("nil", []string{}) {
		t.Error("The filter must be not accept 0 arguments")
	}
}

func TestGlob(t *testing.T) {
	// Data for testing.
	filenames := [...]string{
		// Correct results (with *.png).
		"Photo.png",
		"Image.png",
		"Graphic.png",
		".private.png",
		// Incorrect results.
		"song.wav",
		"text.text",
	}

	// Apply operations.
	var results []string
	for _, fn := range filenames {
		if Glob(fn, []string{"*.png"}) {
			results = append(results, fn)
		}
	}

	// Verify results.
	// The length must be 3.
	if len(results) != 4 {
		t.Log("Results:", results)
		t.Fatal("The results length must be 3, but I got:", len(results))
	}

	// Verify contents.
	for i := 0; i < 4; i++ {
		if results[i] != filenames[i] {
			t.Logf("Correct (%d): %s", i, filenames[i])
			t.Logf("I got (Result %d): %s", i, results[i])
			t.Error("The result is incorrect.")
		}
	}
}

func TestGlobMultiple(t *testing.T) {
	if Glob("invalid-args", []string{"arg1", "arg2"}) {
		t.Error("Glob should not accept more than one argument ")
	}
}

func TestGlobHidden(t *testing.T) {
	hidden := [...]string{
		".hidden",
		"/path/any/no/visible/.file",
		"../hidden/.any",
		"../.gopher",
	}

	normal := [...]string{
		"visible",
		"/root/of/path",
		"../myfile",
	}

	// Test hidden files.
	t.Log("Using the glob .*")
	for _, file := range hidden {
		if !Glob(file, []string{".*"}) {
			t.Log("File:", file)
			t.Error("The file should pass the filter")
		}
	}

	// Test normal files.
	t.Log("Using the glob .*")
	for _, file := range normal {
		if Glob(file, []string{".*"}) {
			t.Log("File:", file)
			t.Error("The file should pass the filter")
		}
	}
}

func TestGlobInvalid(t *testing.T) {
	invalid := [...]string{
		"][",
		"[",
	}

	for _, glob := range invalid {
		if Glob("invalid-glob", []string{glob}) {
			t.Log("Glob used:", glob)
			t.Error("The glob should not have been accepted")
		}
	}
}

func TestPattern(t *testing.T) {
	// Data for testing.
	filenames := [...]string{
		// Correct results (with info-[a-z]{3}).
		"info-abc",
		"info-xyz",
		"info-dfg",
		// Incorrect results.
		"info-du",
		"info-213",
		"song.wav",
		"text.text",
	}

	// Apply operations.
	var results []string
	for _, fn := range filenames {
		if Pattern(fn, []string{"info-[a-z]{3}"}) {
			results = append(results, fn)
		}
	}

	// Verify results.
	// The length must be 3.
	if len(results) != 3 {
		t.Log("Results:", results)
		t.Fatal("The results length must be 3, but I got:", len(results))
	}

	// Verify contents.
	for i := 0; i < 3; i++ {
		if results[i] != filenames[i] {
			t.Logf("Correct (%d): %s", i, filenames[i])
			t.Logf("I got (Result %d): %s", i, results[i])
			t.Error("The result is incorrect.")
		}
	}
}

func TestPatternMultiple(t *testing.T) {
	if Pattern("invalid-arguments", []string{}) {
		t.Error("Pattern should accept an argument exactly.")
	}

	if Pattern("invalid-arguments", []string{"arg1", "arg2"}) {
		t.Error("Pattern should accept an argument exactly.")
	}
}

func TestPatternInvalid(t *testing.T) {
	if Pattern("invalid-pattern", []string{`][`}) {
		t.Error("Regular expression should not have been accepted.")
	}
}
