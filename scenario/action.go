package scenario

import (
	"context"

	"go.dolittle.io/chaostests/results"
)

// Action represents a testing action to perform
type Action interface {
	Name() string
	Kind() string
	Perform(context.Context, *results.Writer) error
}
