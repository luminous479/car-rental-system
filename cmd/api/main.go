package main

import (
	
	"fmt"
	
	"time"
)

type CarInfo struct {
	Name string 
}
type Availability struct {
  Available bool

}
type Pricing struct{ PricePerDay float64 }
type Reviews struct{ AverageRating float64 }
type InsuranceQuote struct{ Quote float64 }

func getCarInfo(carID string) CarInfo {
	time.Sleep(1 * time.Second) 
	return CarInfo{Name: "Toyota Corolla"}
}

func getAvailability(carID string) Availability {
	time.Sleep(1 * time.Second)
	return Availability{Available: true}
}

func getPricing(carID string) Pricing {
	time.Sleep(1 * time.Second)
	return Pricing{PricePerDay: 35.0}
}

func getReviews(carID string) Reviews {
	time.Sleep(1 * time.Second)
	return Reviews{AverageRating: 4.5}
}

func getInsuranceQuote(carID string) InsuranceQuote {
	time.Sleep(1 * time.Second)
	return InsuranceQuote{Quote: 15.0}
}


func main() {

	start := time.Now()
	carID := "car-123"
	
	carInfo := getCarInfo(carID)
	availability := getAvailability(carID)
	pricing := getPricing(carID)
	reviews := getReviews(carID)
	insuranceQuote := getInsuranceQuote(carID)

	elapsed := time.Since(start)
	fmt.Printf("Car Info: %+v\n", carInfo)
	fmt.Printf("Availability: %+v\n", availability)
	fmt.Printf("Pricing: %+v\n", pricing)
	fmt.Printf("Reviews: %+v\n", reviews)
	fmt.Printf("Insurance Quote: %+v\n", insuranceQuote)
	fmt.Printf("Time taken: %s\n", elapsed)
}
