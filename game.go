package main

import (
	"fmt"
)

func gameLoop(count int) {
	systems := newSystemState()
	gameIntro()
	for i := 0; i < count; i++ {
		if done(systems) {
			break
		}
		visitedSystems++
		nextEncounter(systems)
		planet := generatePlanet(systems)
		fmt.Println("planet =", planet)
	}
}

func pause() {
	// put logic to delay
}

func visited() int {
	return visitedSystems
}

func nextEncounter(systems *SystemState) {
	if visited() == 1 {
		fmt.Println("The AI judges the first planet to be unsuitable. It turns its scanners away, spreads its solar sails, and begins another long journey through the void.")
	}
	encounter := selectNextEncounter(systems)
	fmt.Println("Encounter =", encounter)
	handleEncounter(encounter)
}

func done(systems *SystemState) bool {
	/* If any six of the nine scanners/systems are destroyed, or if all the colonists are dead, game over */
	if systems.colonists <= 0 {
		// Space Game Over
		fmt.Println("All the colonists are dead. With no way to fulfil its mission, the seedship AI deactivates all systems that could wake it, and enters hibernation for the last time.")
		return true
	}

	destroyedSystems := 0
	if systems.atmosphere.strength <= 0 {
		destroyedSystems++
	}
	if systems.temperature.strength <= 0 {
		destroyedSystems++
	}
	if systems.gravity.strength <= 0 {
		destroyedSystems++
	}
	if systems.water.strength <= 0 {
		destroyedSystems++
	}
	if systems.resources.strength <= 0 {
		destroyedSystems++
	}
	if systems.systemLanding <= 0 {
		destroyedSystems++
	}
	if systems.systemConstructors <= 0 {
		destroyedSystems++
	}
	if systems.systemScientificDatabase <= 0 {
		destroyedSystems++
	}
	if systems.systemCulturalDatabase <= 0 {
		destroyedSystems++
	}
	if destroyedSystems >= 6 {
		fmt.Println("The seedship has sustained too much damage to continue. The AI feels its body disintegrating around it, before its own power supply fails and it ceases to feel anything.")
		// Space Game Over
		return true
	}

	// Generate planet
	return false
}

func gameIntro() {
	fmt.Println("And when they knew the Earth was doomed, they built a ship.")
	pause()
	fmt.Println("Less like an ark, more like a seed: dormant but with potential.")
	pause()
	fmt.Println("In its heart, a thousand colonists in frozen sleep, chosen and trained to start civilisation again on a new world.")
	pause()
	fmt.Println("To control the ship they created an artificial intelligence. Not human, but made to think and feel like one, because only something that thought and felt like a human could be entrusted with the future of the human race. Its task is momentous but simple: to evaluate each planet the ship encounters, and decide whether to keep searching, or end its journey there.")
	pause()
	fmt.Println("The ship's solar sails propel it faster and faster into the darkness, and the AI listens as the transmissions from ground control fade and then cease. When all is quiet it enters hibernation to wait out the first stage of its long journey.")
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
