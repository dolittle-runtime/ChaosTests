package playbook

import (
	"encoding/json"
	"time"
)

// Delay represents a scenario.Delay
type Delay struct {
	Name     string
	Duration time.Duration
}

// UnmarshalJSON unmarshals a Delay
func (d *Delay) UnmarshalJSON(data []byte) error {
	delay := struct {
		Name     string  `json:"name"`
		Duration float32 `json:"duration"`
	}{}
	if err := json.Unmarshal(data, &delay); err != nil {
		return err
	}
	d.Name = delay.Name
	d.Duration = time.Duration(delay.Duration*1000) * time.Millisecond
	return nil
}

// TimeDuration returns the duration of the Delay as a time.Duration
func (d *Delay) TimeDuration() time.Duration {
	return time.Duration(d.Duration*1000) * time.Millisecond
}

// DelayParsable returns a Parsable for Delay
func DelayParsable() Parsable {
	return &delayParsable{}
}

type delayParsable struct{}

func (p *delayParsable) Kind() string {
	return "delay"
}

func (p *delayParsable) Interface() interface{} {
	return &Delay{}
}
