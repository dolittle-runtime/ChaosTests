package playbook

import (
	"fmt"
)

// UnknownActionType returns an error to use when an unknown action type is encountered while building a kind
func UnknownActionType(kind string, action interface{}) error {
	return fmt.Errorf("Unknown action type %T in %s", action, kind)
}
