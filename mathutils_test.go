package main

import (
	"testing"
)

func TestClamp(t *testing.T) {
	x := clamp(5, 2, 6)
	if x != 5 {
		t.Errorf("Clamp was incorrect, got %d, expected 5", x)
	}
	x = clamp(10, 1, 8)
	if x != 8 {
		t.Errorf("Clamp was incorrect, got %d, expected 8", x)
	}
	x = clamp(3, 6, 9)
	if x != 6 {
		t.Errorf("Clamp was incorrect, got %d, expected 6", x)
	}
}

func TestMin(t *testing.T) {
	x := min(5, 18)
	if x != 5 {
		t.Errorf("Clamp was incorrect, got %d, expected 5", x)
	}
	x = min(23, 6)
	if x != 6 {
		t.Errorf("Clamp was incorrect, got %d, expected 6", x)
	}
}

func TestRandom(t *testing.T) {
	low := false
	high := false
	for i := 0; i < 1000; i++ {
		x := random(20, 52)
		if x < 20 {
			t.Error("random out of range [20, 52], got", x)
			return
		}
		if x > 52 {
			t.Error("random out of range [20, 52], got", x)
			return
		}
		if x == 20 {
			low = true
		} else if x == 52 {
			high = true
		}
	}
	if !low {
		t.Error("random lower bound never met")
	}
	if !high {
		t.Error("random higher bound never met")
	}
}
