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
}

type KaijuGenerator struct {
	Prefixes     []string
	Suffixes     []string
	Locations    []string
	ThreatLevels []string
}

func NewKaijuGenerator() *KaijuGenerator {
	return &KaijuGenerator{
		Prefixes: []string{
			"Giga", "Mega", "Ultra", "Titan", "Apex", "Omega", "Hydro", "Pyro",
			"Cryo", "Neo",
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
	}
}

func (kg *KaijuGenerator) Generate() Kaiju {
	return Kaiju{
		Name:        kg.Prefixes[rand.Intn(len(kg.Prefixes))] + kg.Suffixes[rand.Intn(len(kg.Suffixes))],
		Location:    kg.Locations[rand.Intn(len(kg.Locations))],
		ThreatLevel: kg.ThreatLevels[rand.Intn(len(kg.ThreatLevels))],
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	generator := NewKaijuGenerator()
	kaiju := generator.Generate()

	fmt.Printf("🚨 KAIJU SIGHTING ALERT 🚨\n")
	fmt.Printf("Name: %s\n", kaiju.Name)
	fmt.Printf("Location: %s\n", kaiju.Location)
	fmt.Printf("Threat Level: %s\n", kaiju.ThreatLevel)
}
