package playbook

import (
	"encoding/json"

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
	for _, message := range p.Actions {
		extras := parallelExtras{
			RunToCompletion: true,
		}
		parsed, err := ParallelActions.ParseWithExtras(message, &extras)
		if err != nil {
			return err
		}
		switch action := parsed.(type) {
		case *Delay:
			parallel.AppendDelay(action.Name, action.Duration)
		case *Sequence:
			child := parallel.AppendNewSequence(action.Name, extras.RunToCompletion)
			if err := action.Build(child); err != nil {
				return err
			}
		default:
			return UnknownActionType("parallel", action)
		}
	}
	return nil
}

type parallelExtras struct {
	RunToCompletion bool `json:"waitForCompletion"`
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
