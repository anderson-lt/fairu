// Copyright (C) 2022, Anderson Lizarazo Tellez

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/anderson-lt/fairu/config"
)

func main() {
	// Parse arguments.
	showVersion := flag.Bool("version", false, "Print information about the version.")
	editConf := flag.Bool("edit-config", false, "Execute an text editor in the configuration file.")
	flag.Parse()

	// Configure logging.
	log.SetFlags(0)
	log.SetPrefix("Fairu: ")

	switch {
	case *showVersion:
		fmt.Println("Fairu 0.1.0 (Preview)")
		fmt.Println("Copyright (C) 2022, Anderson Lizarazo Tellez")
		return

	case *editConf:
		editConfig()
		return
	}

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
