package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func clamp(num, min, max int) int {
	if num <= min {
		return min
	}
	if num >= max {
		return max
	}
	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// random returns an integer [min, max]
func random(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomDamageLow() int {
	return random(1, 15)
}
func randomDamageMedium() int {
	return random(15, 35)
}

func randomDamageHigh() int {
	return random(35, 65)
}

func either(choices ...string) string {
	r := random(0, len(choices)-1)
	return choices[r]
}

func pickOne(choices []string) string {
	r := random(0, len(choices)-1)
	return choices[r]
}

func choose(choices ...interface{}) interface{} {
	r := random(0, len(choices)-1)
	return choices[r]
}
