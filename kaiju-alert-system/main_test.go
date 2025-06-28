package main

import (
	kaijuSightingsGenerator "kaiju-sightings-generator"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	criticalAlertKaiju = kaijuSightingsGenerator.Kaiju{Name: "Godzilla", Location: "Tokyo", ThreatLevel: "Critical"}
	highAlertKaiju     = kaijuSightingsGenerator.Kaiju{Name: "Megalon", Location: "Syndney", ThreatLevel: "High"}
	lowAlertKaiju      = kaijuSightingsGenerator.Kaiju{Name: "Mothra", Location: "New York", ThreatLevel: "Low"}
)

func TestShouldAlert(t *testing.T) {

	if !shouldAlert(criticalAlertKaiju.ThreatLevel) {
		t.Errorf("Expected shouldAlert to return true for Critical threat level, but got false")
	}

	if !shouldAlert(highAlertKaiju.ThreatLevel) {
		t.Errorf("Expected shouldAlert to return true for High threat level, but got false")
	}

}

func TestShouldNotAlert(t *testing.T) {
	if shouldAlert(lowAlertKaiju.ThreatLevel) {
		t.Errorf("Expected shouldAlert to return false for Low threat level, but got true")
	}
}

func TestFormatAlert(t *testing.T) {
	testTime := time.Date(2024, 1, 15, 14, 30, 0, 0, time.UTC)

	criticalSighting := kaijuSightingsGenerator.Sighting{Kaiju: criticalAlertKaiju, Timestamp: testTime}
	criticalAlert := formatAlert(criticalSighting)
	if !strings.Contains(criticalAlert, "2024-01-15 14:30:00") {
		t.Errorf("Expected formatAlert to include timestamp at beginning, but got %q", criticalAlert)
	}
	if !strings.Contains(criticalAlert, "WARNING: A CRITICAL LEVEL KAIJU HAS BEEN SPOTTED!  EVACUATE IMMEDIATELY.") {
		t.Errorf("Expected formatAlert to contain critical alert message, but got %q", criticalAlert)
	}

	highSighting := kaijuSightingsGenerator.Sighting{Kaiju: highAlertKaiju, Timestamp: testTime}
	highAlert := formatAlert(highSighting)
	if !strings.Contains(highAlert, "2024-01-15 14:30:00") {
		t.Errorf("Expected formatAlert to include timestamp at beginning, but got %q", highAlert)
	}
	if !strings.Contains(highAlert, "WARNING: A HIGH LEVEL KAIJU HAS BEEN SPOTTED!  PREPARE DEFENSES.") {
		t.Errorf("Expected formatAlert to contain high alert message, but got %q", highAlert)
	}

	// Added for completeness.  Ideally we will not see this level of threat in an alert
	lowSighting := kaijuSightingsGenerator.Sighting{Kaiju: lowAlertKaiju, Timestamp: testTime}
	lowAlert := formatAlert(lowSighting)
	if !strings.Contains(lowAlert, "2024-01-15 14:30:00") {
		t.Errorf("Expected formatAlert to include timestamp at beginning, but got %q", lowAlert)
	}
	if !strings.Contains(lowAlert, "WARNING: A LOW LEVEL KAIJU HAS BEEN SPOTTED!  MONITOR SITUATION.") {
		t.Errorf("Expected formatAlert to contain low alert message, but got %q", lowAlert)
	}
}

// NOTE this will nuke your log file, which doesn't matter for a silly project like this but
// production ready code would need a better way to test this.
func TestMainWritesToLogFile(t *testing.T) {
	os.Remove("alert.log")
	main()

	if _, err := os.Stat("alert.log"); os.IsNotExist(err) {
		t.Errorf("Expected alert.log file to be created, but it doesn't exist")
	}

	os.Remove("alert.log")
}
