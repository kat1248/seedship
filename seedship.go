package main

import "log"

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
		log.Println("done")
		return
	}
	planet := generatePlanet(systems)
	log.Println("planet =", planet)
}
