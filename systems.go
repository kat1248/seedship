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

/* ship systems type */
type ShipSystem int

const (
	atmosphereScanner  ShipSystem = iota
	gravityScanner     ShipSystem = iota
	temperatureScanner ShipSystem = iota
	resourcesScanner   ShipSystem = iota
	waterScanner       ShipSystem = iota
	landingSystem      ShipSystem = iota
	constructionSystem ShipSystem = iota
	scientificDatabase ShipSystem = iota
	culturalDatabase   ShipSystem = iota
	surfaceProbes      ShipSystem = iota
	sleepChambers      ShipSystem = iota
)

func (system ShipSystem) String() string {
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
	if system < atmosphereScanner || system > sleepChambers {
		return "unknown"
	}
	return systems[system]
}

/*
<<set $MO_encounters_first_two =
	[	"MO Impact Choice",
		"MO Comet Choice",
		"MO Micrometeorite",
		"MO Radiation Burst",
		"MO Overheating"] >>

<<set $MO_encounters_uneventful =
	[	"MO Uneventful 1",
		"MO Uneventful 2",
		"MO Uneventful 3",
		"MO Uneventful 4",
		"MO Uneventful 5" ] >>

<<set $MO_encounters_common =
	[	"MO Impact Choice",
		"MO Comet Choice",
		"MO Micrometeorite",
		"MO Protoplanetary Disc",
		"MO Avoid Dust",
		"MO Radiation Burst",
		"MO Sensor Anomaly",
		"MO Overheating"] >>

<<set $MO_encounters_rare =
	[	"MO Racist Program",
		"MO Trailing Drone",
		"MO Alien Signal",
		"MO Alien Derelict",
		"MO Alien Probe",
		"MO Read Databases"] >>

<<set $MO_encounters_malfunction =
	[	"MO Probe Malfunction",
		"MO Computer Failure",
		"MO Stasis Failure",
		"MO System Failure",
		"MO Scanner Failure"] >>



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
	planetsVisited           int
}

func newSystemState() *systemState {
	systems := systemState{
		atmosphere:               scanner{strength: maxStrength, level: 0},
		gravity:                  scanner{strength: maxStrength, level: 0},
		temperature:              scanner{strength: maxStrength, level: 0},
		resources:                scanner{strength: maxStrength, level: 0},
		water:                    scanner{strength: maxStrength, level: 0},
		offCourse:                false,
		surfaceProbeUsed:         false,
		surfaceProbes:            maxProbes,
		colonists:                maxColonists,
		systemLanding:            maxStrength,
		systemConstructors:       maxStrength,
		systemCulturalDatabase:   maxStrength,
		systemScientificDatabase: maxStrength,
		planetsVisited:           0,
	}
	return &systems
}

func damageSystem(systems *systemState, system ShipSystem, amount int) {
	/* Silently apply damage to a system */
	/* system should be the name of the system */
	/* amount should be the amount of damage */
	switch system {
	case atmosphereScanner:
		systems.atmosphere.strength -= min(amount, systems.atmosphere.strength)
		newIntegrity = systems.atmosphere.strength
	case gravityScanner:
		systems.gravity.strength -= min(amount, systems.gravity.strength)
		newIntegrity = systems.gravity.strength
	case temperatureScanner:
		systems.temperature.strength -= min(amount, systems.temperature.strength)
		newIntegrity = systems.temperature.strength
	case resourcesScanner:
		systems.resources.strength -= min(amount, systems.resources.strength)
		newIntegrity = systems.resources.strength
	case waterScanner:
		systems.water.strength -= min(amount, systems.water.strength)
		newIntegrity = systems.water.strength
	case landingSystem:
		systems.systemLanding -= min(amount, systems.systemLanding)
		newIntegrity = systems.systemLanding
	case constructionSystem:
		systems.systemConstructors -= min(amount, systems.systemConstructors)
		newIntegrity = systems.systemConstructors
	case culturalDatabase:
		systems.systemCulturalDatabase -= min(amount, systems.systemCulturalDatabase)
		newIntegrity = systems.systemCulturalDatabase
	case scientificDatabase:
		systems.systemScientificDatabase -= min(amount, systems.systemScientificDatabase)
		newIntegrity = systems.systemScientificDatabase
	case sleepChambers:
		/* Special: damage to the sleep chambers kills colonists */
		/* Never kill exactly 1 colonist or leave exactly 1 alive */
		systems.colonists -= clamp(amount, 2, systems.colonists)
		if systems.colonists == 1 {
			systems.colonists = 0
		}
		newIntegrity = systems.colonists
	case surfaceProbes:
		/* Special: damage to the surface probes destroys a surface probe */
		systems.surfaceProbes -= min(1, systems.surfaceProbes)
		newIntegrity = systems.surfaceProbes
	default:
		log.Println("damage_system: Unexpected system", system)
	}
}

func getSystemStrength(systems *systemState, system ShipSystem) int {
	var systemStrength int
	switch system {
	case atmosphereScanner:
		systemStrength = systems.atmosphere.strength
	case gravityScanner:
		systemStrength = systems.gravity.strength
	case temperatureScanner:
		systemStrength = systems.temperature.strength
	case resourcesScanner:
		systemStrength = systems.resources.strength
	case waterScanner:
		systemStrength = systems.water.strength
	case landingSystem:
		systemStrength = systems.systemLanding
	case constructionSystem:
		systemStrength = systems.systemConstructors
	case culturalDatabase:
		systemStrength = systems.systemCulturalDatabase
	case scientificDatabase:
		systemStrength = systems.systemScientificDatabase
	case sleepChambers:
		systemStrength = systems.colonists
	case surfaceProbes:
		systemStrength = systems.surfaceProbes
	default:
		log.Println("damage_system: Unexpected system", system)
	}
	return systemStrength
}
