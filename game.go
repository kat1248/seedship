package main

import (
	"fmt"
	"strings"
)

func gameLoop() {
	systems := newSystemState()
	gameIntro()
	for !done(systems) {
		planet := generatePlanet(systems)
		fmt.Println("planet =", planet)
	}
}

func pause() {
	// put logic to delay
}

func visited() int {
	// FIXME, track number of times
	return 1
}

func done(systems *systemState) bool {
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

type planetData struct {
	name            string
	temperature     string
	gravity         string
	resources       string
	atmosphere      string
	water           string
	nativeTechLevel int
	surfaceFeatures []string
	anomalies       []string
}

func (p planetData) String() string {
	s := p.name + "\n" +
		"temperature: " + p.temperature + "\n" +
		"gravity: " + p.gravity + "\n" +
		"resources: " + p.resources + "\n" +
		"atmosphere: " + p.atmosphere + "\n" +
		"water: " + p.water + "\n" +
		"natives: " + fmt.Sprint(p.nativeTechLevel)
	if len(p.surfaceFeatures) > 0 {
		s += "\n" + "surface features: " + strings.Join(p.surfaceFeatures, ", ")
	}
	if len(p.anomalies) > 0 {
		s += "\n" + "anomalies: " + strings.Join(p.anomalies, ", ")
	}

	return s
}

func generatePlanet(systems *systemState) *planetData {
	planet := planetData{
		name:            "fred",
		temperature:     "",
		gravity:         "",
		resources:       "",
		atmosphere:      "",
		water:           "",
		nativeTechLevel: 0,
	}
	systems.surfaceProbeUsed = false
	scannerAtmosphereSuccess := random(0, 99) < systems.atmosphere.strength
	scannerGravitySuccess := random(0, 99) < systems.gravity.strength
	scannerTemperatureSuccess := random(0, 99) < systems.temperature.strength
	scannerResourcesSuccess := random(0, 99) < systems.resources.strength
	scannerWaterSuccess := random(0, 99) < systems.water.strength

	greenChance := 2
	yellowChance := 3
	redChance := 5

	var r int

	/* ATMOSPHERE */
	if visited() == 1 {
		r = random(greenChance, greenChance+yellowChance+redChance-1)
	} else if !scannerAtmosphereSuccess || systems.atmosphere.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.atmosphere.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.atmosphere.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner atmosphere level was", systems.atmosphere.level)
	}

	if r < greenChance {
		planet.atmosphere = "Breathable"
	} else if r < greenChance+yellowChance {
		planet.atmosphere = "Marginal"
	} else {
		planet.atmosphere = either("Corrosive", "Toxic", "Non-breathable", "None")
	}

	/* GRAVITY */
	if visited() == 1 {
		r = random(greenChance, greenChance+yellowChance+redChance-1)
	} else if !scannerGravitySuccess || systems.gravity.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.gravity.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.gravity.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner gravity level was", systems.gravity.level)
	}

	if r < greenChance {
		planet.gravity = "Moderate"
	} else if r < greenChance+yellowChance {
		planet.gravity = either("Low", "High")
	} else {
		planet.gravity = either("Very low", "Very high")
	}

	/* TEMPERATURE */
	if visited() == 1 {
		r = random(greenChance, greenChance+yellowChance+redChance-1)
	} else if !scannerTemperatureSuccess || systems.temperature.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.temperature.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.temperature.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner temperature level was", systems.temperature.level)
	}

	if r < greenChance {
		planet.temperature = "Moderate"
	} else if r < greenChance+yellowChance {
		planet.temperature = either("Cold", "Hot")
	} else {
		planet.temperature = either("Very cold", "Very hot")
	}

	/* WATER */
	if visited() == 1 {
		r = random(greenChance, greenChance+yellowChance+redChance-1)
	} else if !scannerWaterSuccess || systems.water.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.water.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.water.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner water level was", systems.water.level)
	}

	if r < greenChance {
		planet.water = "Oceans"
	} else if r < greenChance+yellowChance {
		planet.water = "Planet-wide ocean"
	} else {
		planet.water = either("Trace", "None")
	}
	/* RESOURCES */
	if visited() == 1 {
		r = random(1, greenChance+yellowChance+redChance-1)
	} else if !scannerResourcesSuccess || systems.resources.level == 0 || systems.offCourse {
		r = random(0, greenChance+yellowChance+redChance-1)
	} else if systems.resources.level == 1 {
		r = random(0, greenChance+yellowChance-1)
	} else if systems.resources.level == 2 {
		r = 0
	} else {
		fmt.Println("Scanner resources level was", systems.resources.level)
	}

	if r < greenChance {
		planet.resources = "Rich"
	} else if r < greenChance+yellowChance {
		planet.resources = "Poor"
	} else {
		planet.resources = "None"
	}

	/* Freeze the oceans at low temperatures */
	if planet.temperature == "Very cold" || (planet.temperature == "Cold" && random(0, 1) == 0) {
		if planet.water == "Oceans" {
			planet.water = "Ice caps"
		} else if planet.water == "Planet-wide ocean" {
			planet.water = "Ice-covered surface"
		}
	}

	/* No liquid water without atmosphere */
	if (planet.water == "Oceans" || planet.water == "Planet-wide ocean") && planet.atmosphere == "None" {
		planet.atmosphere = "Non-breathable"
	}

	/* SURFACE FEATURES */
	planet.surfaceFeatures = []string{}
	planet.anomalies = []string{}

	/* Moon - affects technology result */
	/* Moon is first, because the surface probe investigates it before landing on the planet itself */
	/* The higher the planet's gravity, the more likely it is to have a moon */
	moonChance := 0
	switch planet.gravity {
	case "Very low":
		moonChance = 10
	case "Low":
		moonChance = 20
	case "Moderate":
		moonChance = 30
	case "High":
		moonChance = 40
	case "Very high":
		moonChance = 50
	}

	if random(0, 99) < moonChance {
		planet.surfaceFeatures = append(planet.surfaceFeatures, either("Barren moon", "Metal-rich moon", "Unstable moon"))
		planet.anomalies = append(planet.anomalies, "Moon")
	}

	/* Aesthetics - affect culture result */
	/* Flat 20% chance of one or the other */
	if random(0, 4) == 0 {
		planet.surfaceFeatures = append(planet.surfaceFeatures, either("Outstanding beauty", "Outstanding ugliness"))
	}

	/* Caves? */
	/* No caves if surface is covered entirely in water or ice. Otherwise chance of caves is based on gravity. */
	cavesChance := 0
	if planet.water != "Planet-wide ocean" && planet.water != "Ice-covered surface" {
		switch planet.gravity {
		case "Very low":
			cavesChance = 50
		case "Low":
			cavesChance = 40
		case "Moderate":
			cavesChance = 30
		case "High":
			cavesChance = 20
		case "Very high":
			cavesChance = 10
		}
		if random(0, 99) < cavesChance {
			if random(0, 2) == 0 {
				planet.surfaceFeatures = append(planet.surfaceFeatures, "Airtight caves")
			}
			if random(0, 2) == 0 {
				planet.surfaceFeatures = append(planet.surfaceFeatures, "Insulated caves")
			}
			if !stringInSlice("Airtight caves", planet.surfaceFeatures) &&
				!stringInSlice("Insulated caves", planet.surfaceFeatures) {
				planet.surfaceFeatures = append(planet.surfaceFeatures, "Unstable geology")
			}
			planet.anomalies = append(planet.anomalies, "Geological anomalies")
		}
	}

	/* LIFE */
	/* Plant life? */
	plantsChance := 0
	if true {
		switch planet.atmosphere {
		case "Breathable":
			plantsChance = 100
		case "Marginal":
			plantsChance = 100
		case "Non-breathable":
			plantsChance = 25
		case "Toxic":
			plantsChance = 25
		case "None":
			plantsChance = 5
		case "Corrosive":
			plantsChance = 5
		}
		if random(0, 99) < plantsChance {
			planet.surfaceFeatures = append(planet.surfaceFeatures, either("Plant life", "Edible plants", "Poisonous plants"))
			planet.anomalies = append(planet.anomalies, "Vegetation")
		}
	}
	animalsChance := 0
	/* If plants, possibly animals */
	if stringInSlice("Plant life", planet.surfaceFeatures) ||
		stringInSlice("Edible plants", planet.surfaceFeatures) ||
		stringInSlice("Poisonous plants", planet.surfaceFeatures) {
		animalsChance = 50
		if random(0, 99) < animalsChance {
			planet.surfaceFeatures = append(planet.surfaceFeatures, either("Animal life", "Useful animals", "Dangerous animals"))
			planet.anomalies = append(planet.anomalies, "Animal life")
		}
	}

	planet.nativeTechLevel = 0
	sentientsChance := 0
	/* If animals, possibly sentient life */
	if stringInSlice("Animal life", planet.surfaceFeatures) ||
		stringInSlice("Useful animals", planet.surfaceFeatures) ||
		stringInSlice("Dangerous animals", planet.surfaceFeatures) {
		sentientsChance = 50
		if random(0, 99) < sentientsChance {
			planet.surfaceFeatures = append(planet.surfaceFeatures, "Intelligent life")
			/* Max tech level is determined by resources */
			if planet.resources == "Rich" {
				planet.nativeTechLevel = random(0, 10)
			} else if planet.resources == "Poor" {
				planet.nativeTechLevel = random(0, 6)
			} else if planet.resources == "None" {
				planet.nativeTechLevel = random(0, 4)
			} else {
				fmt.Println("Unexpected resources value", planet.resources)
			}
			if planet.nativeTechLevel >= 3 {
				/* Neolithic or higher */
				planet.anomalies = append(planet.anomalies, "Possible structures")
			}
			if planet.nativeTechLevel >= 8 {
				/* Atomic or higher */
				planet.anomalies = append(planet.anomalies, "Electromagnetic activity")
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
		planet.surfaceFeatures = append(planet.surfaceFeatures, "Monumental ruins")
		if !stringInSlice("Possible structures", planet.anomalies) {
			planet.anomalies = append(planet.anomalies, "Possible structures")
		}
	}
	if random(0, 99) < ruinsChance {
		planet.surfaceFeatures = append(planet.surfaceFeatures, "High-tech ruins")
		if !stringInSlice("Possible structures", planet.anomalies) {
			planet.anomalies = append(planet.anomalies, "Possible structures")
		}
	}
	if random(0, 99) < ruinsChance {
		planet.surfaceFeatures = append(planet.surfaceFeatures, "Dangerous ruins")
		if !stringInSlice("Possible structures", planet.anomalies) {
			planet.anomalies = append(planet.anomalies, "Possible structures")
		}
	}
	if random(0, 99) < ruinsChance &&
		!stringInSlice("Monumental ruins", planet.surfaceFeatures) &&
		!stringInSlice("High-tech ruins", planet.surfaceFeatures) &&
		!stringInSlice("Dangerous ruins", planet.surfaceFeatures) {
		planet.surfaceFeatures = append(planet.surfaceFeatures, "Regular geological formations")
		if !stringInSlice("Possible structures", planet.anomalies) {
			planet.anomalies = append(planet.anomalies, "Possible structures")
		}
	}
	// [[Orbit planet]]

	return &planet
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
