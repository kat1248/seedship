package main

import "flag"

var (
	count          int
	newIntegrity   int
	visitedSystems int
	skip           bool
)

func init() {
	flag.IntVar(&count, "count", 1, "number of planets to visit")
	flag.BoolVar(&skip, "skip", false, "skip intro")
	flag.Parse()

	visitedSystems = 0
}

func main() {
	gameLoop(count, skip)
}
