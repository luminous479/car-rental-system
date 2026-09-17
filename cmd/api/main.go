package main

import (
	"fmt"
	"time"
)

type CarInfo struct{ Name string }
type Availability struct{ Available bool }
type Pricing struct{ PricePerDay float64 }
type Reviews struct{ AverageRating float64 }

func getCarInfo(carID string, ch chan<- CarInfo) {
	time.Sleep(1 * time.Second)
	ch <- CarInfo{Name: "Toyota Corolla"}
}

func getAvailability(carID string, ch chan<- Availability) {
	time.Sleep(1 * time.Second)
	ch <- Availability{Available: true}
}

func getPricing(carID string, ch chan<- Pricing) {
	time.Sleep(1 * time.Second)
	ch <- Pricing{PricePerDay: 35.0}
}

func getReviews(carID string, ch chan<- Reviews) {
	time.Sleep(1 * time.Second)
	ch <- Reviews{AverageRating: 4.5}
}

func main() {
	start := time.Now()
	carID := "car-123"

	carCh := make(chan CarInfo, 1)
	availCh := make(chan Availability, 1)
	priceCh := make(chan Pricing, 1)
	reviewsCh := make(chan Reviews, 1)

	go getCarInfo(carID, carCh)
	go getAvailability(carID, availCh)
	go getPricing(carID, priceCh)
	go getReviews(carID, reviewsCh)

	// Receive from each channel — order here is just the order we choose
	// to *read* results, not the order they *finish* in.
	info := <-carCh
	avail := <-availCh
	price := <-priceCh
	reviews := <-reviewsCh

	fmt.Printf("Car: %+v\n", info)
	fmt.Printf("Availability: %+v\n", avail)
	fmt.Printf("Pricing: %+v\n", price)
	fmt.Printf("Reviews: %+v\n", reviews)
	fmt.Println("Total time:", time.Since(start))
}