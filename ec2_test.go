package aws

import (
	"log"
	"testing"
)

// TestVolFile
func TestVolFile(t *testing.T) {

	tinv := &Inventory{Name: "ap-northeast-2"}

	// Just read one specific file
	pattern := "tests/ap-northeast/vols-ap-northeast-2.json"
	tinv.ReadFile(pattern)
	if len(tinv.Volumes) < 2 {
		t.Fatalf("expected lots of volumes got (%d) ", len(tinv.Volumes))
	}
	sc, vc := tinv.Sizes()
	if sc < 1 && vc < 1 {
		t.Error("We have nothing for inventory")
	}
	log.Printf("%+v", tinv)
}

func TestInstFile(t *testing.T) {

	tinv := &Inventory{Name: "eu-central-1"}
	pattern := "tests/eu-central/instances-eu-central-1.json"
	tinv.ReadFile(pattern)
	if len(tinv.Instances) < 2 {
		t.Errorf("expected lots of instances got (%d)", len(tinv.Instances))
	}

	// See what we have for inventory
	sc, vc := tinv.Sizes()
	if sc < 1 && vc < 1 {
		t.Error("We have nothing for inventory")
	}
}
