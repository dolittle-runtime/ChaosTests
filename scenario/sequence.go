package scenario

import (
	"context"
	"math/rand"
	"time"

	"go.dolittle.io/chaostests/results"
)

var _ Action = &Sequence{}

// Sequence represents a sequence of Actions to perform
type Sequence struct {
	name    string
	actions []sequenceAction
}

// Name returns the name of the Sequence
func (s *Sequence) Name() string {
	return s.name
}

// Kind returns the kind of the Sequence
func (s *Sequence) Kind() string {
	return "sequence"
}

// Perform performs all of the Actions in the Sequence
func (s *Sequence) Perform(ctx context.Context, writer *results.Writer) error {
	writer.Log("Performing sequence")
	for _, action := range s.actions {
		actionWriter := writer.NewFor(action.Name(), action.Kind())
		shouldContinue, err := action.Perform(ctx, actionWriter)
		if err != nil {
			return err
		}
		if !shouldContinue {
			return nil
		}
		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}
	return nil
}

// AppendNewSequence creates a new Sequence and appends it to the Sequence
func (s *Sequence) AppendNewSequence(name string) *Sequence {
	sequence := &Sequence{
		name: name,
	}
	s.AppendAction(sequence)
	return sequence
}

// AppendNewParallel creates a new Parallel and appends it to the Sequence
func (s *Sequence) AppendNewParallel(name string) *Parallel {
	parallel := &Parallel{
		name: name,
	}
	s.AppendAction(parallel)
	return parallel
}

// AppendAction appends Action to be performed to the Sequence
func (s *Sequence) AppendAction(action Action) {
	s.actions = append(s.actions, &actionSequenceAction{
		action: action,
	})
}

// AppendDelay appens a delay to the Sequence
func (s *Sequence) AppendDelay(name string, duration time.Duration) {
	s.AppendAction(&delay{
		name:     name,
		duration: duration,
	})
}

// AppendStop appends a probabilistic stop to the Sequence
func (s *Sequence) AppendStop(name string, probability float32) {
	s.actions = append(s.actions, &stopSequenceAction{
		name:        name,
		probability: probability,
	})
}

type sequenceAction interface {
	Name() string
	Kind() string
	Perform(context.Context, *results.Writer) (bool, error)
}

type actionSequenceAction struct {
	action Action
}

func (a *actionSequenceAction) Name() string {
	return a.action.Name()
}

func (a *actionSequenceAction) Kind() string {
	return a.action.Kind()
}

func (a *actionSequenceAction) Perform(ctx context.Context, writer *results.Writer) (bool, error) {
	err := a.action.Perform(ctx, writer)
	return true, err
}

type stopSequenceAction struct {
	name        string
	probability float32
}

func (a *stopSequenceAction) Name() string {
	return a.name
}

func (a *stopSequenceAction) Kind() string {
	return "stop"
}

func (a *stopSequenceAction) Perform(ctx context.Context, writer *results.Writer) (bool, error) {
	if r := rand.Float32(); r < a.probability {
		return false, nil
	}
	return true, nil
}
