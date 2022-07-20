// Copyright (C) 2022, Anderson Lizarazo Tellez

package config_test

import (
	"fmt"
	"os"

	"github.com/anderson-lt/fairu/config"
)

func ExampleReadConfig() {
	file, err := os.Open("testdata/example_config.yaml")
	if err != nil {
		panic(err)
	}

	config, err := config.ReadConfig(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
	// Output: <Config: 3 Rules>
}
