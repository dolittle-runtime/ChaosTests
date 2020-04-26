package scenario

var _ Action = &Scenario{}

// Scenario represents a test scenario that can be performed
type Scenario struct {
	Sequence
}

// NewScenario creates a new Scenario with the provided name
func NewScenario(name string) *Scenario {
	return &Scenario{
		Sequence: Sequence{
			name: name,
		},
	}
}

// Kind returns the kind of the Scenario
func (s *Scenario) Kind() string {
	return "scenario"
}
