package main

type Atmosphere int

const (
	atmBreathable Atmosphere = iota
	atmMarginal
	atmCorrosive
	atmToxic
	atmNonBreathable
	atmNone
)

func (atmosphere Atmosphere) String() string {
	atmospheres := [...]string{
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

type Gravity int

const (
	grvVeryLow Gravity = iota
	grvLow
	grvModerate
	grvHigh
	grvVeryHigh
)

func (gravity Gravity) String() string {
	gravities := [...]string{
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

type Temperature int

const (
	tmpVeryCold Temperature = iota
	tmpCold
	tmpModerate
	tmpHot
	tmpVeryHot
)

func (temperature Temperature) String() string {
	temperatures := [...]string{
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

type Water int

const (
	wtrNone Water = iota
	wtrTrace
	wtrOceans
	wtrPlanetWideOcean
	wtrIceCoveredSurface
	wtrIceCaps
)

func (water Water) String() string {
	waters := [...]string{
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

type Resources int

const (
	rscNone Resources = iota
	rscPoor
	rscRich
)

func (resource Resources) String() string {
	resources := [...]string{
		"None",
		"Poor",
		"Rich",
	}
	if resource < rscNone || resource > rscRich {
		return "Unknown"
	}
	return resources[resource]
}

type SurfaceFeature int

type SurfaceFeatureList []SurfaceFeature

const (
	sfBarrenMoon SurfaceFeature = iota
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

func (feature SurfaceFeature) String() string {
	features := [...]string{
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

func (features SurfaceFeatureList) String() string {
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

func featureInList(a SurfaceFeature, list SurfaceFeatureList) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type Anomaly int

type AnomalyList []Anomaly

const (
	anMoon Anomaly = iota
	anGeologicalAnomaly
	anVegetation
	anAnimalLife
	anPossibleStructures
	anElectromagneticActivity
)

func (anomaly Anomaly) String() string {
	anomalies := [...]string{
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

func (anomalies AnomalyList) String() string {
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

func anomalyInList(a Anomaly, list AnomalyList) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
