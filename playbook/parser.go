package playbook

import (
	"encoding/json"
	"fmt"
)

// Parsable represents the type of a playbook action
type Parsable interface {
	Kind() string
	Interface() interface{}
}

// ParsableGroup is a group of Parsables that can be used to parse playbook actions
type ParsableGroup struct {
	kind      string
	parsables map[string]Parsable
}

// NewParsableGroup creates a new ParsableGroup with a given Kind from a set of Parsables
func NewParsableGroup(kind string, parsables ...Parsable) *ParsableGroup {
	group := &ParsableGroup{
		kind:      kind,
		parsables: make(map[string]Parsable),
	}
	for _, parsable := range parsables {
		group.parsables[parsable.Kind()] = parsable
	}
	return group
}

// WithExtraParsables creates a new ParsableGroup with given Kind by extending the ParsableGroup with extra Parsables
func (g *ParsableGroup) WithExtraParsables(kind string, parsables ...Parsable) *ParsableGroup {
	group := &ParsableGroup{
		kind:      kind,
		parsables: make(map[string]Parsable),
	}
	for kind, parsable := range g.parsables {
		group.parsables[kind] = parsable
	}
	for _, parsable := range parsables {
		group.parsables[parsable.Kind()] = parsable
	}
	return group
}

// Parse unmarshals a json.RawMessage into one of the Parsables in the ParsableGroup
func (g *ParsableGroup) Parse(message json.RawMessage) (interface{}, error) {
	action := rawAction{}
	if err := json.Unmarshal(message, &action); err != nil {
		return nil, err
	}
	for kind, parsable := range g.parsables {
		if action.Kind == kind {
			value := parsable.Interface()
			if err := json.Unmarshal(message, value); err != nil {
				return nil, err
			}
			return value, nil
		}
	}
	return nil, fmt.Errorf("Kind %s not applicable in %s", action.Kind, g.kind)
}

// ParseWithExtras unmarshals a json.RawMessage into one of the Parsables in the ParsableGroup and unmarshals extra properties into extras
func (g *ParsableGroup) ParseWithExtras(message json.RawMessage, extras interface{}) (interface{}, error) {
	result, err := g.Parse(message)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(message, extras); err != nil {
		return nil, err
	}
	return result, nil
}

type rawAction struct {
	Kind string `json:"kind"`
}
