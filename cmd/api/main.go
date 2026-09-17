package main

import (
	"fmt"
	"time"
)

func getCarInfo(carID string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Car info fetched")
}

func getAvailability(carID string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Availability fetched")
}

func getPricing(carID string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Pricing fetched")
}

func getReviews(carID string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Reviews fetched")
}

func main() {
	start := time.Now()
	carID := "car-123"

	go getCarInfo(carID)
	go getAvailability(carID)
	go getPricing(carID)
	go getReviews(carID)

	fmt.Println("main() reached the end after:", time.Since(start))
	time.Sleep(2 * time.Second) // Wait for goroutines to finish
}