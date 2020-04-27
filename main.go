package main

import (
	"context"
	"fmt"

	"go.dolittle.io/chaostests/playbook"

	"go.dolittle.io/chaostests/results"
)

func main() {
	col := results.NewCollector()
	ctx := context.Background()
	sc, err := playbook.ParseScenarioFromFile("first.chaostest.json")
	if err != nil {
		fmt.Println("Error parsing scenario:", err)
		return
	}
	wr := col.NewWriterForScenario(sc.Name())
	if err := sc.Perform(ctx, wr); err != nil {
		fmt.Println("Error performing scenario:", err)
	}

	col.Flush()
}
