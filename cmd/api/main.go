package main

import (
	"context"
	"fmt"
	"time"

)

// type CarInfo struct{ Name string }
// type Availability struct{ Available bool }
// type Pricing struct{ PricePerDay float64 }
// type Reviews struct{ AverageRating float64 }

// func getCarInfo(ctx context.Context, carID string, ch chan<- CarInfo) {
// 	time.Sleep(1 * time.Second)
// 	ch <- CarInfo{Name: "Toyota Corolla"}
// }

// func getAvailability(ctx context.Context, carID string, ch chan<- Availability) {
// 	time.Sleep(1 * time.Second)
// 	ch <- Availability{Available: true}
// }

// func getPricing(ctx context.Context, carID string, ch chan<- Pricing) {
// 	time.Sleep(1 * time.Second)
// 	ch <- Pricing{PricePerDay: 35.0}
// }

// func getReviews(ctx context.Context, carID string, ch chan<- Reviews) {
// 	time.Sleep(1 * time.Second)
// 	ch <- Reviews{AverageRating: 4.5}
// }

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: received cancellation signal, exiting...\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(500 * time.Millisecond) // Simulate work
		}
	}
}

func main() {
	// ctx := context.Background()
	// carID := "car-123"

	// carCh := make(chan CarInfo, 1)
	// availCh := make(chan Availability, 1)
	// priceCh := make(chan Pricing, 1)
	// reviewsCh := make(chan Reviews, 1)

	// go getCarInfo(ctx, carID, carCh)
	// go getAvailability(ctx, carID, availCh)
	// go getPricing(ctx, carID, priceCh)
	// go getReviews(ctx, carID, reviewsCh)

	// info := <-carCh
	// avail := <-availCh
	// price := <-priceCh
	// reviews := <-reviewsCh

	// fmt.Printf("Car: %+v\n", info)
	// fmt.Printf("Availability: %+v\n", avail)
	// fmt.Printf("Pricing: %+v\n", price)
	// fmt.Printf("Reviews: %+v\n", reviews)

	ctx, cancel := context.WithCancel(context.Background())
	id := 1

	go worker(ctx, id)

	time.Sleep(2 * time.Second) // Let the worker run for a while
	fmt.Println("Main: sending cancellation signal to worker...")
	cancel() // Send cancellation signal to the worker

	time.Sleep(1 * time.Second) // Give some time for the worker to exit
	fmt.Println("Main: exiting.")
}