package main

import "log"

func getMaxProfit(stockPrices []int) int {
	var lowestPrice int
	var highestPrice int

	for _, price := range stockPrices {
		if price < lowestPrice || lowestPrice == 0 {
			lowestPrice = price
		}
		if price > highestPrice {
			highestPrice = price
		}
	}

	log.Printf("Lowest: %v, Highest: %v\n", lowestPrice, highestPrice)
	return highestPrice - lowestPrice
}
