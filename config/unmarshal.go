package config

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

// The YAML package used has a very peculiar way to represent the maps. These
// are represented using a slice of pair length, where each key is followed by
// its value, therefore the key "A" is in position 0 and its value in position
// 1, the key "B" is in position 2 and its value in position 3, and so on.

// Rules represents a set of rules.
type Rules []Rule

// UnmarshalYAML implements yaml.Unmarshaler.
func (r *Rules) UnmarshalYAML(node *yaml.Node) error {
	// I need a map type.
	if node.Kind != yaml.MappingNode {
		return errors.New("invalid structure of the document")
	}

	rules := make([]Rule, 0, len(node.Content)/2)
	for key, value := 0, 1; key < len(node.Content); key, value = key+2, value+2 {
		var name string
		err := node.Content[key].Decode(&name)
		if err != nil {
			return errors.New("the names of the rules must be strings")
		}

		rule := Rule{Name: name}
		err = rule.UnmarshalYAML(node.Content[value])
		if err != nil {
			return err
		}

		rules = append(rules, rule)
	}

	*r = rules

	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (r *Rule) UnmarshalYAML(node *yaml.Node) error {
	// I need a map type.
	// We do not decode in the native map type, because it does not store the
	// order of the elements, therefore we have to do it ourselves to preserve
	// the order.
	if node.Kind != yaml.MappingNode {
		return errors.New("invalid structure of the document")
	}

	// Get filters and actions.
	filters, actions, err := decodeRule(node)
	if err != nil {
		return err
	}

	// Set filters and actions.
	r.Filters = filters
	r.Actions = actions

	return nil
}

// decodeRule gets the filters and actions of the rule.
func decodeRule(node *yaml.Node) ([]Filter, []Action, error) {
	var filters []Filter
	var actions []Action

	var decodedFilters bool
	for key, value := 0, 1; key < len(node.Content); key, value = key+2, value+2 {
		// Get name of filter or action
		var name string
		err := node.Content[key].Decode(&name)
		if err != nil {
			return nil, nil, errors.New("filter and actions names must be strings")
		}

		// Get arguments.
		args, err := getArguments(node.Content[value])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid format of arguments: %w", err)
		}

		// Get filter or action.
		// Try to find a filter.
		filter := getFilter(name)
		if filter != nil {
			// Verify that the filter is not out of place, that is, it appears
			// after it is defined after an action.
			if decodedFilters {
				return nil, nil, errors.New("filter is out of place: " + name)
			}

			// Add filter to the list.
			filters = append(filters, Filter{Filter: filter, Arguments: args})
			continue
		}

		// Try to find an action.
		action := getAction(name)
		if action != nil {
			// Report that now it should only be left to decode actions.
			decodedFilters = true

			// Add action to the list.
			actions = append(actions, Action{Action: action, Arguments: args})
			continue
		}

		// If we reach this point it means that the given name is not filter
		// or action.
		return nil, nil, errors.New("unknown name: " + name)
	}

	return filters, actions, nil
}

func getArguments(node *yaml.Node) ([]string, error) {
	if node.ShortTag() == "!!null" {
		return nil, nil
	}

	switch node.Kind {
	case yaml.ScalarNode:
		var arg string
		err := node.Decode(&arg)
		if err != nil {
			return nil, errors.New("the arguments must be strings")
		}
		return []string{arg}, nil
	case yaml.SequenceNode:
		var args []string
		err := node.Decode(&args)
		if err != nil {
			return nil, fmt.Errorf("the arguments must be strings: %w", err)
		}
		return args, nil
	default:
		return nil, errors.New("unknown argument type")
	}
}
