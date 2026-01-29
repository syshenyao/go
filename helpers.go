package utils

import "fmt"

// PrintMessage is a utility function to print messages
func PrintMessage(message string) {
	fmt.Printf("Utility: %s\n", message)
}

// CalculateAverage calculates the average of a slice of numbers
func CalculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	var sum float64
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}
