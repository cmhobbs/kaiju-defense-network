package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Sighting struct {
	Kaiju     Kaiju
	Timestamp time.Time
}

type Kaiju struct {
	Name        string
	Location    string
	ThreatLevel string
	Size        string
	Behavior    string
}

// TODO update this to include the new Kaiju fields
func simulateKaijuSighting() Sighting {
	kaiju := Kaiju{
		Name:        "Baragon",
		Location:    "Tokyo",
		ThreatLevel: "High",
	}

	sighting := Sighting{
		Kaiju:     kaiju,
		Timestamp: time.Now(),
	}

	return sighting
}

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

func formatAlert(sighting Sighting) string {
	kaiju := sighting.Kaiju
	threatLevel := strings.ToUpper(kaiju.ThreatLevel)
	action := requiredAction(kaiju.ThreatLevel)

	return fmt.Sprintf("%s - WARNING: A %s LEVEL KAIJU HAS BEEN SPOTTED!  %s.", sighting.Timestamp.Format(time.DateTime), threatLevel, action)
}

func main() {
	sighting := simulateKaijuSighting()
	kaiju := sighting.Kaiju

	if shouldAlert(kaiju.ThreatLevel) {
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
	}
}
