package main

import "fmt"

// --- Abstract Product Interfaces ---

// Chair interface
type Chair interface {
	SitOn()
	GetDescription() string
}

// Sofa interface
type Sofa interface {
	LayOn()
	GetDescription() string
}

// CoffeeTable interface
type CoffeeTable interface {
	PlaceItems()
	GetDescription() string
}

// --- Concrete Products for Modern Style ---

// ModernChair
type ModernChair struct{}

func (m *ModernChair) SitOn() {
	fmt.Println("Sitting on a modern chair.")
}

func (m *ModernChair) GetDescription() string {
	return "This is a modern chair."
}

// ModernSofa
type ModernSofa struct{}

func (m *ModernSofa) LayOn() {
	fmt.Println("Lying on a modern sofa.")
}

func (m *ModernSofa) GetDescription() string {
	return "This is a modern sofa."
}

// ModernCoffeeTable
type ModernCoffeeTable struct{}

func (m *ModernCoffeeTable) PlaceItems() {
	fmt.Println("Placing items on a modern coffee table.")
}

func (m *ModernCoffeeTable) GetDescription() string {
	return "This is a modern coffee table."
}

// --- Concrete Products for Victorian Style ---

// VictorianChair
type VictorianChair struct{}

func (v *VictorianChair) SitOn() {
	fmt.Println("Sitting on a Victorian chair.")
}

func (v *VictorianChair) GetDescription() string {
	return "This is a Victorian chair."
}

// VictorianSofa
type VictorianSofa struct{}

func (v *VictorianSofa) LayOn() {
	fmt.Println("Lying on a Victorian sofa.")
}

func (v *VictorianSofa) GetDescription() string {
	return "This is a Victorian sofa."
}

// VictorianCoffeeTable
type VictorianCoffeeTable struct{}

func (v *VictorianCoffeeTable) PlaceItems() {
	fmt.Println("Placing items on a Victorian coffee table.")
}

func (v *VictorianCoffeeTable) GetDescription() string {
	return "This is a Victorian coffee table."
}

// --- Abstract Factory Interface ---

// FurnitureFactory
type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
	CreateCoffeeTable() CoffeeTable
}

// --- Concrete Factories ---

// ModernFurnitureFactory
type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (m *ModernFurnitureFactory) CreateSofa() Sofa {
	return &ModernSofa{}
}

func (m *ModernFurnitureFactory) CreateCoffeeTable() CoffeeTable {
	return &ModernCoffeeTable{}
}

// VictorianFurnitureFactory
type VictorianFurnitureFactory struct{}

func (v *VictorianFurnitureFactory) CreateChair() Chair {
	return &VictorianChair{}
}

func (v *VictorianFurnitureFactory) CreateSofa() Sofa {
	return &VictorianSofa{}
}

func (v *VictorianFurnitureFactory) CreateCoffeeTable() CoffeeTable {
	return &VictorianCoffeeTable{}
}

// --- Factory Initialization ---

func GetFactory(style string) FurnitureFactory {
	switch style {
	case "modern":
		return &ModernFurnitureFactory{}
	case "victorian":
		return &VictorianFurnitureFactory{}
	default:
		panic("Unknown style!")
	}
}

// --- Client Code ---

func main() {
	// Example: Initialize the factory based on configuration
	style := "modern" // Change to "victorian" for Victorian style
	factory := GetFactory(style)

	// Use the factory to create products
	chair := factory.CreateChair()
	sofa := factory.CreateSofa()
	table := factory.CreateCoffeeTable()

	// Interact with the products
	chair.SitOn()
	fmt.Println(chair.GetDescription())

	sofa.LayOn()
	fmt.Println(sofa.GetDescription())

	table.PlaceItems()
	fmt.Println(table.GetDescription())
}
