package main

import (
	"log"
)

func getMaxProfit(stockPrices []int) int {
	var lowestPrice int
	var highestPrice int
	var lowestPriceIndex int
	var highestPriceIndex int

	for i, price := range stockPrices {
		if (price < lowestPrice || lowestPrice == 0) && i != len(stockPrices)-1 {
			lowestPrice = price
			lowestPriceIndex = i
			highestPrice = 0
		} else if (price > highestPrice) && (i > lowestPriceIndex) {
			highestPrice = price
			highestPriceIndex = i
		}
	}

	log.Printf("Lowest[%v]: %v, Highest[%v]: %v\n", lowestPriceIndex, lowestPrice, highestPriceIndex, highestPrice)
	return highestPrice - lowestPrice
}
