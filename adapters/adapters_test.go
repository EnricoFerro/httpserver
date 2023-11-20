package adapters_test

import (
	"httpserver/adapters"
	"testing"
)

func TestAdapters(t *testing.T) {

	awins, err := adapters.Adapters()
	if err != nil {
		t.Errorf("got panic %s", err.Error())
	}
	for _, awin := range awins {
		t.Logf("name=%s, description=%s\n", awin.Name, awin.Description)
		for _, ipnet := range awin.IPNets {
			t.Logf("addr=%s, mask=%s\n", ipnet.IP, ipnet.Mask)
		}
	}
}
