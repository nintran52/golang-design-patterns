package main

import "fmt"

// Truck struct
type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Delivering by truck.")
}

func PlanDelivery(truck Truck) {
	truck.Deliver()
}

func main() {
	// Create a logistics app that only handles trucks
	truck := Truck{}

	// Plan a delivery
	PlanDelivery(truck)

	// Now we need to add support for Ships
	// This would require rewriting significant portions of the app.
}
