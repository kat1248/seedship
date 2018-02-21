package main

import "math/rand"

func init() {
	rand.Seed(42) // time.Now().UnixNano()
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
