package main

import "testing"

func nilError(t *testing.T, obj interface{}) {
	if obj == nil {
		t.Error("failed object is nil")
	}
}

func TestRegions(t *testing.T) {
	regs := Regions()
	nilError(t, regs)

	if len(regs) < 1 {
		t.Error("regions should be a longer list")
	}
}
