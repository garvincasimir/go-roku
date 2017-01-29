package rokuclient

import (
	"testing"
)

//This test will fail if devices are not present on the network
func TestDiscover(t *testing.T) {
	devices := discover(5)

	if len(devices) == 0 {
		t.Fail()
	}

	t.Logf("%v",devices)
}
