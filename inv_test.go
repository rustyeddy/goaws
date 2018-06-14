package main

import "testing"

func TestNewInventory(t *testing.T) {
	inv := NewInventory("tst")
	failNotEqual(t, inv.Name, "tst")
}
