package playbook

import (
	"encoding/json"
	"io/ioutil"

	"go.dolittle.io/chaostests/scenario"
)

// Scenario represents a scenario.Scenario
type Scenario struct {
	Sequence
}

// ParseScenario unmarshals a scenario.Scenario from JSON encoded data
func ParseScenario(data []byte) (*scenario.Scenario, error) {
	parsed := Scenario{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, err
	}

	result := scenario.NewScenario(parsed.Name)
	return result, parsed.Build(&result.Sequence)
}

// ParseScenarioFromFile reads the provided file and unmarshals a scenario.Scenario from the contents
func ParseScenarioFromFile(filename string) (*scenario.Scenario, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseScenario(data)
}
