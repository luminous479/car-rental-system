package main

import (
	"fmt"
	"sync"
	"time"
)

func getCarInfo(carID string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Car info fetched")
}

func getAvailability(carID string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Availability fetched")
}

func getPricing(carID string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println("Pricing fetched")
}

func getReviews(carID string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Reviews fetched")
}

func main() {
	start := time.Now()
	carID := "car-123"

	var wg sync.WaitGroup
	wg.Add(4)

	go getCarInfo(carID, &wg)
	go getAvailability(carID, &wg)
	go getPricing(carID, &wg)
	go getReviews(carID, &wg)

	wg.Wait()

	fmt.Println("All done after:", time.Since(start))
}