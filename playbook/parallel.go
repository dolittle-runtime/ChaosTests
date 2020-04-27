package playbook

import (
	"encoding/json"
	"fmt"

	"go.dolittle.io/chaostests/scenario"
)

// ParallelActions is a ParsableGroup for actions allowed in a Parallel
var ParallelActions = NewParsableGroup("parallel", DelayParsable(), SequenceParsable())

// Parallel represents a scenario.Parallel
type Parallel struct {
	Name    string            `json:"name"`
	Actions []json.RawMessage `json:"actions"`
}

// Build appends all Parallel actions to a scenario.Parallel
func (p *Parallel) Build(parallel *scenario.Parallel) error {
	duplicateAction := func(duplicates int, action func(string) error) error {
		if duplicates < 2 {
			return action("")
		}
		for n := 0; n < duplicates; n++ {
			suffix := fmt.Sprintf("[%d]", n)
			if err := action(suffix); err != nil {
				return err
			}
		}
		return nil
	}

	for _, message := range p.Actions {
		extras := struct {
			RunToCompletion bool `json:"waitForCompletion"`
			Duplicate       int  `json:"duplicate"`
		}{true, 0}

		parsed, err := ParallelActions.ParseWithExtras(message, &extras)
		if err != nil {
			return err
		}

		switch action := parsed.(type) {
		case *Delay:
			duplicateAction(extras.Duplicate, func(suffix string) error {
				parallel.AppendDelay(action.Name+suffix, action.Duration)
				return nil
			})
		case *Sequence:
			err := duplicateAction(extras.Duplicate, func(suffix string) error {
				child := parallel.AppendNewSequence(action.Name+suffix, extras.RunToCompletion)
				if err := action.Build(child); err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				return err
			}
		default:
			return UnknownActionType("parallel", action)
		}
	}
	return nil
}

// ParallelParsable returns a Parsable for Parallel
func ParallelParsable() Parsable {
	return &parallelParsable{}
}

type parallelParsable struct{}

func (p *parallelParsable) Kind() string {
	return "parallel"
}

func (p *parallelParsable) Interface() interface{} {
	return &Parallel{}
}
