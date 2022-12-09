package main

import "fmt"

// Planet represents a celestial body in our solar system
type Planet struct {
	Name       string
	Mass       float64
	Diameter   float64
	Distance   float64
	NumMoons   int
	Rings      bool
}

func main() {
	// Create a slice of planets
	planets := []Planet{
		{
			Name:     "Mercury",
			Mass:     3.285e23,
			Diameter: 4879,
			Distance: 57.9e6,
			NumMoons: 0,
			Rings:    false,
		},
		{
			Name:     "Venus",
			Mass:     4.867e24,
			Diameter: 12104,
			Distance: 108.2e6,
			NumMoons: 0,
			Rings:    false,
		},
		// ...
	}

	// Print the name and distance of each planet
	for _, planet := range planets {
		fmt.Printf("%s is %.1f million km away from the Sun\n", 
			planet.Name, planet.Distance/1e6)
	}
}