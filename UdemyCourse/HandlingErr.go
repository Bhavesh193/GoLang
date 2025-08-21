package main

import (
	// "errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

func ProcessTruck(truck Truck) error {
	fmt.Printf("Processing truck : %s\n", truck.id)
	
	return nil
	// return errors.New("There is some error in the ProcessTruck")
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
		{id: "Truck-4"},
	}
	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)

		err := ProcessTruck(truck)
		if err != nil {
			log.Printf("Error processing truck: %s", err)
			// log.Fatalf("Error processin truck: %s" , err)
		}
	}
}
