package main

import "fmt"

var (
	systems      systemState
	newIntegrity int
)

func init() {
}

func main() {
	systems := newSystemState()
	gameIntro()
	if done(systems) {
		fmt.Println("done")
		return
	}
	planet := generatePlanet(systems)
	fmt.Println("planet =", planet)
}
