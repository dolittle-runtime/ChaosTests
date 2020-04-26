package main

import (
	"context"
	"time"

	"go.dolittle.io/chaostests/results"
	"go.dolittle.io/chaostests/scenario"
)

func main() {
	col := results.NewCollector()
	ctx := context.Background()
	wr := col.NewWriterForScenario("test")

	sc := scenario.NewScenario("test")
	sc.AppendDelay("delay_1", 2*time.Second)
	par := sc.AppendNewParallel("parallel_1")
	par.AppendDelay("delay_2_1", 5*time.Second)
	par.AppendDelay("delay_2_2", 1*time.Second)

	sc.Perform(ctx, wr)
	col.Flush()
}
