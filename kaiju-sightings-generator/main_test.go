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

	kaiju := generator.Generate()

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
}

func TestGenerateMultiple(t *testing.T) {
	sightings := generateMultiple(10)

	if len(sightings) != 10 {
		t.Errorf("Expected 10 sightings, got %d", len(sightings))
	}

	// TODO make this less brittle
	allSame := true
	for i := 1; i < len(sightings); i++ {
		if sightings[i].Kaiju.Name != sightings[0].Kaiju.Name ||
			sightings[i].Kaiju.Location != sightings[0].Kaiju.Location ||
			sightings[i].Kaiju.ThreatLevel != sightings[0].Kaiju.ThreatLevel ||
			sightings[i].Kaiju.Size != sightings[0].Kaiju.Size ||
			sightings[i].Kaiju.Behavior != sightings[0].Kaiju.Behavior {
			allSame = false
			break
		}
	}

	// Naively test that our monsters are random
	if allSame {
		t.Error("All generated kaiju are identical - randomness may not be working")
	}

	// Test that all sightings have valid timestamps
	for i, sighting := range sightings {
		if sighting.Timestamp.IsZero() {
			t.Errorf("Sighting %d has zero timestamp", i)
		}
		if sighting.Timestamp.Unix() <= 0 {
			t.Errorf("Sighting %d has invalid Unix timestamp: %v", i, sighting.Timestamp.Unix())
		}
	}
}
