package main

import "fmt"

// ModernChair
type ModernChair struct{}

func (m *ModernChair) SitOn() {
	fmt.Println("Sitting on a modern chair.")
}

// VictorianChair
type VictorianChair struct{}

func (v *VictorianChair) SitOn() {
	fmt.Println("Sitting on a Victorian chair.")
}

// ModernSofa
type ModernSofa struct{}

func (m *ModernSofa) LayOn() {
	fmt.Println("Lying on a modern sofa.")
}

// VictorianSofa
type VictorianSofa struct{}

func (v *VictorianSofa) LayOn() {
	fmt.Println("Lying on a Victorian sofa.")
}

func main() {
	style := "modern" // Change to "victorian" for Victorian style

	var chair interface{}
	var sofa interface{}

	// Directly using constructors based on condition
	if style == "modern" {
		chair = &ModernChair{}
		sofa = &ModernSofa{}
	} else if style == "victorian" {
		chair = &VictorianChair{}
		sofa = &VictorianSofa{}
	} else {
		panic("Unknown style!")
	}

	// Using type assertion to call methods (not type-safe)
	if c, ok := chair.(*ModernChair); ok {
		c.SitOn()
	} else if c, ok := chair.(*VictorianChair); ok {
		c.SitOn()
	}

	if s, ok := sofa.(*ModernSofa); ok {
		s.LayOn()
	} else if s, ok := sofa.(*VictorianSofa); ok {
		s.LayOn()
	}
}
