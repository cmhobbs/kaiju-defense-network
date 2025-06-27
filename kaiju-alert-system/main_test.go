package main

import (
	"testing"
)

var (
	criticalAlertKaiju = Kaiju{Name: "Godzilla", Location: "Tokyo", ThreatLevel: "Critical"}
	highAlertKaiju     = Kaiju{Name: "Megalon", Location: "Syndney", ThreatLevel: "High"}
	lowAlertKaiju      = Kaiju{Name: "Mothra", Location: "New York", ThreatLevel: "Low"}
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
	criticalAlert := formatAlert(criticalAlertKaiju)
	expectedCritical := "WARNING: A CRITICAL LEVEL KAIJU HAS BEEN SPOTTED.  EVACUATE IMMEDIATELY."
	if criticalAlert != expectedCritical {
		t.Errorf("Expected formatAlert for critical threat to return %q, but got %q", expectedCritical, criticalAlert)
	}

	highAlert := formatAlert(highAlertKaiju)
	expectedHigh := "WARNING: A HIGH LEVEL KAIJU HAS BEEN SPOTTED.  PREPARE DEFENSES."
	if highAlert != expectedHigh {
		t.Errorf("Expected formatAlert for high threat to return %q, but got %q", expectedHigh, highAlert)
	}

	// Added for completeness.  Ideally we will not see this level of threat in an alert
	lowAlert := formatAlert(lowAlertKaiju)
	expectedLow := "WARNING: A LOW LEVEL KAIJU HAS BEEN SPOTTED.  MONITOR SITUATION."
	if lowAlert != expectedLow {
		t.Errorf("Expected formatAlert for low threat to return %q, but got %q", expectedLow, lowAlert)
	}
}
