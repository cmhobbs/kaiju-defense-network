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
}

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
}

func TestGenerateMultiple(t *testing.T) {
	generator := NewKaijuGenerator()

	kaijus := make([]Kaiju, 10)
	for i := 0; i < 10; i++ {
		kaijus[i] = generator.Generate()
	}

	allSame := true
	for i := 1; i < len(kaijus); i++ {
		if kaijus[i].Name != kaijus[0].Name ||
			kaijus[i].Location != kaijus[0].Location ||
			kaijus[i].ThreatLevel != kaijus[0].ThreatLevel {
			allSame = false
			break
		}
	}

	if allSame {
		t.Error("All generated kaiju are identical - randomness may not be working")
	}
}
