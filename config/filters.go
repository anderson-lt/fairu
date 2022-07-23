// Copyright (C) 2022, Anderson Lizarazo Tellez

package config

import (
	"strings"

	"github.com/anderson-lt/fairu/filter"
)

// getFilter contains all filters available. It returns nil if the filter is
// non-existent.
func getFilter(name string) filter.Filter {
	// Verify if it is a negated filter.
	var negated bool
	if strings.HasSuffix(name, "!") && len(name) > 1 {
		name = name[:len(name)-1]
		negated = true
	}

	filterFunc := getFilterName(name)
	if filterFunc == nil {
		return nil
	}

	// Denying the filter, if the negated version was requested.
	if negated {
		return filter.Negate(filterFunc)
	}

	return filterFunc
}

func getFilterName(name string) filter.Filter {
	switch name {
	case "Name":
		return filter.Name
	case "Glob":
		return filter.Glob
	case "Pattern":
		return filter.Pattern
	case "Type":
		return filter.Type
	case "Identifier":
		return filter.Identifier
	case "Size":
		return filter.Size
	case "Consumes":
		return filter.Consumes
	}

	return nil
}
