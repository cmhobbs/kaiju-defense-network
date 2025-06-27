package main

import (
	"fmt"
	"math/rand"
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

type KaijuGenerator struct {
	Prefixes     []string
	Suffixes     []string
	Locations    []string
	ThreatLevels []string
	Sizes        []string
	Behaviors    []string
}

func NewKaijuGenerator() *KaijuGenerator {
	return &KaijuGenerator{
		Prefixes: []string{
			"Giga", "Mega", "Ultra", "Apex", "Omega", "Hydro", "Pyro",
			"Electro", "Cryo", "Neo",
		},
		Suffixes: []string{
			"zilla", "tron", "saurus", "moth", "beast", "king", "lord", "demon",
			"dragon", "crusher",
		},
		Locations: []string{
			"New York", "Tokyo", "Paris", "Sydney", "London", "Berlin", "Rome",
			"Moscow", "Beijing", "Rio de Janeiro",
		},
		ThreatLevels: []string{
			"Low", "Medium", "High", "Critical",
		},
		Sizes: []string{
			"Large", "Huge", "Gigantic", "Colossal", "Titan", "Supreme",
		},
		Behaviors: []string{
			"Aggressive", "Defensive", "Ambush", "Patrol", "Scout",
		},
	}
}

// Generate a single kaiju.
func (kg *KaijuGenerator) Generate() Kaiju {
	kaiju := Kaiju{
		Name:        kg.Prefixes[rand.Intn(len(kg.Prefixes))] + kg.Suffixes[rand.Intn(len(kg.Suffixes))],
		Location:    kg.Locations[rand.Intn(len(kg.Locations))],
		ThreatLevel: kg.ThreatLevels[rand.Intn(len(kg.ThreatLevels))],
		Size:        kg.Sizes[rand.Intn(len(kg.Sizes))],
		Behavior:    kg.Behaviors[rand.Intn(len(kg.Behaviors))],
	}
	return kaiju
}

// Generate multiple kaiju sightings.
func generateMultiple(count int) []Sighting {
	var sightings []Sighting
	for i := 0; i < count; i++ {
		generator := NewKaijuGenerator()
		kaiju := generator.Generate()
		sightings = append(sightings, Sighting{
			Kaiju:     kaiju,
			Timestamp: time.Now(),
		})
	}
	return sightings
}

// Print output for a single kaiju sighting to stdout.
func main() {
	rand.Seed(time.Now().UnixNano())

	sighting := generateMultiple(1)
	kaiju := sighting[0].Kaiju
	sightingTime := sighting[0].Timestamp

	fmt.Printf("🚨 KAIJU SIGHTING ALERT AT %s\n", sightingTime.Format(time.DateTime))
	fmt.Printf("────────────────────────────────────────────────────────────\n")
	fmt.Printf("A %s %s has been spotted in %s,\nexhibiting %s behavior!\n\nIt is a %s threat.\n", kaiju.Size, kaiju.Name, kaiju.Location, kaiju.Behavior, kaiju.ThreatLevel)

	fmt.Printf("\n⚠️  SIGHTING DETAILS\n")
	fmt.Printf("────────────────────────────────────────────────────────────\n")
	fmt.Printf("👾 Name: %s\n", kaiju.Name)
	fmt.Printf("📍 Location: %s\n", kaiju.Location)
	fmt.Printf("⚡ Threat Level: %s\n", kaiju.ThreatLevel)
	fmt.Printf("📏 Size: %s\n", kaiju.Size)
	fmt.Printf("🎭 Behavior: %s\n", kaiju.Behavior)
	fmt.Printf("⏰ Timestamp: %s\n", sightingTime.Format(time.DateTime))
	fmt.Printf("────────────────────────────────────────────────────────────\n")
}
