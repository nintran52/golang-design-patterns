package main

import "fmt"

// Product interface
type Transport interface {
	Deliver()
}

// Concrete Products
type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Delivering by land using a truck.")
}

type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Delivering by sea using a ship.")
}

// Creator Interface
type Logistics interface {
	CreateTransport() Transport
	PlanDelivery()
}

type BaseLogistics struct{}

func (b *BaseLogistics) PlanDelivery() {
	transport := b.CreateTransport()
	transport.Deliver()
}

func (b *BaseLogistics) CreateTransport() Transport {
	return nil
}

// Concrete Creators
type RoadLogistics struct {
	BaseLogistics
}

func (r *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

// SeaLogistics struct
type SeaLogistics struct {
	BaseLogistics
}

func (s *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

// Client code
func main() {
	// Create a road logistics instance
	roadLogistics := &RoadLogistics{}
	fmt.Println("Road Logistics Delivery:")
	roadLogistics.PlanDelivery()

	// Create a sea logistics instance
	seaLogistics := &SeaLogistics{}
	fmt.Println("Sea Logistics Delivery:")
	seaLogistics.PlanDelivery()
}
