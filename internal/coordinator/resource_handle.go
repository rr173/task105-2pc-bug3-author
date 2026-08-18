package coordinator

import "fmt"

func requireResourceHandle(name string, resource Resource, ok bool) error {
	if !ok || resource == nil {
		return fmt.Errorf("participant handle unavailable: %s", name)
	}
	return nil
}
