package main

type atmosphereType int

const (
	atmBreathable atmosphereType = iota + 1
	atmMarginal
	atmCorrosive
	atmToxic
	atmNonBreathable
	atmNone
)

func (atmosphere atmosphereType) String() string {
	atmospheres := [...]string{
		"uninitialized",
		"Breathable",
		"Marginal",
		"Corrosive",
		"Toxic",
		"Non-breathable",
		"None",
	}
	if atmosphere < atmBreathable || atmosphere > atmNone {
		return "Unknown"
	}
	return atmospheres[atmosphere]
}

type gravityType int

const (
	grvVeryLow gravityType = iota + 1
	grvLow
	grvModerate
	grvHigh
	grvVeryHigh
)

func (gravity gravityType) String() string {
	gravities := [...]string{
		"uninitialized",
		"Very low",
		"Low",
		"Moderate",
		"High",
		"Very high",
	}
	if gravity < grvVeryLow || gravity > grvVeryHigh {
		return "Unknown"
	}
	return gravities[gravity]
}

type temperatureType int

const (
	tmpVeryCold temperatureType = iota + 1
	tmpCold
	tmpModerate
	tmpHot
	tmpVeryHot
)

func (temperature temperatureType) String() string {
	temperatures := [...]string{
		"uninitialized",
		"Very cold",
		"Cold",
		"Moderate",
		"Hot",
		"Very hot",
	}
	if temperature < tmpVeryCold || temperature > tmpVeryHot {
		return "Unknown"
	}
	return temperatures[temperature]
}

type waterType int

const (
	wtrNone waterType = iota + 1
	wtrTrace
	wtrOceans
	wtrPlanetWideOcean
	wtrIceCoveredSurface
	wtrIceCaps
)

func (water waterType) String() string {
	waters := [...]string{
		"uninitialized",
		"None",
		"Trace",
		"Oceans",
		"Planet-wide ocean",
		"Ice-covered surface",
		"Ice caps",
	}
	if water < wtrNone || water > wtrIceCaps {
		return "Unknown"
	}
	return waters[water]
}

type resourcesType int

const (
	rscNone resourcesType = iota + 1
	rscPoor
	rscRich
)

func (resource resourcesType) String() string {
	resources := [...]string{
		"uninitialized",
		"None",
		"Poor",
		"Rich",
	}
	if resource < rscNone || resource > rscRich {
		return "Unknown"
	}
	return resources[resource]
}

type surfaceFeatureType int

type surfaceFeatureList []surfaceFeatureType

const (
	sfBarrenMoon surfaceFeatureType = iota + 1
	sfMetalRichMoon
	sfUnstableMoon
	sfOutstandingBeauty
	sfOutstandingUgliness
	sfAirtightCaves
	sfInsulatedCaves
	sfUnstableGeology
	sfPlantLife
	sfEdiblePlants
	sfPoisonousPlants
	sfAnimalLife
	sfUsefulAnimals
	sfDangerousAnimals
	sfIntelligentLife
	sfMonumentalRuins
	sfHighTechRuins
	sfDangerousRuins
	sfRegularGeologicalFormations
)

func (feature surfaceFeatureType) String() string {
	features := [...]string{
		"uninitialized",
		"Barren moon",
		"Metal-rich moon",
		"Unstable moon",
		"Outstanding beauty",
		"Outstanding ugliness",
		"Airtight caves",
		"Insulated caves",
		"Unstable geology",
		"Plant life",
		"Edible plants",
		"Poisonous plants",
		"Animal life",
		"Useful animals",
		"Dangerous animals",
		"Intelligent life",
		"Monumental ruins",
		"High-tech ruins",
		"Dangerous ruins",
		"Regular geological formations",
	}
	if feature < sfBarrenMoon || feature > sfRegularGeologicalFormations {
		return "unknown"
	}
	return features[feature]
}

func (features surfaceFeatureList) String() string {
	s := ""
	first := true
	for _, v := range features {
		if !first {
			s += ", "
		}
		first = false
		s += v.String()
	}
	return s
}

func featureInList(a surfaceFeatureType, list surfaceFeatureList) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type anomalyType int

type anomalyList []anomalyType

const (
	anMoon anomalyType = iota + 1
	anGeologicalAnomaly
	anVegetation
	anAnimalLife
	anPossibleStructures
	anElectromagneticActivity
)

func (anomaly anomalyType) String() string {
	anomalies := [...]string{
		"uninitialized",
		"Moon",
		"Geological anomalies",
		"Vegetation",
		"Animal life",
		"Possible structures",
		"Electromagnetic activity",
	}
	if anomaly < anMoon || anomaly > anElectromagneticActivity {
		return "unknown"
	}
	return anomalies[anomaly]
}

func (anomalies anomalyList) String() string {
	s := ""
	first := true
	for _, v := range anomalies {
		if !first {
			s += ", "
		}
		first = false
		s += v.String()
	}
	return s
}

func anomalyInList(a anomalyType, list anomalyList) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
