package libdocker

import (
	"testing"
)

func TestAllContainers(t *testing.T) {
	_, err := AllContainers()
	if err != nil {
		t.Errorf("Gettting Docker containers info should not fail. Error: %v", err)
	}
}
