package main

import "fmt"

// encounterType is the id of a random encounter
type encounterType int

const (
	encNone = iota
	encImpactChoice
	encCometChoice
	encMicrometeorite
	encRadiationBurst
	encOverheating
	encUneventful1
	encUneventful2
	encUneventful3
	encUneventful4
	encUneventful5
	encProtoplanetaryDisc
	encAvoidDust
	encSensorAnomaly
	encRacistProgram
	encTrailingDrone
	encAlienSignal
	encAlienDerelict
	encAlienProbe
	encReadDatabase
	encProbeMalfunction
	encComputerFailure
	encStasisFailure
	encSystemFailure
	encScannerFailure
)

// encounterList is a list of encounters
type encounterList []encounterType

// encounterCategory is which category an encounter is
type encounterCategory int

const (
	catCommon encounterCategory = iota
	catUneventful
	catRare
	catMalfunction
)

var (
	encountersFirstTwo = []encounterType{
		encImpactChoice,
		encCometChoice,
		encMicrometeorite,
		encOverheating,
	}

	encountersUneventfulBase = []encounterType{
		encUneventful1,
		encUneventful2,
		encUneventful3,
		encUneventful4,
		encUneventful5,
	}

	encountersCommon = []encounterType{
		encImpactChoice,
		encCometChoice,
		encMicrometeorite,
		encOverheating,
		encProtoplanetaryDisc,
		encAvoidDust,
		encSensorAnomaly,
		encRadiationBurst,
	}

	encountersRareBase = []encounterType{
		encRacistProgram,
		encTrailingDrone,
		encAlienSignal,
		encAlienDerelict,
		encAlienProbe,
		encReadDatabase,
	}

	encountersMalfunction = []encounterType{
		encProbeMalfunction,
		encComputerFailure,
		encStasisFailure,
		encSystemFailure,
		encScannerFailure,
	}
	encountersUneventful = encountersUneventfulBase
	encountersRare       = encountersRareBase
)

func (encounter encounterType) String() string {
	encounters := [...]string{
		"None",
		"Impact Choice",
		"Comet Choice",
		"Micrometeorite",
		"Radiation Burst",
		"Overheating",
		"Uneventful 1",
		"Uneventful 2",
		"Uneventful 3",
		"Uneventful 4",
		"Uneventful 5",
		"Protoplanetary Disc",
		"Avoid Dust",
		"Sensor Anomaly",
		"Racist Program",
		"Trailing Drone",
		"Alien Signal",
		"Alien Derelict",
		"Alien Probe",
		"Read Databases",
		"Probe Malfunction",
		"Computer Failure",
		"Stasis Failure",
		"System Failure",
		"Scanner Failure",
	}
	if encounter < encNone || encounter > encScannerFailure {
		return "unknown"
	}
	return encounters[encounter]
}

func (category encounterCategory) String() string {
	categories := [...]string{
		"Common",
		"Uneventful",
		"Rare",
		"Malfunction",
	}
	if category < catCommon || category > catMalfunction {
		return "unknown"
	}
	return categories[category]
}

func selectNextEncounter(systems *SystemState) encounterType {
	// Initial: 2/10 uneventful, 7/10 common, 1/10 rare
	// Final: 5/10 malfunction, 4/10 common, 1/10 rare

	malfunctionChance := min(visited()-5, 5)

	encounter := systems.lastEncounter
	for encounter == systems.lastEncounter {
		r := random(0, 9)
		if r < malfunctionChance {
			/* MALFUNCTIONS */
			encounter = chooseEncounter(encountersMalfunction)
			systems.lastEncounterCategory = catMalfunction
		} else if r < 2 && len(encountersUneventful) > 0 && systems.lastEncounterCategory != catUneventful {
			/* NOTHING INTERESTING HAPPENS - once each per playthrough */
			r2 := random(0, len(encountersUneventful)-1)
			encounter = encountersUneventful[r2]
			encountersUneventful = append(encountersUneventful[:r2], encountersUneventful[r2+1:]...)
			systems.lastEncounterCategory = catUneventful
		} else if visited() < 3 {
			/* FIRST TWO EVENTS */
			encounter = chooseEncounter(encountersFirstTwo)
			systems.lastEncounterCategory = catCommon
		} else if r == 9 && len(encountersRare) > 0 && systems.lastEncounterCategory != catRare {
			/* RARE EVENTS - once each per playthrough */
			r2 := random(0, len(encountersRare)-1)
			encounter = encountersRare[r2]
			encountersRare = append(encountersRare[:r2], encountersRare[r2+1:]...)
			systems.lastEncounterCategory = catRare
		} else {
			/* COMMON EVENTS */
			encounter = chooseEncounter(encountersCommon)
			systems.lastEncounterCategory = catCommon
		}
	}
	systems.lastEncounter = encounter

	return encounter
}

func chooseEncounter(choices encounterList) encounterType {
	r := random(0, len(choices)-1)
	return choices[r]
}

func handleEncounter(encounter encounterType) {
	fmt.Println("*** handling", encounter, "***")
}
