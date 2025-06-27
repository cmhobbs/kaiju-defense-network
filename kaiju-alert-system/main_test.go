package main

import (
	"testing"
)

var (
	highAlertKaiju     = Kaiju{Name: "Megalon", Location: "Syndney", ThreatLevel: "High"}
	criticalAlertKaiju = Kaiju{Name: "Godzilla", Location: "Tokyo", ThreatLevel: "Critical"}
	lowAlertKaiju      = Kaiju{Name: "Mothra", Location: "New York", ThreatLevel: "Low"}
)

func TestShouldAlert(t *testing.T) {

	if !shouldAlert(highAlertKaiju.ThreatLevel) {
		t.Errorf("Expected shouldAlert to return true for High threat level, but got false")
	}

	if !shouldAlert(criticalAlertKaiju.ThreatLevel) {
		t.Errorf("Expected shouldAlert to return true for Critical threat level, but got false")
	}
}

func TestShouldNotAlert(t *testing.T) {
	if shouldAlert(lowAlertKaiju.ThreatLevel) {
		t.Errorf("Expected shouldAlert to return false for Low threat level, but got true")
	}
}

func TestFormatAlert(t *testing.T) {
	highAlert := formatAlert(highAlertKaiju)
	expectedHigh := "WARNING: A HIGH LEVEL KAIJU HAS BEEN SPOTTED"
	if highAlert != expectedHigh {
		t.Errorf("Expected formatAlert for high threat to return %q, but got %q", expectedHigh, highAlert)
	}

	criticalAlert := formatAlert(criticalAlertKaiju)
	expectedCritical := "WARNING: A CRITICAL LEVEL KAIJU HAS BEEN SPOTTED"
	if criticalAlert != expectedCritical {
		t.Errorf("Expected formatAlert for critical threat to return %q, but got %q", expectedCritical, criticalAlert)
	}

	// Added for completeness.  Ideally we will not see this level of threat in an alert
	lowAlert := formatAlert(lowAlertKaiju)
	expectedLow := "WARNING: A LOW LEVEL KAIJU HAS BEEN SPOTTED"
	if lowAlert != expectedLow {
		t.Errorf("Expected formatAlert for low threat to return %q, but got %q", expectedLow, lowAlert)
	}
}
