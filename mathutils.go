package main

import (
	"log"
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

func random(min, max int) int {
	return min + rand.Intn(max-min)
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
	r := random(0, len(choices))
	log.Println("choices =", choices)
	log.Println("r =", r)
	return choices[r]
}
