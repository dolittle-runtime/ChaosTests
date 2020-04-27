package playbook

import (
	"encoding/json"
	"fmt"

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
	repeatAction := func(repeat int, action func(string) error) error {
		if repeat < 2 {
			return action("")
		}
		for n := 0; n < repeat; n++ {
			suffix := fmt.Sprintf("[%d]", n)
			if err := action(suffix); err != nil {
				return err
			}
		}
		return nil
	}

	for _, message := range s.Actions {
		extras := struct {
			Repeat int `json:"repeat"`
		}{0}

		parsed, err := SequenceActions.ParseWithExtras(message, &extras)
		if err != nil {
			return err
		}
		switch action := parsed.(type) {
		case *Delay:
			repeatAction(extras.Repeat, func(suffix string) error {
				sequence.AppendDelay(action.Name+suffix, action.Duration)
				return nil
			})
		case *Stop:
			repeatAction(extras.Repeat, func(suffix string) error {
				sequence.AppendStop(action.Name+suffix, action.Probability)
				return nil
			})
		case *Sequence:
			err := repeatAction(extras.Repeat, func(suffix string) error {
				child := sequence.AppendNewSequence(action.Name)
				if err := action.Build(child); err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				return err
			}
		case *Parallel:
			err := repeatAction(extras.Repeat, func(suffix string) error {
				child := sequence.AppendNewParallel(action.Name)
				if err := action.Build(child); err != nil {
					return err
				}
				return nil
			})
			if err != nil {
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
