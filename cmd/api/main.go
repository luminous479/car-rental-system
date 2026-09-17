package main

import (
	"context"
	"fmt"
	"time"
)

type Pricing struct {
	PricePerDay float64
}

func getPricing(ctx context.Context, carID string, ch chan<- Pricing) {
	select {
	case <-time.After(3 * time.Second): // simulated slow pricing service
		ch <- Pricing{PricePerDay: 35.0}
	case <-ctx.Done():
		fmt.Println("getPricing: gave up,", ctx.Err())
		// note: nothing gets sent to ch in this branch!
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	priceCh := make(chan Pricing, 1)
	go getPricing(ctx, "car-123", priceCh)

	select {
	case price := <-priceCh:
		fmt.Printf("Got pricing: %+v\n", price)
	case <-ctx.Done():
		fmt.Println("main: giving up on pricing,", ctx.Err())
	}
}
