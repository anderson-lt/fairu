// Copyright (C) 2022, Anderson Lizarazo Tellez

package config

import (
	"testing"

	"github.com/anderson-lt/fairu/filter"
)

func TestRule(t *testing.T) {
	Data := Rule{
		Filters: []Filter{
			{
				Filter:    filter.Name,
				Arguments: []string{"Example"},
			},
		},
	}

	if !Data.Accept("/home/gopher/Example") {
		t.Error("the filter must be passed")
	}
}
