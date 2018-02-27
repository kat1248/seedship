package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Planet is the random elemets of a planet
type Planet struct {
	Name            string
	Temperature     temperatureType
	Gravity         gravityType
	Resources       resourcesType
	Atmosphere      atmosphereType
	Water           waterType
	NativeTechLevel int
	Anomalies       anomalyList
	SurfaceFeatures surfaceFeatureList
}

const planetTemplate = `{{.Name}}
temperature: {{.Temperature}}
gravity: {{.Gravity}}
resources: {{.Resources}}
atmosphere: {{.Atmosphere}}
water: {{.Water}}
natives: {{.NativeTechLevel}}{{if .Anomalies}}
anomalies: {{.Anomalies}}{{end}}{{if .SurfaceFeatures}}
surface features: {{.SurfaceFeatures}}{{end}}`

func (p Planet) String() string {
	tmpl := template.New("planet")
	tmpl, err := tmpl.Parse(planetTemplate)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var tpl bytes.Buffer
	if err = tmpl.Execute(&tpl, p); err != nil {
		fmt.Println(err)
		return ""
	}

	return tpl.String()
}

func generatePlanet(systems *SystemState) *Planet {
	planet := Planet{
		Name:            "fred",
		Temperature:     tmpModerate,
		Gravity:         grvModerate,
		Resources:       rscNone,
		Atmosphere:      atmNone,
		Water:           wtrNone,
		NativeTechLevel: 0,
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
		planet.Atmosphere = atmBreathable
	} else if r < greenChance+yellowChance {
		planet.Atmosphere = atmMarginal
	} else {
		planet.Atmosphere = choose(atmCorrosive, atmToxic, atmNonBreathable, atmNone).(atmosphereType)
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
		planet.Gravity = grvModerate
	} else if r < greenChance+yellowChance {
		planet.Gravity = choose(grvLow, grvHigh).(gravityType)
	} else {
		planet.Gravity = choose(grvVeryLow, grvVeryHigh).(gravityType)
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
		planet.Temperature = tmpModerate
	} else if r < greenChance+yellowChance {
		planet.Temperature = choose(tmpCold, tmpHot).(temperatureType)
	} else {
		planet.Temperature = choose(tmpVeryCold, tmpVeryHot).(temperatureType)
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
		planet.Water = wtrOceans
	} else if r < greenChance+yellowChance {
		planet.Water = wtrPlanetWideOcean
	} else {
		planet.Water = choose(wtrTrace, wtrNone).(waterType)
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
		planet.Resources = rscRich
	} else if r < greenChance+yellowChance {
		planet.Resources = rscPoor
	} else {
		planet.Resources = rscNone
	}

	/* Freeze the oceans at low temperatures */
	if planet.Temperature == tmpVeryCold || (planet.Temperature == tmpCold && random(0, 1) == 0) {
		if planet.Water == wtrOceans {
			planet.Water = wtrIceCaps
		} else if planet.Water == wtrPlanetWideOcean {
			planet.Water = wtrIceCoveredSurface
		}
	}

	/* No liquid water without atmosphere */
	if (planet.Water == wtrOceans || planet.Water == wtrPlanetWideOcean) && planet.Atmosphere == atmNone {
		planet.Atmosphere = atmNonBreathable
	}

	/* SURFACE FEATURES */
	planet.SurfaceFeatures = []surfaceFeatureType{}
	planet.Anomalies = []anomalyType{}

	/* Moon - affects technology result */
	/* Moon is first, because the surface probe investigates it before landing on the planet itself */
	/* The higher the planet's gravity, the more likely it is to have a moon */
	moonChance := 0
	switch planet.Gravity {
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
		planet.SurfaceFeatures = append(planet.SurfaceFeatures, choose(sfBarrenMoon, sfMetalRichMoon, sfUnstableMoon).(surfaceFeatureType))
		planet.Anomalies = append(planet.Anomalies, anMoon)
	}

	/* Aesthetics - affect culture result */
	/* Flat 20% chance of one or the other */
	if random(0, 4) == 0 {
		planet.SurfaceFeatures = append(planet.SurfaceFeatures, choose(sfOutstandingBeauty, sfOutstandingUgliness).(surfaceFeatureType))
	}

	/* Caves? */
	/* No caves if surface is covered entirely in water or ice. Otherwise chance of caves is based on gravity. */
	cavesChance := 0
	if planet.Water != wtrPlanetWideOcean && planet.Water != wtrIceCoveredSurface {
		switch planet.Gravity {
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
				planet.SurfaceFeatures = append(planet.SurfaceFeatures, sfAirtightCaves)
			}
			if random(0, 2) == 0 {
				planet.SurfaceFeatures = append(planet.SurfaceFeatures, sfInsulatedCaves)
			}
			if !featureInList(sfAirtightCaves, planet.SurfaceFeatures) &&
				!featureInList(sfInsulatedCaves, planet.SurfaceFeatures) {
				planet.SurfaceFeatures = append(planet.SurfaceFeatures, sfUnstableGeology)
			}
			planet.Anomalies = append(planet.Anomalies, anGeologicalAnomaly)
		}
	}

	/* LIFE */
	/* Plant life? */
	plantsChance := 0
	if true {
		switch planet.Atmosphere {
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
			planet.SurfaceFeatures = append(planet.SurfaceFeatures, choose(sfPlantLife, sfEdiblePlants, sfPoisonousPlants).(surfaceFeatureType))
			planet.Anomalies = append(planet.Anomalies, anVegetation)
		}
	}
	animalsChance := 0
	/* If plants, possibly animals */
	if featureInList(sfPlantLife, planet.SurfaceFeatures) ||
		featureInList(sfEdiblePlants, planet.SurfaceFeatures) ||
		featureInList(sfPoisonousPlants, planet.SurfaceFeatures) {
		animalsChance = 50
		if random(0, 99) < animalsChance {
			planet.SurfaceFeatures = append(planet.SurfaceFeatures, choose(sfAnimalLife, sfUsefulAnimals, sfDangerousAnimals).(surfaceFeatureType))
			planet.Anomalies = append(planet.Anomalies, anAnimalLife)
		}
	}

	planet.NativeTechLevel = 0
	sentientsChance := 0
	/* If animals, possibly sentient life */
	if featureInList(sfAnimalLife, planet.SurfaceFeatures) ||
		featureInList(sfUsefulAnimals, planet.SurfaceFeatures) ||
		featureInList(sfDangerousAnimals, planet.SurfaceFeatures) {
		sentientsChance = 50
		if random(0, 99) < sentientsChance {
			planet.SurfaceFeatures = append(planet.SurfaceFeatures, sfIntelligentLife)
			/* Max tech level is determined by resources */
			if planet.Resources == rscRich {
				planet.NativeTechLevel = random(0, 10)
			} else if planet.Resources == rscPoor {
				planet.NativeTechLevel = random(0, 6)
			} else if planet.Resources == rscNone {
				planet.NativeTechLevel = random(0, 4)
			} else {
				fmt.Println("Unexpected resources value", planet.Resources)
			}
			if planet.NativeTechLevel >= 3 {
				/* Neolithic or higher */
				planet.Anomalies = append(planet.Anomalies, anPossibleStructures)
			}
			if planet.NativeTechLevel >= 8 {
				/* Atomic or higher */
				planet.Anomalies = append(planet.Anomalies, anElectromagneticActivity)
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
		planet.SurfaceFeatures = append(planet.SurfaceFeatures, sfMonumentalRuins)
		if !anomalyInList(anPossibleStructures, planet.Anomalies) {
			planet.Anomalies = append(planet.Anomalies, anPossibleStructures)
		}
	}
	if random(0, 99) < ruinsChance {
		planet.SurfaceFeatures = append(planet.SurfaceFeatures, sfHighTechRuins)
		if !anomalyInList(anPossibleStructures, planet.Anomalies) {
			planet.Anomalies = append(planet.Anomalies, anPossibleStructures)
		}
	}
	if random(0, 99) < ruinsChance {
		planet.SurfaceFeatures = append(planet.SurfaceFeatures, sfDangerousRuins)
		if !anomalyInList(anPossibleStructures, planet.Anomalies) {
			planet.Anomalies = append(planet.Anomalies, anPossibleStructures)
		}
	}
	if random(0, 99) < ruinsChance &&
		!featureInList(sfMonumentalRuins, planet.SurfaceFeatures) &&
		!featureInList(sfHighTechRuins, planet.SurfaceFeatures) &&
		!featureInList(sfDangerousRuins, planet.SurfaceFeatures) {
		planet.SurfaceFeatures = append(planet.SurfaceFeatures, sfRegularGeologicalFormations)
		if !anomalyInList(anPossibleStructures, planet.Anomalies) {
			planet.Anomalies = append(planet.Anomalies, anPossibleStructures)
		}
	}

	name(&planet)

	return &planet
}

func name(p *Planet) {
	backupNames := []string{}
	possibleNames := []string{}

	backupNames = append(backupNames, "This World")

	if p.Atmosphere == atmBreathable &&
		p.Temperature == tmpModerate &&
		p.Gravity == grvModerate &&
		p.Water == wtrOceans &&
		p.Resources == rscRich {
		possibleNames = append(possibleNames, "Eden")
		possibleNames = append(possibleNames, "Paradise")
		possibleNames = append(possibleNames, "Terra Nova")
		possibleNames = append(possibleNames, "New Earth")
	} else {
		switch p.Temperature {
		case tmpVeryHot:
			possibleNames = append(possibleNames, "Inferno")
		case tmpHot:
			backupNames = append(backupNames, "Inferno")
		case tmpVeryCold:
			possibleNames = append(possibleNames, "Arctica")
		case tmpCold:
			backupNames = append(backupNames, "Arctica")
		}

		if p.Water != wtrPlanetWideOcean {
			switch p.Gravity {
			case grvVeryHigh:
				possibleNames = append(possibleNames, "Cueball")
			case grvHigh:
				backupNames = append(backupNames, "Cueball")
			case grvVeryLow:
				possibleNames = append(possibleNames, "Crag")
			case grvLow:
				backupNames = append(backupNames, "Crag")
			}
		}

		switch p.Water {
		case wtrPlanetWideOcean:
			possibleNames = append(possibleNames, "Atlantis")
			possibleNames = append(possibleNames, "Oceanus")
		case wtrIceCoveredSurface:
			possibleNames = append(possibleNames, "Snowball")
			possibleNames = append(possibleNames, "Iceball")
		case wtrTrace, wtrNone:
			possibleNames = append(possibleNames, "Arid")
			possibleNames = append(possibleNames, "Desert")
		}

		if p.Resources == rscRich {
			possibleNames = append(possibleNames, "Bounty")
			possibleNames = append(possibleNames, "El Dorado")
		}

		if featureInList(sfPlantLife, p.SurfaceFeatures) ||
			featureInList(sfEdiblePlants, p.SurfaceFeatures) ||
			featureInList(sfPoisonousPlants, p.SurfaceFeatures) {
			possibleNames = append(possibleNames, "Garden")
			possibleNames = append(possibleNames, "Arcadia")
		}

		if featureInList(sfAirtightCaves, p.SurfaceFeatures) ||
			featureInList(sfInsulatedCaves, p.SurfaceFeatures) {
			possibleNames = append(possibleNames, "Warren")
			possibleNames = append(possibleNames, "Honeycomb")
		}
	}

	/* Pick a name */
	if len(possibleNames) == 0 {
		possibleNames = backupNames
	}

	p.Name = pickOne(possibleNames)
}

func describe(state *SystemState, p *Planet) string {
	/* Now the first impression text */
	s := ""
	if state.colonists >= maxColonists {
		s += "The colonists"
	} else if state.colonists >= maxColonists/2 {
		s += "The surviving colonists"
	} else if state.colonists >= 0 {
		s += "The few surviving colonists"
	}

	s += " wake from their sleep chambers and survey their new home.\n"

	/* Landscape */
	if p.Water == wtrPlanetWideOcean {
		switch p.Gravity {
		case grvVeryLow:
			s += "The seedship tosses on an ocean of slow, kilometres-high waves, beneath"
		case grvLow:
			s += "The seedship rocks on the surface of an ocean of slow, tall waves, beneath"
		case grvModerate:
			s += "The ship floats on the surface of an ocean that stretches to the horizon"
		case grvHigh:
			s += "The ship floats on the surface of a calm ocean that stretches to the horizon"
		case grvVeryHigh:
			s += "The ship rests on the surface of a mirror-flat ocean beneath"
		default:
			fmt.Println("error: Unexpected gravity value $gravity.")
		}
	} else if featureInList(sfPlantLife, p.SurfaceFeatures) ||
		featureInList(sfEdiblePlants, p.SurfaceFeatures) ||
		featureInList(sfPoisonousPlants, p.SurfaceFeatures) {
		switch p.Gravity {
		case grvVeryLow:
			s += "Forests of impossibly slender alien plants reach kilometres into"
		case grvLow:
			s += "Forests of tall alien plants reach hundreds of metres into"
		case grvModerate:
			s += "Forests of alien vegetation stretch away beneath"
		case grvHigh:
			s += "Fields of squat, thick-limbed alien plants stretch away beneath"
		case grvVeryHigh:
			s += "Level fields of alien moss stretch away beneath"
		default:
			fmt.Println("error: Unexpected gravity value $gravity.")
		}
	} else if p.Water == wtrIceCoveredSurface {
		switch p.Gravity {
		case grvVeryLow:
			s += "Kilometres-high spires of ice reach into"
		case grvLow:
			s += "Tall, jagged spires of ice reach into"
		case grvModerate:
			s += "Jagged shards of ice stretch away beneath"
		case grvHigh:
			s += "A crumpled plain of ice stretches away beneath"
		case grvVeryHigh:
			s += "A perfectly flat plain of ice stretches away beneath"
		default:
			fmt.Println("error: Unexpected gravity value $gravity.")
		}
	} else if p.Atmosphere == atmNone {
		switch p.Gravity {
		case grvVeryLow:
			s += "A jagged landscape of high crater walls and towering shards of rock stretches away beneath"
		case grvLow:
			s += "A jagged, cratered landscape stretches away beneath"
		case grvModerate:
			s += "A perfectly still, cratered landscape stretches away beneath"
		case grvHigh:
			s += "A perfectly still landscape, flat except for a few shallow craters, stretches away beneath"
		case grvVeryHigh:
			s += "A perfectly still, flat landscape stretches away beneath"
		default:
			fmt.Println("error: Unexpected gravity value $gravity.")
		}
	} else {
		switch p.Gravity {
		case grvVeryLow:
			s += "A landscape of spindly rock outcroppings and impossibly tall mountains stretches away beneath"
		case grvLow:
			s += "A landscape of huge boulders and towering mountains stretches away beneath"
		case grvModerate:
			s += "A barren, rocky landscape stretches away beneath"
		case grvHigh:
			s += "A flat, rocky landscape stretches away beneath"
		case grvVeryHigh:
			s += "A perfectly flat, barren landscape stretches away beneath"
		default:
			fmt.Println("error: Unexpected gravity value $gravity.")
		}
	}

	/* Sky (come back to this and add more) */
	switch p.Atmosphere {
	case atmBreathable:
		s += " a blue sky."
	case atmMarginal:
		s += " a pale blue sky."
	case atmNone:
		s += " a black, star-studded sky."
	case atmNonBreathable:
		s += " an alien sky."
	case atmToxic:
		s += " a poisonous sky."
	case atmCorrosive:
		s += " a sky filled with corrosive clouds."
	default:
		fmt.Println("error: Unexpected atmosphere value $atmosphere.")
	}

	if state.colonists < maxColonists {
		s += "\nThey build a memorial to the print $max_colonists-$colonists colonists who did not survive the journey, and"
	} else {
		s += "\nThey"
	}

	s += " name their new world " + p.Name + "\n"

	return s
}
