// builder_pattern.go: Unified Builder Pattern Implementation

package main

import "fmt"

// iBuilder.go: Builder interface
// Step 1: Define the construction steps clearly for all product representations.
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

// Factory function to return the appropriate builder instance.
func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// normalBuilder.go: Concrete builder
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// Constructor for NormalBuilder
func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

// Define construction steps for NormalBuilder
func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

// Retrieve the constructed product
func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// iglooBuilder.go: Concrete builder
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// Constructor for IglooBuilder
func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

// Define construction steps for IglooBuilder
func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

// Retrieve the constructed product
func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// house.go: Product
// Represents the final product.
type House struct {
	windowType string
	doorType   string
	floor      int
}

// director.go: Director
// Step 4: Create a director class to manage the construction process.
type Director struct {
	builder IBuilder
}

// Constructor for Director
func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// Change the builder used by the director
func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

// Direct the building process
func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// main.go: Client code
// Step 5: Client creates both builder and director and initiates the building process.
func main() {
	// Initialize builders for different types of houses.
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	// Create a director and assign it a builder.
	director := newDirector(normalBuilder)

	// Build a normal house.
	normalHouse := director.buildHouse()

	// Display the details of the normal house.
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	// Switch to the igloo builder and build an igloo house.
	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	// Display the details of the igloo house.
	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}

// Steps Summary:
// 1. Define the construction steps clearly for all product representations in the IBuilder interface.
// 2. Implement concrete builders (NormalBuilder, IglooBuilder) to provide specific product representations.
// 3. Create a product struct (House) to hold the result of the construction.
// 4. Create a Director class to manage the sequence of construction steps.
// 5. Use the client to initiate the construction process by associating a builder with the director.
// 6. Retrieve the constructed product either from the director or the builder.
