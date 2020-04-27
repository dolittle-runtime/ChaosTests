package playbook

import (
	"encoding/json"

	"go.dolittle.io/chaostests/scenario"
)

// SequenceActions is a ParsableGroup for actions allowed in a Sequence
var SequenceActions = NewParsableGroup("sequence",
	DelayParsable(),
	SequenceParsable(),
	ParallelParsable(),
	SequenceStopParsable())

// Sequence represents a scenario.Sequence
type Sequence struct {
	Name    string `json:"name"`
	Actions []json.RawMessage
}

// Build appends all Sequence actions to a scenario.Sequence
func (s *Sequence) Build(sequence *scenario.Sequence) error {
	for _, message := range s.Actions {
		parsed, err := SequenceActions.Parse(message)
		if err != nil {
			return err
		}
		switch action := parsed.(type) {
		case *Delay:
			sequence.AppendDelay(action.Name, action.Duration)
		case *Stop:
			sequence.AppendStop(action.Name, action.Probability)
		case *Sequence:
			child := sequence.AppendNewSequence(action.Name)
			if err := action.Build(child); err != nil {
				return err
			}
		case *Parallel:
			child := sequence.AppendNewParallel(action.Name)
			if err := action.Build(child); err != nil {
				return err
			}
		default:
			return UnknownActionType("sequence", action)
		}
	}
	return nil
}

// SequenceParsable returns a Parsable for Sequence
func SequenceParsable() Parsable {
	return &sequenceParsable{}
}

type sequenceParsable struct{}

func (p *sequenceParsable) Kind() string {
	return "sequence"
}

func (p *sequenceParsable) Interface() interface{} {
	return &Sequence{}
}

// Stop represents a scenario.Sequence.Stop action
type Stop struct {
	Name        string  `json:"name"`
	Probability float32 `json:"probability"`
}

// SequenceStopParsable returns a Parsable for Sequence.Stop
func SequenceStopParsable() Parsable {
	return &sequenceStopParsable{}
}

type sequenceStopParsable struct{}

func (p *sequenceStopParsable) Kind() string {
	return "stop"
}

func (p *sequenceStopParsable) Interface() interface{} {
	return &Stop{}
}
