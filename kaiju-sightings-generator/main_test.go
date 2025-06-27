package main

import (
	"math/rand"
	"strings"
	"testing"
)

func TestNewKaijuGenerator(t *testing.T) {
	generator := NewKaijuGenerator()

	if len(generator.Prefixes) == 0 {
		t.Error("Expected prefixes to be populated")
	}
	if len(generator.Suffixes) == 0 {
		t.Error("Expected suffixes to be populated")
	}
	if len(generator.Locations) == 0 {
		t.Error("Expected locations to be populated")
	}
	if len(generator.ThreatLevels) == 0 {
		t.Error("Expected threat levels to be populated")
	}
	if len(generator.Sizes) == 0 {
		t.Error("Expected sizes to be populated")
	}
	if len(generator.Behaviors) == 0 {
		t.Error("Expected behaviors to be populated")
	}
}

// TODO clean this test up, it is doing a lot of things
func TestGenerate(t *testing.T) {
	rand.Seed(42)
	generator := NewKaijuGenerator()

	kaiju, sightingTime := generator.Generate()

	if kaiju.Name == "" {
		t.Error("Expected kaiju name to be generated")
	}

	hasValidPrefix := false
	for _, prefix := range generator.Prefixes {
		if strings.HasPrefix(kaiju.Name, prefix) {
			hasValidPrefix = true
			break
		}
	}
	if !hasValidPrefix {
		t.Errorf("Kaiju name '%s' doesn't start with valid prefix", kaiju.Name)
	}

	validLocation := false
	for _, location := range generator.Locations {
		if kaiju.Location == location {
			validLocation = true
			break
		}
	}
	if !validLocation {
		t.Errorf("Invalid location: %s", kaiju.Location)
	}

	validThreatLevel := false
	for _, level := range generator.ThreatLevels {
		if kaiju.ThreatLevel == level {
			validThreatLevel = true
			break
		}
	}
	if !validThreatLevel {
		t.Errorf("Invalid threat level: %s", kaiju.ThreatLevel)
	}

	validSize := false
	for _, size := range generator.Sizes {
		if kaiju.Size == size {
			validSize = true
			break
		}
	}
	if !validSize {
		t.Errorf("Invalid size: %s", kaiju.Size)
	}

	validBehavior := false
	for _, behavior := range generator.Behaviors {
		if kaiju.Behavior == behavior {
			validBehavior = true
			break
		}
	}
	if !validBehavior {
		t.Errorf("Invalid behavior: %s", kaiju.Behavior)
	}

	// test that sightingTime is actually a time type
	if sightingTime.Unix() <= 0 {
		t.Errorf("Sighting time should have a valid Unix timestamp: %v", sightingTime.Unix())
	}

	// test that sightingTime has proper date components
	year, month, day := sightingTime.Date()
	if year < 2000 || int(month) < 1 || int(month) > 12 || day < 1 || day > 31 {
		t.Errorf("Sighting time has invalid date components: year=%d, month=%d, day=%d", year, month, day)
	}

	// test that sightingTime has proper time components
	hour, min, sec := sightingTime.Clock()
	if hour < 0 || hour > 23 || min < 0 || min > 59 || sec < 0 || sec > 59 {
		t.Errorf("Sighting time has invalid time components: hour=%d, min=%d, sec=%d", hour, min, sec)
	}
}

func TestGenerateMultiple(t *testing.T) {
	kaijus := generateMultiple(10)

	allSame := true
	for i := 1; i < len(kaijus); i++ {
		if kaijus[i].Name != kaijus[0].Name ||
			kaijus[i].Location != kaijus[0].Location ||
			kaijus[i].ThreatLevel != kaijus[0].ThreatLevel ||
			kaijus[i].Size != kaijus[0].Size ||
			kaijus[i].Behavior != kaijus[0].Behavior {
			allSame = false
			break
		}
	}

	if allSame {
		t.Error("All generated kaiju are identical - randomness may not be working")
	}
}
