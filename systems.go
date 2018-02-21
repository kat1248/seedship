package main

import (
	"log"
)

type systemState struct {
	scannerAtmosphere        int
	scannerGravity           int
	scannerTemperature       int
	scannerResources         int
	scannerWater             int
	systemLanding            int
	systemConstructors       int
	systemCulturalDatabase   int
	systemScientificDatabase int
	colonists                int
	surfaceProbes            int
}

func damageSystem(systems *systemState, system string, amount int) {
	/* Silently apply damage to a system */
	/* system should be the name of the system */
	/* amount should be the amount of damage */
	switch system {
	case "atmosphere scanner":
		systems.scannerAtmosphere -= min(amount, systems.scannerAtmosphere)
		newIntegrity = systems.scannerAtmosphere
	case "gravity scanner":
		systems.scannerGravity -= min(amount, systems.scannerGravity)
		newIntegrity = systems.scannerGravity
	case "temperature scanner":
		systems.scannerTemperature -= min(amount, systems.scannerTemperature)
		newIntegrity = systems.scannerTemperature
	case "resources scanner":
		systems.scannerResources -= min(amount, systems.scannerResources)
		newIntegrity = systems.scannerResources
	case "water scanner":
		systems.scannerWater -= min(amount, systems.scannerWater)
		newIntegrity = systems.scannerWater
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
		systemStrength = systems.scannerAtmosphere
	case "gravity scanner":
		systemStrength = systems.scannerGravity
	case "temperature scanner":
		systemStrength = systems.scannerTemperature
	case "resources scanner":
		systemStrength = systems.scannerResources
	case "water scanner":
		systemStrength = systems.scannerWater
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
