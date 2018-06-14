package main

import (
	"testing"
)

var (
	tinv Inventory // test inventory
)

func init() {
	tinv.name = "testInventory"
	tinv.path = "etc"
}

// TestInv
func TestInv(t *testing.T) {
	if tinv.instances != nil {
		t.Errorf("  nil instances expected got (%v)", tinv.instances)
	}
}

// TestVolFile
func TestVolFile(t *testing.T) {
	// Just read one specific file
	pattern := "etc/vols-ap-northeast-2.json"
	tinv.ProcessFile(pattern)
	if len(tinv.volumes) < 2 {
		t.Fatalf("expected lots of volumes got (%d) ", len(tinv.volumes))
	}

	// See what we have for inventory
	sc, vc := tinv.Sizes()
	if sc < 1 && vc < 1 {
		t.Error("We have nothing for inventory")
	}

}

func TestInstFile(t *testing.T) {
	pattern := "etc/instances-eu-central-1.json"
	tinv.ProcessFile(pattern)
	if len(tinv.instances) < 2 {
		t.Errorf("expected lots of instances got (%d)", len(tinv.instances))
	}

	// See what we have for inventory
	sc, vc := tinv.Sizes()
	if sc < 1 && vc < 1 {
		t.Error("We have nothing for inventory")
	}
}
