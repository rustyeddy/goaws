package main

import "testing"

var (
	regs []string
)

func init() {
	regs = nil
}

func nilError(t *testing.T, obj interface{}) {
	if obj == nil {
		t.Error("failed object is nil")
	}
}

func TestRegions(t *testing.T) {
	regs = Regions()
	nilError(t, regs)

	if len(regs) < 1 {
		t.Error("regions should be a longer list")
	}
}

func TestString(t *testing.T) {
	str := regs.String()
	failNotEqual(t, str, "somestring")
}

// TestSaveReadRegions
func TestSaveRegions(t *testing.T) {

	fname := "run/test-regions.json"
	if f, err := os.Stat(fname); os.IsExist(err) {
		ioutil.Rm(fname)
	}
	if err != nil {
		log.Error("%+v", err)
	}
	saveRegions(fname, Regions)
}
