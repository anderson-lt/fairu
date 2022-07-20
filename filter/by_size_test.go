// Copyright (C) 2022, Anderson Lizarazo Tellez

package filter

import "testing"

func TestSize(t *testing.T) {
	// The file heavyfile weighs 32KB.
	if !Size("testdata/heavyfile", []string{"20KB"}) {
		t.Fatal("The file should not be filtered")
	}

	// The file lightfile weighs 4KB.
	if Size("testdata/lightfile", []string{"3MB"}) {
		t.Fatal("The file should be filtered")
	}
}

func TestConsumes(t *testing.T) {
	// The file mediumfile weighs 8KB.
	if !Consumes("testdata/mediumfile", []string{"2KB", "10KB"}) {
		t.Fatal("The file should not be filtered")
	}

	if Consumes("testdata/mediumfile", []string{"1B", "1KB"}) {
		t.Fatal("The file should be filtered")
	}
}

func TestParseSize(t *testing.T) {
	Sizes := map[string]int64{
		"1":    1,
		"3B":   3,
		"1KB":  1000,
		"1KiB": 1024,
		"5KB":  5000,
		"1MB":  1000000,
		"1MiB": 1048576,
		"5MiB": 5242880,
		"1GB":  1000000000,
		"1GiB": 1073741824,
		"2TB":  2000000000000,
		"5TiB": 5497558138880,
	}

	// Test valid values.
	for key, value := range Sizes {
		result, err := parseSize(key)
		if err != nil {
			t.Error("Unexpected error:", err)
			continue
		}

		if result != value {
			t.Log("Expression:", key)
			t.Log("Correct:", value)
			t.Error("Incorrect result:", result)
		}
	}

	// Test invalid values.
	_, err := parseSize("mb")
	if err == nil {
		t.Error(`parseSize("mb") no return error`)
	}

	_, err = parseSize("25invalid")
	if err == nil {
		t.Error(`parseSize("25invalid") no return error`)
	}
}
