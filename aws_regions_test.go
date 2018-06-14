package main

import (
	"testing"

	"github.com/rustyeddy/golib"
)

var (
	tstBasedir string
	tstRegions []string
)

func init() {
	tstRegions = []string{"us-west-1", "us-west-2", "us-central-1"}
}

func tErrorNil(t *testing.T, obj interface{}) {
	if obj == nil {
		t.Error("failed object is nil")
	}
}

// Perhaps create a interface and call prod or test
func tRegions() []string {
	regions = tstRegions
	return regions
}

// TestRegionsInit will ensure that the regions are nil before
// a call to acquire the regions.
func TestRegionsInit(t *testing.T) {
	if regions != nil || len(regions) > 0 {
		t.Errorf("expected len (0) got (%d) for regions ", len(regions))
	}
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
func TestSaveReadRegions(t *testing.T) {

	regs := Regions()
	fname := "tests/regions.json"
	saveRegions(fname, regs)

	//log.

	if golib.FileNotExists(fname) {
		t.Errorf("expected %s to exist, it does not", fname)
	}

	// Reset the regions and reread
	regs, regions = nil, nil
	if regs = readRegions(fname); regs == nil {
		t.Errorf("expected read %s but failed", fname)
	}

	if len(regs) < 1 {
		t.Errorf("expected many regions got (%d)", len(regs))
	}
}

// Regions() now the normal regions test with covers all abovce
func TestRegions(t *testing.T) {
	regs := Regions()
	if regs == nil || len(regs) < 1 {
		t.Error("failed to read regions")
	}
}
