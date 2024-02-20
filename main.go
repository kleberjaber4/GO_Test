package main

import (
	"fmt"
	"strings"
	"time"
)

func isInList(value string, list []string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func formatLicensePlate(input string) string {
	// Verwijder alle streepjes en spaties uit de invoer
	input = strings.ReplaceAll(input, "-", "")
	input = strings.ReplaceAll(input, " ", "")

	// Voeg streepjes toe op de juiste posities
	formatted := fmt.Sprintf("%s-%s-%s", input[:2], input[2:5], input[5:])
	return formatted
}

func main() {

	kentekens := []string{"AB-123-C", "XY-456-Z", "CD-789-E", "FG-012-D", "HI-345-F"}

	var licensePlate string
	fmt.Print("Voer uw kenteken in: ")
	fmt.Scanln(&licensePlate)
	formattedInput := formatLicensePlate(licensePlate)

	licensePlate = strings.ToUpper(formattedInput)
	fmt.Println("Uw ingevulde kenteken:", licensePlate)

	// Check if the input value is in the list
	if isInList(licensePlate, kentekens) {
		currentTime := time.Now()
		hour := currentTime.Hour()

		var greeting string

		switch {
		case hour >= 7 && hour < 12:
			greeting = "Goedemorgen. Welkom bij Fonteyn Vakantieparken!"
		case hour >= 12 && hour < 18:
			greeting = "Goedemiddag. Welkom bij Fonteyn Vakantieparken!"
		case hour >= 18 && hour < 23:
			greeting = "Goedenavond. Welkom bij Fonteyn Vakantieparken!"
		default:
			greeting = "Sorry, de parkeerplaats is ’s nachts gesloten!"
		}
		fmt.Println(greeting)

	} else {
		fmt.Printf("U heeft helaas geen toegang tot het parkeerterrein")
	}

}
