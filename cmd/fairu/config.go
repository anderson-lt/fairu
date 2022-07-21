package main

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"

	"github.com/anderson-lt/fairu/config"
)

func editConfig() {
	path, err := config.GetConfigFile()
	if err != nil {
		log.Fatal("I can not get the path of the configuration file.")
	}

	log.Println("Configuration file path:", path)

	editor := os.Getenv("EDITOR")
	if editor == "" {
		log.Fatal("The environment variable EDITOR is empty or is not defined.")
	}

	// Execute text editor.
	cmd := exec.Command(editor, path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() *config.Config {
	conf, err := config.Load()
	var perr *fs.PathError
	switch {
	case errors.Is(err, fs.ErrNotExist) && errors.As(err, &perr):
		log.Fatal("Non-existent configuration file:\n", perr.Path)

	case err == io.EOF:
		path, err := config.GetConfigFile()
		if err != nil {
			panic("Unexpected error: " + err.Error())
		}
		log.Fatal("Empty configuration file:\n", path)

	case err != nil:
		log.Fatalln("Fatal error:", err)
	}

	return conf
}
