package main

import (
	"context"
	"fmt"
	"time"
)

type CarInfo struct{ Name string }
type Availability struct{ Available bool }
type Pricing struct{ PricePerDay float64 }
type Reviews struct{ AverageRating float64 }

func getCarInfo(ctx context.Context, carID string, ch chan<- CarInfo) {
	time.Sleep(1 * time.Second)
	ch <- CarInfo{Name: "Toyota Corolla"}
}

func getAvailability(ctx context.Context, carID string, ch chan<- Availability) {
	time.Sleep(1 * time.Second)
	ch <- Availability{Available: true}
}

func getPricing(ctx context.Context, carID string, ch chan<- Pricing) {
	time.Sleep(1 * time.Second)
	ch <- Pricing{PricePerDay: 35.0}
}

func getReviews(ctx context.Context, carID string, ch chan<- Reviews) {
	time.Sleep(1 * time.Second)
	ch <- Reviews{AverageRating: 4.5}
}

func main() {
	ctx := context.Background()
	carID := "car-123"

	carCh := make(chan CarInfo, 1)
	availCh := make(chan Availability, 1)
	priceCh := make(chan Pricing, 1)
	reviewsCh := make(chan Reviews, 1)

	go getCarInfo(ctx, carID, carCh)
	go getAvailability(ctx, carID, availCh)
	go getPricing(ctx, carID, priceCh)
	go getReviews(ctx, carID, reviewsCh)

	info := <-carCh
	avail := <-availCh
	price := <-priceCh
	reviews := <-reviewsCh

	fmt.Printf("Car: %+v\n", info)
	fmt.Printf("Availability: %+v\n", avail)
	fmt.Printf("Pricing: %+v\n", price)
	fmt.Printf("Reviews: %+v\n", reviews)
}