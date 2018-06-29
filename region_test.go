package goaws

import (
	"testing"
)

func TestNoCache(t *testing.T) {
}

// TestSaveReadRegions this file actualy calls AWS, should read from
// mock response to avoid call overhead.
func TestFetchRegions(t *testing.T) {
	regions = nil // reset regions
	regs := fetchRegions()
	if regs == nil {
		t.Error("failed to fetch regions")
		t.FailNow()
	}
	lregs, lregions := len(regs), len(regions)
	if lregs < 1 {
		t.Errorf("regions expected (%d) got (%d) ", lregions, lregs)
	}
	if lregions != lregs {
		t.Errorf("expected len (%d) got (%d)", lregs, lregions)
	}
}

// TestSaveReadRegions will make sure our regions are being saved
// and read to and from the disk as expected.
func TestSaveRegions(t *testing.T) {
	regs := Regions()
	_, err := cache.StoreObject("regions", regs)
	if err != nil {
		t.Error("failed to store regions ", err)
	}
}

// Regions() now the normal regions test with covers all abovce
func TestRegions(t *testing.T) {
	regs := Regions()
	if regs == nil || len(regs) < 1 {
		t.Error("failed to read regions")
	}
}
