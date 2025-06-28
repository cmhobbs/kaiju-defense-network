package main

import (
	"fmt"
	kaijuSightingsGenerator "kaiju-sightings-generator"
	"log"
	"os"
	"strings"
	"time"
)

func shouldAlert(threatLevel string) bool {
	return threatLevel == "High" || threatLevel == "Critical"
}

func requiredAction(threatLevel string) string {
	switch threatLevel {
	case "Critical":
		return "EVACUATE IMMEDIATELY"
	case "High":
		return "PREPARE DEFENSES"
	case "Medium":
		return "SHELTER IN PLACE"
	default:
		return "MONITOR SITUATION"
	}
}

func formatAlert(sighting kaijuSightingsGenerator.Sighting) string {
	kaiju := sighting.Kaiju
	threatLevel := strings.ToUpper(kaiju.ThreatLevel)
	action := requiredAction(kaiju.ThreatLevel)

	return fmt.Sprintf("%s - WARNING: A %s LEVEL KAIJU HAS BEEN SPOTTED!  %s.", sighting.Timestamp.Format(time.DateTime), threatLevel, action)
}

func main() {
	sightings := kaijuSightingsGenerator.GenerateMultiple(5)

	for _, sighting := range sightings {
		kaijuSightingsGenerator.PrintSingleSighting(sighting)
		kaiju := sighting.Kaiju

		if shouldAlert(kaiju.ThreatLevel) {
			fmt.Println("Threshold hit, logging alert...")
			fmt.Println()

			alert := formatAlert(sighting)

			file, err := os.OpenFile("alert.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			_, err = file.WriteString(alert + "\n")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Threshold too low, alert will not be logged...")
			fmt.Println()
		}
	}
}
