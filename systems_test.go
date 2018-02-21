package main

import (
	"testing"
)

func TestDamageSystem(t *testing.T) {
	var systems systemState
	systems.scannerAtmosphere = 100
	damageSystem(&systems, "atmosphere scanner", 10)
	if systems.scannerAtmosphere != 90 {
		t.Errorf("damage system incorrect, got %d, expected 90", systems.scannerAtmosphere)
	}
}

func TestGetSystemStrength(t *testing.T) {
	var systems systemState
	systems.scannerAtmosphere = 90
	foo := getSystemStrength(&systems, "atmosphere scanner")
	if foo != 90 {
		t.Errorf("system strenght incorrect, got %d, expected 90", foo)
	}
}
