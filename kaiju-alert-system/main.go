package main

import (
	"fmt"
	"strings"
)

type Kaiju struct {
	Name        string
	Location    string
	ThreatLevel string
}

// TODO update this to include the new Kaiju fields
func simulateKaijuSighting() Kaiju {
	return Kaiju{
		Name:        "Baragon",
		Location:    "Tokyo",
		ThreatLevel: "High",
	}
}

func shouldAlert(threatLevel string) bool {
	return threatLevel == "High" || threatLevel == "Critical"
}

func formatAlert(kaiju Kaiju) string {
	threatLevel := strings.ToUpper(kaiju.ThreatLevel)
	return fmt.Sprintf("WARNING: A %s LEVEL KAIJU HAS BEEN SPOTTED", threatLevel)
}

func main() {
	kaiju := simulateKaijuSighting()
	if shouldAlert(kaiju.ThreatLevel) {
		alert := formatAlert(kaiju)
		fmt.Println(alert)
	}
}
