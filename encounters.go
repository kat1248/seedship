package main

import "fmt"

type Encounter int

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

type EncounterList []Encounter

type EncounterCategory int

const (
	catCommon EncounterCategory = iota
	catUneventful
	catRare
	catMalfunction
)

var (
	encountersFirstTwo = []Encounter{
		encImpactChoice,
		encCometChoice,
		encMicrometeorite,
		encOverheating,
	}

	encountersUneventful = []Encounter{
		encUneventful1,
		encUneventful2,
		encUneventful3,
		encUneventful4,
		encUneventful5,
	}

	encountersCommon = []Encounter{
		encImpactChoice,
		encCometChoice,
		encMicrometeorite,
		encOverheating,
		encProtoplanetaryDisc,
		encAvoidDust,
		encSensorAnomaly,
		encRadiationBurst,
	}

	encountersRare = []Encounter{
		encRacistProgram,
		encTrailingDrone,
		encAlienSignal,
		encAlienDerelict,
		encAlienProbe,
		encReadDatabase,
	}

	encountersMalfunction = []Encounter{
		encProbeMalfunction,
		encComputerFailure,
		encStasisFailure,
		encSystemFailure,
		encScannerFailure,
	}
)

func (encounter Encounter) String() string {
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

func (category EncounterCategory) String() string {
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

func selectNextEncounter(systems *SystemState) Encounter {
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
			/* FIXME - remove an uneventful event from list */
			// r2 = random(0,(encountersUneventful.length)-1)
			// encounter = encountersUneventful[_r2]
			// encountersUneventful.splice(_r2,1)
			encounter = chooseEncounter(encountersMalfunction)
			systems.lastEncounterCategory = catUneventful
		} else if visited() < 3 {
			/* FIRST TWO EVENTS */
			encounter = chooseEncounter(encountersFirstTwo)
			systems.lastEncounterCategory = catCommon
		} else if r == 9 && len(encountersRare) > 0 && systems.lastEncounterCategory != catRare {
			/* RARE EVENTS - once each per playthrough */
			/* FIXME - remove rare event from list */
			// r2 = random(0,(encountersRare.length)-1)>>
			// encounter = encountersRare[_r2]>>
			// encountersRare.splice(_r2,1)>>
			encounter = chooseEncounter(encountersRare)
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

func chooseEncounter(choices EncounterList) Encounter {
	r := random(0, len(choices))
	return choices[r]
}

func handleEncounter(encounter Encounter) {
	fmt.Println("*** handling", encounter, "***")
}
