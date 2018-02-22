package main

import "flag"

var (
	count          int
	newIntegrity   int
	visitedSystems int
)

func init() {
	flag.IntVar(&count, "count", 1, "number of planets to visit")
	flag.Parse()

	visitedSystems = 0
}

func main() {
	gameLoop(count)
}
