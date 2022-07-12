package filter

import "testing"

func TestName(t *testing.T) {
	// Data for testing.
	var filenames = [...]string{
		// Correct results (with Image.png).
		"Image.png",
		"../Image.png",
		"/home/user/Image.png",
		// Incorrect results.
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
}

func TestGlob(t *testing.T) {
	// Data for testing.
	var filenames = [...]string{
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

func TestPattern(t *testing.T) {
	// Data for testing.
	var filenames = [...]string{
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
