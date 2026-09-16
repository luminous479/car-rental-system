package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CarInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	client := &http.Client{}

	fmt.Println("Before request:", time.Now())

	res, err := client.Get("https://jsonplaceholder.typicode.com/users/2")

	fmt.Println("After request:", time.Now())

	if err != nil {
		fmt.Println("request failed", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("request failed with status", res.Status)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("failed to read response body", err)
		return
	}

	var carInfo CarInfo
	err = json.Unmarshal(body, &carInfo)
	if err != nil {
		fmt.Println("failed to unmarshal response body", err)
		return
	}

	fmt.Printf("Car Info: %+v\n", carInfo)
}
