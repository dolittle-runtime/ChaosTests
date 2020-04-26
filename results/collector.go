package results

import (
	"fmt"
	"time"
)

// Collector collects results from Writers
type Collector struct {
	results chan interface{}
	done    chan interface{}
}

// NewCollector creates a new Collector
func NewCollector() *Collector {
	collector := &Collector{
		results: make(chan interface{}, 100),
		done:    make(chan interface{}),
	}
	go collector.process()
	return collector
}

// NewWriterForScenario creates a new Writer that writes to the Collector
func (c *Collector) NewWriterForScenario(name string) *Writer {
	return &Writer{
		collector: c,
		name:      name,
		kind:      "scenario",
	}
}

// Flush flushes the collected results
func (c *Collector) Flush() {
	close(c.results)
	<-c.done
}

func (c *Collector) process() {
	for message := range c.results {
		switch result := message.(type) {
		case logResult:
			fmt.Printf("%s [%s] (%s) %s\n", result.timestamp.Format(time.StampMilli), result.name, result.kind, result.message)
		}
	}
	close(c.done)
}

func (c *Collector) log(name, kind, format string, a ...interface{}) {
	c.results <- logResult{
		timestamp: time.Now(),
		name:      name,
		kind:      kind,
		message:   fmt.Sprintf(format, a...),
	}
}

type logResult struct {
	timestamp time.Time
	name      string
	kind      string
	message   string
}
