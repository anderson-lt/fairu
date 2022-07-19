package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/anderson-lt/fairu/config"
)

func main() {
	// Parse arguments.
	showVersion := flag.Bool("version", false, "Print information about the version.")
	flag.Parse()

	if *showVersion {
		fmt.Println("Fairu 0.1.0 (Preview)")
		fmt.Println("Copyright (C) 2022, Anderson Lizarazo Tellez")
		return
	}

	// Configure logging.
	log.SetFlags(0)
	log.SetPrefix("Fairu: ")

	// Execute rules.
	root, err := os.Getwd()
	if err != nil {
		log.Fatal("I can not get the path to the current work directory.")
	}
	applyRules(root, loadConfig())
}

func applyRules(root string, config *config.Config) {
	walkFunc := func(path string, _ fs.DirEntry, _ error) error {
		// Skip current working directory.
		if path == root {
			return nil
		}

		for _, rule := range config.Rules {
			if rule.Accept(path) {
				rule.Execute(path)
				break
			}
		}

		return nil
	}

	if err := filepath.WalkDir(root, walkFunc); err != nil {
		panic("Unexpected error: " + err.Error())
	}
}

func loadConfig() *config.Config {
	conf, err := config.Load()
	var perr *fs.PathError
	if errors.Is(err, fs.ErrNotExist) && errors.As(err, &perr) {
		log.Fatal("Non-existent configuration file:\n", perr.Path)
	} else if err == io.EOF {
		path, err := config.GetConfigFile()
		if err != nil {
			panic("Unexpected error: " + err.Error())
		}
		log.Fatal("Empty configuration file:\n", path)
	} else if err != nil {
		log.Fatalln("Fatal error:", err)
	}

	return conf
}
