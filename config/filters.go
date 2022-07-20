// Copyright (C) 2022, Anderson Lizarazo Tellez

package config

import "github.com/anderson-lt/fairu/filter"

// getFilter contains all filters available. It returns nil if the filter is
// non-existent.
func getFilter(name string) filter.Filter {
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
