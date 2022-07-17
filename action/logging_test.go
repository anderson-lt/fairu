package action

import (
	"testing"
)

func TestExpandEnviron(t *testing.T) {
	t.Setenv("HOME", "/home")

	Path := "/home/gopher/mydata/data.yaml"
	Data := map[string]string{
		"$Path":                             Path,
		"$BaseName":                         "data.yaml",
		"$Extension":                        "YAML",
		"$ShellPath":                        "~/gopher/mydata/data.yaml",
		"$ShortPath":                        "~/g/m/data.yaml",
		"File: '$BaseName' Ext: $Extension": "File: 'data.yaml' Ext: YAML",
		"$NoExistentEnviromentVariable":     "",
		"$HOME":                             "/home",
		"":                                  "",
		"Gopher":                            "Gopher",
	}

	for key, value := range Data {
		if expandEnviron(Path, key) != value {
			t.Log("Data:", key)
			t.Log("Correct:", value)
			t.Error("Incorrect result:", expandEnviron(Path, key))
		}
	}
}

func TestToShellPath(t *testing.T) {
	t.Setenv("HOME", "/home")

	Data := map[string]string{
		"/home":              "~",
		"/":                  "/",
		"/home/gopher":       "~/gopher",
		"/root":              "/root",
		"/gophers/home/blue": "/gophers/home/blue",
		"/home/home/home":    "~/home/home",
	}

	for key, value := range Data {
		if toShellPath(key) != value {
			t.Log("Data:", key)
			t.Log("Correct:", value)
			t.Error("Incorrect conversion:", toShellPath(key))
		}
	}
}

func TestToShortPath(t *testing.T) {
	t.Setenv("HOME", "/home")

	Data := map[string]string{
		"/home":                            "~",
		"/":                                "/",
		"/home/gopher":                     "~/gopher",
		"/root":                            "/root",
		"/gophers/home/blue":               "/g/h/blue",
		"/home/home/home":                  "~/h/home",
		"/it/is/a/very/long/path/for/test": "/i/i/a/v/l/p/f/test",
	}

	for key, value := range Data {
		if toShortPath(key) != value {
			t.Log("Data:", key)
			t.Log("Correct:", value)
			t.Error("Incorrect conversion:", toShortPath(key))
		}
	}
}
