package main

import (
	"testing"
)

func TestDamageSystem(t *testing.T) {
	systems := newSystemState()
	damageSystem(systems, atmosphereScanner, 10)
	if systems.atmosphere.strength != 90 {
		t.Errorf("damage system incorrect, got %d, expected 90", systems.atmosphere.strength)
	}
}

func TestGetSystemStrength(t *testing.T) {
	systems := newSystemState()
	foo := getSystemStrength(systems, atmosphereScanner)
	if foo != 100 {
		t.Errorf("system strength incorrect, got %d, expected 90", foo)
	}
}

func TestSystemToString(t *testing.T) {
	x := culturalDatabase
	if x.String() != "cultural database" {
		t.Errorf("string conversion incorrect, got %s, expected \"cultural database\"", x)
	}
}
