package scenario

import (
	"context"
	"time"

	"go.dolittle.io/chaostests/results"
)

var _ Action = &delay{}

type delay struct {
	name     string
	duration time.Duration
}

func (d *delay) Name() string {
	return d.name
}

func (d *delay) Kind() string {
	return "delay"
}

func (d *delay) Perform(ctx context.Context, writer *results.Writer) error {
	timer := time.NewTimer(d.duration)
	writer.Log("Waiting for %.4f seconds", float32(d.duration/time.Millisecond)/1000)
	select {
	case <-timer.C:
	case <-ctx.Done():
	}
	return nil
}
