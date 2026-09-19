package main

import (
	"context"
	"net/http"

	"fmt"
	"math/rand"
	"time"
)

// type CarInfo struct{ Name string }
// type Availability struct{ Available bool }
// type Pricing struct{ PricePerDay float64 }
// type Reviews struct{ AverageRating float64 }

// // Generic-ish result wrapper: carries either a value or an error, never both.
// type carInfoResult struct {
// 	value CarInfo
// 	err   error
// }
// type availabilityResult struct {
// 	value Availability
// 	err   error
// }
// type pricingResult struct {
// 	value Pricing
// 	err   error
// }
// type reviewsResult struct {
// 	value Reviews
// 	err   error
// }

// func getCarInfo(ctx context.Context, carID string, ch chan<- carInfoResult) {
// 	select {
// 	case <-time.After(500 * time.Millisecond):
// 		ch <- carInfoResult{value: CarInfo{Name: "Toyota Corolla"}}
// 	case <-ctx.Done():
// 		ch <- carInfoResult{err: ctx.Err()}
// 	}
// }

// func getAvailability(ctx context.Context, carID string, ch chan<- availabilityResult) {
// 	select {
// 	case <-time.After(500 * time.Millisecond):
// 		// simulate an occasional hard failure from the service itself
// 		if rand.Intn(4) == 0 {
// 			ch <- availabilityResult{err: errors.New("fleet service unreachable")}
// 			return
// 		}
// 		ch <- availabilityResult{value: Availability{Available: true}}
// 	case <-ctx.Done():
// 		ch <- availabilityResult{err: ctx.Err()}
// 	}
// }

// func getPricing(ctx context.Context, carID string, ch chan<- pricingResult) {
// 	select {
// 	case <-time.After(700 * time.Millisecond):
// 		ch <- pricingResult{value: Pricing{PricePerDay: 35.0}}
// 	case <-ctx.Done():
// 		ch <- pricingResult{err: ctx.Err()}
// 	}
// }

// func getReviews(ctx context.Context, carID string, ch chan<- reviewsResult) {
// 	select {
// 	case <-time.After(600 * time.Millisecond):
// 		// simulate reviews being flaky — but this is an OPTIONAL dependency
// 		if rand.Intn(3) == 0 {
// 			ch <- reviewsResult{err: errors.New("reviews service timed out")}
// 			return
// 		}
// 		ch <- reviewsResult{value: Reviews{AverageRating: 4.5}}
// 	case <-ctx.Done():
// 		ch <- reviewsResult{err: ctx.Err()}
// 	}
// }

func main() {
	// rand.Seed(time.Now().UnixNano())

	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()

	// carID := "car-123"

	// carCh := make(chan carInfoResult, 1)
	// availCh := make(chan availabilityResult, 1)
	// priceCh := make(chan pricingResult, 1)
	// reviewsCh := make(chan reviewsResult, 1)

	// go getCarInfo(ctx, carID, carCh)
	// go getAvailability(ctx, carID, availCh)
	// go getPricing(ctx, carID, priceCh)
	// go getReviews(ctx, carID, reviewsCh)

	// // --- Step A: car info is critical — no car, no response at all ---
	// carRes := <-carCh
	// if carRes.err != nil {
	// 	cancel()
	// 	fmt.Println("FATAL: could not get car info:", carRes.err)
	// 	return
	// }

	// // --- Step B: availability is critical — cancel everything else if it fails ---
	// availRes := <-availCh
	// if availRes.err != nil || !availRes.value.Available {
	// 	cancel() // tell pricing/reviews to give up if still running
	// 	if availRes.err != nil {
	// 		fmt.Println("FATAL: could not check availability:", availRes.err)
	// 	} else {
	// 		fmt.Println("Car is not available — no need to check pricing or reviews")
	// 	}
	// 	return
	// }

	// // --- Step C: pricing is also critical (can't rent without a price) ---
	// priceRes := <-priceCh
	// if priceRes.err != nil {
	// 	cancel()
	// 	fmt.Println("FATAL: could not get pricing:", priceRes.err)
	// 	return
	// }

	// // --- Step D: reviews is OPTIONAL — degrade gracefully on failure ---
	// reviewsRes := <-reviewsCh
	// var reviews *Reviews
	// if reviewsRes.err != nil {
	// 	fmt.Println("WARNING: reviews unavailable:", reviewsRes.err)
	// 	reviews = nil // response will just omit reviews
	// } else {
	// 	reviews = &reviewsRes.value
	// }

	// // --- Final combined response ---
	// fmt.Printf("\nCar: %+v\n", carRes.value)
	// fmt.Printf("Availability: %+v\n", availRes.value)
	// fmt.Printf("Pricing: %+v\n", priceRes.value)
	// if reviews != nil {
	// 	fmt.Printf("Reviews: %+v\n", *reviews)
	// } else {
	// 	fmt.Println("Reviews: unavailable")
	// }


	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		
		fmt.Fprintln(w, "car rental API")
	})

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}