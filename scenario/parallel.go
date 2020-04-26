package scenario

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"go.dolittle.io/chaostests/results"
)

var _ Action = &Parallel{}

// Parallel represents a set of Sequences to perform in parallel
type Parallel struct {
	name    string
	actions []parallelAction
}

// Name returns the name of the Parallel
func (p *Parallel) Name() string {
	return p.name
}

// Kind returns the kind of the Parallel
func (p *Parallel) Kind() string {
	return "parallel"
}

// Perform performs all of the Sequences in the Parallel
func (p *Parallel) Perform(ctx context.Context, writer *results.Writer) error {
	parallelContext, cancel := context.WithCancel(ctx)
	shouldComplete := sync.WaitGroup{}
	all := sync.WaitGroup{}
	errors := make(chan error, len(p.actions))

	for _, action := range p.actions {
		all.Add(1)
		if action.runToCompletion {
			shouldComplete.Add(1)
		}
		actionWriter := writer.NewFor(action.action.Name(), action.action.Kind())
		go func(action Action, writer *results.Writer, shouldSignalRunToCompletion bool) {
			if err := action.Perform(parallelContext, writer); err != nil {
				errors <- err
			}
			all.Done()
			if shouldSignalRunToCompletion {
				shouldComplete.Done()
			}
		}(action.action, actionWriter, action.runToCompletion)
	}

	shouldComplete.Wait()
	cancel()
	all.Wait()
	close(errors)

	return combineParallelErrors(errors)
}

// AppendNewSequence creates a new Sequence and appends it to the Parallel
func (p *Parallel) AppendNewSequence(name string, runToCompletion bool) *Sequence {
	sequence := &Sequence{
		name: name,
	}
	p.AppendAction(sequence, runToCompletion)
	return sequence
}

// AppendAction appends Action to be performed to the Parallel
func (p *Parallel) AppendAction(action Action, runToCompletion bool) {
	p.actions = append(p.actions, parallelAction{
		action:          action,
		runToCompletion: runToCompletion,
	})
}

// AppendDelay appends a delay to the Parallel
func (p *Parallel) AppendDelay(name string, duration time.Duration) {
	p.AppendAction(&delay{
		name:     name,
		duration: duration,
	}, true)
}

type parallelAction struct {
	action          Action
	runToCompletion bool
}

func combineParallelErrors(errors <-chan error) error {
	errorStrings := []string{"Parallel errors:"}
	receivedError := false
	for err := range errors {
		errorStrings = append(errorStrings, err.Error())
		receivedError = true
	}
	if receivedError {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}
