package main

import (
	"fmt"
)

// Planet is the random elemets of a planet
type Planet struct {
	name            string
	temperature     Temperature
	gravity         Gravity
	resources       Resources
	atmosphere      Atmosphere
	water           Water
	nativeTechLevel int
	anomalies       AnomalyList
	surfaceFeatures SurfaceFeatureList
}

func (p Planet) String() string {
	s := p.name + "\n" +
		"temperature: " + p.temperature.String() + "\n" +
		"gravity: " + p.gravity.String() + "\n" +
		"resources: " + p.resources.String() + "\n" +
		"atmosphere: " + p.atmosphere.String() + "\n" +
		"water: " + p.water.String() + "\n" +
		"natives: " + fmt.Sprint(p.nativeTechLevel)
	if len(p.anomalies) > 0 {
		s += "\n" + "anomalies: " + p.anomalies.String()
	}
	if len(p.surfaceFeatures) > 0 {
		s += "\n" + "surface features: " + p.surfaceFeatures.String()
	}

	return s
}

func generatePlanet(systems *SystemState) *Planet {
	planet := Planet{
		name:            "fred",
		temperature:     tmpModerate,
		gravity:         grvModerate,
		resources:       rscNone,
		atmosphere:      atmNone,
		water:           wtrNone,
		nativeTechLevel: 0,
	}
	systems.surfaceProbeUsed = false
	systems.atmosphere.success = random(0, 99) < systems.atmosphere.strength
	systems.gravity.success = random(0, 99) < systems.gravity.strength
	systems.temperature.success = random(0, 99) < systems.temperature.strength
	systems.resources.success = random(0, 99) < systems.resources.strength
	systems.water.success = random(0, 99) < systems.water.strength

	greenChance := 2
	yellowChance := 3
	redChance := 5

	var r int

	/* ATMOSPHERE */
	if visited() == 1 {
		r = random(greenChance, greenChance+yellowChance+redChance-1)
	} else if !systems.atmosphere.success || systems.atmosphere.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.atmosphere.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.atmosphere.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner atmosphere level was", systems.atmosphere.level)
	}

	if r < greenChance {
		planet.atmosphere = atmBreathable
	} else if r < greenChance+yellowChance {
		planet.atmosphere = atmMarginal
	} else {
		planet.atmosphere = choose(atmCorrosive, atmToxic, atmNonBreathable, atmNone).(Atmosphere)
	}

	/* GRAVITY */
	if visited() == 1 {
		r = random(greenChance, greenChance+yellowChance+redChance-1)
	} else if !systems.gravity.success || systems.gravity.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.gravity.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.gravity.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner gravity level was", systems.gravity.level)
	}

	if r < greenChance {
		planet.gravity = grvModerate
	} else if r < greenChance+yellowChance {
		planet.gravity = choose(grvLow, grvHigh).(Gravity)
	} else {
		planet.gravity = choose(grvVeryLow, grvVeryHigh).(Gravity)
	}

	/* TEMPERATURE */
	if visited() == 1 {
		r = random(greenChance, greenChance+yellowChance+redChance-1)
	} else if !systems.temperature.success || systems.temperature.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.temperature.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.temperature.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner temperature level was", systems.temperature.level)
	}

	if r < greenChance {
		planet.temperature = tmpModerate
	} else if r < greenChance+yellowChance {
		planet.temperature = choose(tmpCold, tmpHot).(Temperature)
	} else {
		planet.temperature = choose(tmpVeryCold, tmpVeryHot).(Temperature)
	}

	/* WATER */
	if visited() == 1 {
		r = random(greenChance, greenChance+yellowChance+redChance-1)
	} else if !systems.water.success || systems.water.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.water.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.water.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner water level was", systems.water.level)
	}

	if r < greenChance {
		planet.water = wtrOceans
	} else if r < greenChance+yellowChance {
		planet.water = wtrPlanetWideOcean
	} else {
		planet.water = choose(wtrTrace, wtrNone).(Water)
	}

	/* RESOURCES */
	if visited() == 1 {
		r = random(1, greenChance+yellowChance+redChance-1)
	} else if !systems.resources.success || systems.resources.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.resources.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.resources.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner resources level was", systems.resources.level)
	}

	if r < greenChance {
		planet.resources = rscRich
	} else if r < greenChance+yellowChance {
		planet.resources = rscPoor
	} else {
		planet.resources = rscNone
	}

	/* Freeze the oceans at low temperatures */
	if planet.temperature == tmpVeryCold || (planet.temperature == tmpCold && random(0, 1) == 0) {
		if planet.water == wtrOceans {
			planet.water = wtrIceCaps
		} else if planet.water == wtrPlanetWideOcean {
			planet.water = wtrIceCoveredSurface
		}
	}

	/* No liquid water without atmosphere */
	if (planet.water == wtrOceans || planet.water == wtrPlanetWideOcean) && planet.atmosphere == atmNone {
		planet.atmosphere = atmNonBreathable
	}

	/* SURFACE FEATURES */
	planet.surfaceFeatures = []SurfaceFeature{}
	planet.anomalies = []Anomaly{}

	/* Moon - affects technology result */
	/* Moon is first, because the surface probe investigates it before landing on the planet itself */
	/* The higher the planet's gravity, the more likely it is to have a moon */
	moonChance := 0
	switch planet.gravity {
	case grvVeryLow:
		moonChance = 10
	case grvLow:
		moonChance = 20
	case grvModerate:
		moonChance = 30
	case grvHigh:
		moonChance = 40
	case grvVeryHigh:
		moonChance = 50
	}

	if random(0, 99) < moonChance {
		planet.surfaceFeatures = append(planet.surfaceFeatures, choose(sfBarrenMoon, sfMetalRichMoon, sfUnstableMoon).(SurfaceFeature))
		planet.anomalies = append(planet.anomalies, anMoon)
	}

	/* Aesthetics - affect culture result */
	/* Flat 20% chance of one or the other */
	if random(0, 4) == 0 {
		planet.surfaceFeatures = append(planet.surfaceFeatures, choose(sfOutstandingBeauty, sfOutstandingUgliness).(SurfaceFeature))
	}

	/* Caves? */
	/* No caves if surface is covered entirely in water or ice. Otherwise chance of caves is based on gravity. */
	cavesChance := 0
	if planet.water != wtrPlanetWideOcean && planet.water != wtrIceCoveredSurface {
		switch planet.gravity {
		case grvVeryLow:
			cavesChance = 50
		case grvLow:
			cavesChance = 40
		case grvModerate:
			cavesChance = 30
		case grvHigh:
			cavesChance = 20
		case grvVeryHigh:
			cavesChance = 10
		}
		if random(0, 99) < cavesChance {
			if random(0, 2) == 0 {
				planet.surfaceFeatures = append(planet.surfaceFeatures, sfAirtightCaves)
			}
			if random(0, 2) == 0 {
				planet.surfaceFeatures = append(planet.surfaceFeatures, sfInsulatedCaves)
			}
			if !featureInList(sfAirtightCaves, planet.surfaceFeatures) &&
				!featureInList(sfInsulatedCaves, planet.surfaceFeatures) {
				planet.surfaceFeatures = append(planet.surfaceFeatures, sfUnstableGeology)
			}
			planet.anomalies = append(planet.anomalies, anGeologicalAnomaly)
		}
	}

	/* LIFE */
	/* Plant life? */
	plantsChance := 0
	if true {
		switch planet.atmosphere {
		case atmBreathable:
			plantsChance = 100
		case atmMarginal:
			plantsChance = 100
		case atmNonBreathable:
			plantsChance = 25
		case atmToxic:
			plantsChance = 25
		case atmNone:
			plantsChance = 5
		case atmCorrosive:
			plantsChance = 5
		}
		if random(0, 99) < plantsChance {
			planet.surfaceFeatures = append(planet.surfaceFeatures, choose(sfPlantLife, sfEdiblePlants, sfPoisonousPlants).(SurfaceFeature))
			planet.anomalies = append(planet.anomalies, anVegetation)
		}
	}
	animalsChance := 0
	/* If plants, possibly animals */
	if featureInList(sfPlantLife, planet.surfaceFeatures) ||
		featureInList(sfEdiblePlants, planet.surfaceFeatures) ||
		featureInList(sfPoisonousPlants, planet.surfaceFeatures) {
		animalsChance = 50
		if random(0, 99) < animalsChance {
			planet.surfaceFeatures = append(planet.surfaceFeatures, choose(sfAnimalLife, sfUsefulAnimals, sfDangerousAnimals).(SurfaceFeature))
			planet.anomalies = append(planet.anomalies, anAnimalLife)
		}
	}

	planet.nativeTechLevel = 0
	sentientsChance := 0
	/* If animals, possibly sentient life */
	if featureInList(sfAnimalLife, planet.surfaceFeatures) ||
		featureInList(sfUsefulAnimals, planet.surfaceFeatures) ||
		featureInList(sfDangerousAnimals, planet.surfaceFeatures) {
		sentientsChance = 50
		if random(0, 99) < sentientsChance {
			planet.surfaceFeatures = append(planet.surfaceFeatures, sfIntelligentLife)
			/* Max tech level is determined by resources */
			if planet.resources == rscRich {
				planet.nativeTechLevel = random(0, 10)
			} else if planet.resources == rscPoor {
				planet.nativeTechLevel = random(0, 6)
			} else if planet.resources == rscNone {
				planet.nativeTechLevel = random(0, 4)
			} else {
				fmt.Println("Unexpected resources value", planet.resources)
			}
			if planet.nativeTechLevel >= 3 {
				/* Neolithic or higher */
				planet.anomalies = append(planet.anomalies, anPossibleStructures)
			}
			if planet.nativeTechLevel >= 8 {
				/* Atomic or higher */
				planet.anomalies = append(planet.anomalies, anElectromagneticActivity)
			}
		}
	}
	/* Probability notes: */
	/* A breathable planet has a 100% chance of plant life, a 50% chance of animal life, and a 25% chance of sentient life */
	/* A breathable planet with rich resources has a 2.5% chance of a space-age civilisation */

	/* Ruins? */
	/* To do: alter ruins chance depending on existing factors */
	ruinsChance := 10
	if random(0, 99) < ruinsChance {
		planet.surfaceFeatures = append(planet.surfaceFeatures, sfMonumentalRuins)
		if !anomalyInList(anPossibleStructures, planet.anomalies) {
			planet.anomalies = append(planet.anomalies, anPossibleStructures)
		}
	}
	if random(0, 99) < ruinsChance {
		planet.surfaceFeatures = append(planet.surfaceFeatures, sfHighTechRuins)
		if !anomalyInList(anPossibleStructures, planet.anomalies) {
			planet.anomalies = append(planet.anomalies, anPossibleStructures)
		}
	}
	if random(0, 99) < ruinsChance {
		planet.surfaceFeatures = append(planet.surfaceFeatures, sfDangerousRuins)
		if !anomalyInList(anPossibleStructures, planet.anomalies) {
			planet.anomalies = append(planet.anomalies, anPossibleStructures)
		}
	}
	if random(0, 99) < ruinsChance &&
		!featureInList(sfMonumentalRuins, planet.surfaceFeatures) &&
		!featureInList(sfHighTechRuins, planet.surfaceFeatures) &&
		!featureInList(sfDangerousRuins, planet.surfaceFeatures) {
		planet.surfaceFeatures = append(planet.surfaceFeatures, sfRegularGeologicalFormations)
		if !anomalyInList(anPossibleStructures, planet.anomalies) {
			planet.anomalies = append(planet.anomalies, anPossibleStructures)
		}
	}

	// [[Orbit planet]]

	return &planet
}
