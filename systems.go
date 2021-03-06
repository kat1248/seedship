package main

import (
	"log"
)

const (
	maxColonists               = 1000
	maxProbes                  = 10
	maxStrength                = 100
	maxScorePerPlanetAttribute = 500
)

// shipSystem is a set of attributes for a ship
type shipSystem int

const (
	sysAtmosphereScanner shipSystem = iota
	sysGravityScanner
	sysTemperatureScanner
	sysResourcesScanner
	sysWaterScanner
	sysLandingSystem
	sysConstructionSystem
	sysScientificDatabase
	sysCulturalDatabase
	sysSurfaceProbes
	sysSleepChambers
)

func (system shipSystem) String() string {
	systems := [...]string{
		"atmosphere scanner",
		"gravity scanner",
		"temperature scanner",
		"resources scanner",
		"water scanner",
		"landing system",
		"construction system",
		"scientific database",
		"cultural database",
		"surface probes",
		"sleep chambers"}
	if system < sysAtmosphereScanner || system > sysSleepChambers {
		return "unknown"
	}
	return systems[system]
}

/*
<<set $tech_level_names =
	[	"Pre-Stone Age",
		"Paleolithic",
		"Mesolithic",
		"Neolithic",
		"Bronze Age",
		"Iron Age",
		"Medieval",
		"Industrial",
		"Atomic Age",
		"Information Age",
		"Post-Singularity"]>>

<<set $culture_names =
	[	"Savagery",
		"Warring Tribes",
		"Brutal Chieftains",
		"Benevolent Chieftains",
		"Collective Rule",
		"Warring States",
		"Slave-Based Empire",
		"Oppressive Theocracy",
		"Benevolent Monarchy",
		"Egalitarian Republic",
		"Post-Nuclear Wasteland",
		"Warring Superpowers",
		"Dystopian Police State",
		"Rule by Corporations",
		"Corrupt Democracy",
		"Engaged Democracy",
		"Post-Scarcity Utopia",
		"Cosmic Enlightenment"]>>

<<set $native_relations_names =
	[	"Genocide of Colonists",
		"Genocide of Natives",
		"Colonists Enslaved",
		"Natives Enslaved",
		"Isolationism",
		"Friendly",
		"Integrated Societies"]>>

*/

// scannerType is the data for one of the ships scanners
type scannerType struct {
	strength int
	level    int
	success  bool
}

// SystemState is a collection of the different subsystem states
type SystemState struct {
	atmosphere               scannerType
	gravity                  scannerType
	temperature              scannerType
	resources                scannerType
	water                    scannerType
	systemLanding            int
	systemConstructors       int
	systemCulturalDatabase   int
	systemScientificDatabase int
	colonists                int
	surfaceProbes            int
	offCourse                bool
	surfaceProbeUsed         bool
	planetsVisited           int
	lastEncounter            encounterType
	lastEncounterCategory    encounterCategory
}

func (system SystemState) String() string {
	return ""
}

func newSystemState() *SystemState {
	systems := SystemState{
		atmosphere:               scannerType{strength: maxStrength, level: 0, success: false},
		gravity:                  scannerType{strength: maxStrength, level: 0, success: false},
		temperature:              scannerType{strength: maxStrength, level: 0, success: false},
		resources:                scannerType{strength: maxStrength, level: 0, success: false},
		water:                    scannerType{strength: maxStrength, level: 0, success: false},
		offCourse:                false,
		surfaceProbeUsed:         false,
		surfaceProbes:            maxProbes,
		colonists:                maxColonists,
		systemLanding:            maxStrength,
		systemConstructors:       maxStrength,
		systemCulturalDatabase:   maxStrength,
		systemScientificDatabase: maxStrength,
		planetsVisited:           0,
		lastEncounter:            encNone,
		lastEncounterCategory:    catCommon,
	}
	return &systems
}

func damageSystem(systems *SystemState, system shipSystem, amount int) {
	/* Silently apply damage to a system */
	/* system should be the name of the system */
	/* amount should be the amount of damage */
	switch system {
	case sysAtmosphereScanner:
		systems.atmosphere.strength -= min(amount, systems.atmosphere.strength)
		newIntegrity = systems.atmosphere.strength
	case sysGravityScanner:
		systems.gravity.strength -= min(amount, systems.gravity.strength)
		newIntegrity = systems.gravity.strength
	case sysTemperatureScanner:
		systems.temperature.strength -= min(amount, systems.temperature.strength)
		newIntegrity = systems.temperature.strength
	case sysResourcesScanner:
		systems.resources.strength -= min(amount, systems.resources.strength)
		newIntegrity = systems.resources.strength
	case sysWaterScanner:
		systems.water.strength -= min(amount, systems.water.strength)
		newIntegrity = systems.water.strength
	case sysLandingSystem:
		systems.systemLanding -= min(amount, systems.systemLanding)
		newIntegrity = systems.systemLanding
	case sysConstructionSystem:
		systems.systemConstructors -= min(amount, systems.systemConstructors)
		newIntegrity = systems.systemConstructors
	case sysCulturalDatabase:
		systems.systemCulturalDatabase -= min(amount, systems.systemCulturalDatabase)
		newIntegrity = systems.systemCulturalDatabase
	case sysScientificDatabase:
		systems.systemScientificDatabase -= min(amount, systems.systemScientificDatabase)
		newIntegrity = systems.systemScientificDatabase
	case sysSleepChambers:
		/* Special: damage to the sleep chambers kills colonists */
		/* Never kill exactly 1 colonist or leave exactly 1 alive */
		systems.colonists -= clamp(amount, 2, systems.colonists)
		if systems.colonists == 1 {
			systems.colonists = 0
		}
		newIntegrity = systems.colonists
	case sysSurfaceProbes:
		/* Special: damage to the surface probes destroys a surface probe */
		systems.surfaceProbes -= min(1, systems.surfaceProbes)
		newIntegrity = systems.surfaceProbes
	default:
		log.Println("damage_system: Unexpected system", system)
	}
}

func getSystemStrength(systems *SystemState, system shipSystem) int {
	var systemStrength int
	switch system {
	case sysAtmosphereScanner:
		systemStrength = systems.atmosphere.strength
	case sysGravityScanner:
		systemStrength = systems.gravity.strength
	case sysTemperatureScanner:
		systemStrength = systems.temperature.strength
	case sysResourcesScanner:
		systemStrength = systems.resources.strength
	case sysWaterScanner:
		systemStrength = systems.water.strength
	case sysLandingSystem:
		systemStrength = systems.systemLanding
	case sysConstructionSystem:
		systemStrength = systems.systemConstructors
	case sysCulturalDatabase:
		systemStrength = systems.systemCulturalDatabase
	case sysScientificDatabase:
		systemStrength = systems.systemScientificDatabase
	case sysSleepChambers:
		systemStrength = systems.colonists
	case sysSurfaceProbes:
		systemStrength = systems.surfaceProbes
	default:
		log.Println("damage_system: Unexpected system", system)
	}
	return systemStrength
}
