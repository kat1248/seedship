package main

import (
	"log"
)

type scanner struct {
	strength int
	level    int
}

type systemState struct {
	atmosphere               scanner
	gravity                  scanner
	temperature              scanner
	resources                scanner
	water                    scanner
	systemLanding            int
	systemConstructors       int
	systemCulturalDatabase   int
	systemScientificDatabase int
	colonists                int
	surfaceProbes            int
	offCourse                bool
	surfaceProbeUsed         bool
}

func newScanner() scanner {
	s := scanner{strength: 100, level: 0}
	return s
}

func newSystemState() *systemState {
	systems := systemState{
		atmosphere:               newScanner(),
		gravity:                  newScanner(),
		temperature:              newScanner(),
		resources:                newScanner(),
		water:                    newScanner(),
		offCourse:                false,
		surfaceProbeUsed:         false,
		surfaceProbes:            10,
		colonists:                1000,
		systemLanding:            100,
		systemConstructors:       100,
		systemCulturalDatabase:   100,
		systemScientificDatabase: 100,
	}
	return &systems
}

func damageSystem(systems *systemState, system string, amount int) {
	/* Silently apply damage to a system */
	/* system should be the name of the system */
	/* amount should be the amount of damage */
	switch system {
	case "atmosphere scanner":
		systems.atmosphere.strength -= min(amount, systems.atmosphere.strength)
		newIntegrity = systems.atmosphere.strength
	case "gravity scanner":
		systems.gravity.strength -= min(amount, systems.gravity.strength)
		newIntegrity = systems.gravity.strength
	case "temperature scanner":
		systems.temperature.strength -= min(amount, systems.temperature.strength)
		newIntegrity = systems.temperature.strength
	case "resources scanner":
		systems.resources.strength -= min(amount, systems.resources.strength)
		newIntegrity = systems.resources.strength
	case "water scanner":
		systems.water.strength -= min(amount, systems.water.strength)
		newIntegrity = systems.water.strength
	case "landing system":
		systems.systemLanding -= min(amount, systems.systemLanding)
		newIntegrity = systems.systemLanding
	case "construction system":
		systems.systemConstructors -= min(amount, systems.systemConstructors)
		newIntegrity = systems.systemConstructors
	case "cultural database":
		systems.systemCulturalDatabase -= min(amount, systems.systemCulturalDatabase)
		newIntegrity = systems.systemCulturalDatabase
	case "scientific database":
		systems.systemScientificDatabase -= min(amount, systems.systemScientificDatabase)
		newIntegrity = systems.systemScientificDatabase
	case "sleep chambers":
		/* Special: damage to the sleep chambers kills colonists */
		/* Never kill exactly 1 colonist or leave exactly 1 alive */
		systems.colonists -= clamp(amount, 2, systems.colonists)
		if systems.colonists == 1 {
			systems.colonists = 0
		}
		newIntegrity = systems.colonists
	case "surface probes":
		/* Special: damage to the surface probes destroys a surface probe */
		systems.surfaceProbes -= min(1, systems.surfaceProbes)
		newIntegrity = systems.surfaceProbes
	default:
		log.Println("damage_system: Unexpected system name", system)
	}
}

func getSystemStrength(systems *systemState, system string) int {
	var systemStrength int
	switch system {
	case "atmosphere scanner":
		systemStrength = systems.atmosphere.strength
	case "gravity scanner":
		systemStrength = systems.gravity.strength
	case "temperature scanner":
		systemStrength = systems.temperature.strength
	case "resources scanner":
		systemStrength = systems.resources.strength
	case "water scanner":
		systemStrength = systems.water.strength
	case "landing system":
		systemStrength = systems.systemLanding
	case "construction system":
		systemStrength = systems.systemConstructors
	case "cultural database":
		systemStrength = systems.systemCulturalDatabase
	case "scientific database":
		systemStrength = systems.systemScientificDatabase
	case "sleep chambers":
		systemStrength = systems.colonists
	case "surface probes":
		systemStrength = systems.surfaceProbes
	default:
		log.Println("damage_system: Unexpected system name", system)
	}
	return systemStrength
}
