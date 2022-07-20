// Copyright (C) 2022, Anderson Lizarazo Tellez

package config

import "github.com/anderson-lt/fairu/action"

// getAction contains all available actions. It returns nil if the action is
// non-existent.
func getAction(name string) action.Action {
	switch name {
	case "Print":
		return action.Print
	case "Write":
		return action.Write
	}

	return nil
}
