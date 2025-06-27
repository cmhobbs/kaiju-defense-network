package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func (kg *KaijuGenerator) Generate() Kaiju {
	return Kaiju{
		Name:        kg.Prefixes[rand.Intn(len(kg.Prefixes))] + kg.Suffixes[rand.Intn(len(kg.Suffixes))],
		Location:    kg.Locations[rand.Intn(len(kg.Locations))],
		ThreatLevel: kg.ThreatLevels[rand.Intn(len(kg.ThreatLevels))],
		Size:        kg.Sizes[rand.Intn(len(kg.Sizes))],
		Behavior:    kg.Behaviors[rand.Intn(len(kg.Behaviors))],
	}
}

func generateMultiple(count int) []Kaiju {
	var kaijus []Kaiju
	for i := 0; i < count; i++ {
		generator := NewKaijuGenerator()
		kaijus = append(kaijus, generator.Generate())
	}
	return kaijus
}

// Show output for a single kaiju sighting.
func main() {
	rand.Seed(time.Now().UnixNano())

	generator := NewKaijuGenerator()
	kaiju := generator.Generate()

	fmt.Printf("🚨 KAIJU SIGHTING ALERT AT %s\n", time.Now().Format(time.DateTime))
	fmt.Printf("────────────────────────────────────────────────────────────\n")
	fmt.Printf("A %s %s has been spotted in %s,\nexhibiting %s behavior!\n\nIt is a %s threat.\n", kaiju.Size, kaiju.Name, kaiju.Location, kaiju.Behavior, kaiju.ThreatLevel)

	fmt.Printf("\n⚠️  SIGHTING DETAILS\n")
	fmt.Printf("────────────────────────────────────────────────────────────\n")
	fmt.Printf("👾 Name: %s\n", kaiju.Name)
	fmt.Printf("📍 Location: %s\n", kaiju.Location)
	fmt.Printf("⚡ Threat Level: %s\n", kaiju.ThreatLevel)
	fmt.Printf("📏 Size: %s\n", kaiju.Size)
	fmt.Printf("🎭 Behavior: %s\n", kaiju.Behavior)
	fmt.Printf("────────────────────────────────────────────────────────────\n")
}
