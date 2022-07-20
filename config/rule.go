// Copyright (C) 2022, Anderson Lizarazo Tellez

package config

import (
	"github.com/anderson-lt/fairu/action"
	"github.com/anderson-lt/fairu/filter"
)

// Rule represents a rule in the configuration.
type Rule struct {
	Name    string
	Filters []Filter
	Actions []Action
}

// Accept apply all the filters with your arguments on the specified path.
func (r Rule) Accept(path string) bool {
	for _, f := range r.Filters {
		if !f.Call(path) {
			return false
		}
	}

	return true
}

// Execute executes every action on the specified path.
func (r Rule) Execute(path string) {
	for _, a := range r.Actions {
		path = a.Call(path)
	}
}

// Filter contains a filter and its arguments.
type Filter struct {
	Filter    filter.Filter
	Arguments []string
}

// Call calls the filter with its arguments on the specified path.
func (f Filter) Call(path string) bool {
	return f.Filter(path, f.Arguments)
}

// Action contains a action and its arguments.
type Action struct {
	Action    action.Action
	Arguments []string
}

// Call calls the action with its arguments on the specified path.
func (a Action) Call(path string) string {
	return a.Action(path, a.Arguments)
}
