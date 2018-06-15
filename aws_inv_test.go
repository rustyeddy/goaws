package aws

import (
	"strings"
	"testing"
)

var (
	tstinv IMap
)

func init() {
}

// failNotEqual will do just that
func failNotEqual(t *testing.T, s1, s2 string) {
	if strings.Compare(s1, s2) != 0 {
		t.Errorf("expected equality s1 == s2 got (%s) and (%s) to ", s1, s2)
	}
}

func TestInventoryInit(t *testing.T) {
	// Expect inventories to be empty
	invs := Inventories()
	if invs == nil {
		t.Error("expected inventories got nil")
	}
	if len(invs) != 0 {
		t.Errorf("expected inventory len (0) got (%d) ", len(invs))
	}
}

func TestInventoryCreate(t *testing.T) {
	// make sure that our entry does not already exist
	tests := []struct {
		have string
		want *Inventory
	}{
		{"test-west-1", &Inventory{Name: "test-west-1"}},
		{"test-west-2", &Inventory{Name: "test-west-2"}},
	}

	for _, tt := range tests {
		if got := Inventories().Get(tt.have); got != nil {
			failNotEqual(t, tt.have, got.Name)
		}
	}

	// We should have two entries in our inventories database
	if len(Inventories()) != 2 {
		t.Errorf("inventories expected (%d) got (%d)", 2, len(inventories))
	}
}

func TestExist(t *testing.T) {
	invs := Inventories()

	if invs.Exists("us") {
		t.Errorf("expected no inventory for us found some")
	}

	inv := invs.Get("us")
	if inv != nil {
		t.Errorf("expected us inventory (nil) got (%v) ", inv)
	}

	inv = invs.Create("us")
	if inv == nil {
		t.Error("failed to create a new inventory for the us")
	}
	failNotEqual(t, "us", inv.Name)

	exists := invs.Exists("us")
	if exists == false {
		t.Error("expected us to exist but did not")
	}
	inv = invs.Get("us")
	failNotEqual(t, "us", inv.Name)

	if inv == nil {
		t.Error("expected inventory for (us) got none")
	}
	failNotEqual(t, "us", inv.Name)

	inv = &Inventory{Name: "Moon"}
	invs.Set("us", inv)

	inv = invs.Get("us")
	failNotEqual(t, inv.Name, "Moon")

	// Finally the inventory should have a single entry
	if len(invs) != 1 {
		t.Errorf("expected inventory to equal 1")
	}
}
