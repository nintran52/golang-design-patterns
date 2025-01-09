// single.go: Singleton
package main

import (
	"fmt"
	"sync"
)

// Step 1: Add a static field to store the Singleton instance.
var lock = &sync.Mutex{}   // Mutex to handle concurrent access.
var singleInstance *single // Static field to store the Singleton instance.

type single struct {
	// Private fields can be added here if needed.
}

// Step 2: Declare a public static creation method to get the Singleton instance.
func getInstance() *single {
	// Step 3: Implement lazy initialization in the static method.
	if singleInstance == nil {
		lock.Lock() // Acquire lock for thread safety.
		defer lock.Unlock()

		if singleInstance == nil { // Double-check to prevent multiple instances.
			fmt.Println("Creating single instance now.")
			singleInstance = &single{} // Create a new instance.
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance // Return the Singleton instance.
}

// main.go: Client code
func main() {
	// Step 5: Use the static method instead of directly calling the constructor.
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Prevent the program from exiting immediately.
	fmt.Scanln()
}

/*
Summary of Steps to Implement Singleton:
1. Add a static field (`singleInstance`) to store the Singleton instance.
2. Declare a static method (`getInstance`) to return the instance.
3. Implement lazy initialization in the static method to ensure the instance is created only when first requested.
4. Use locking (`sync.Mutex`) and double-checking to handle concurrency and ensure thread safety.
5. Replace direct calls to the constructor with calls to the static method (`getInstance`) in client code.
*/
