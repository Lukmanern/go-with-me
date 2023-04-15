package main

import "fmt"

func main() {
	// The temperature in Celsius
	celsius := 25.0

	// Convert the temperature to Kelvin
	kelvin := celsius + 273.15
	// Convert the temperature to Fahrenheit
	fahrenheit := celsius*9/5 + 32
	// Convert the temperature to Réaumur
	reaumur := celsius * 4 / 5

	// Print the results
	fmt.Printf("%.2f°C = %.2fK\n", celsius, kelvin)
	fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
	fmt.Printf("%.2f°C = %.2f°Ré\n", celsius, reaumur)
}
