package main

import (
	"fmt"
	"log"
	"testing"
)

func TestGetMaxProfit(t *testing.T) {

	cases := []struct {
		stockPriceYesterday []int
		expectedProfit      int
	}{
		{[]int{12, 10, 7, 5, 8, 11, 9}, 6},
		{[]int{10, 12, 7, 5, 8, 11, 4}, 6},
		{[]int{58, 95, 10, 72, 22, 47, 43, 43, 38, 9, 97, 5, 41, 19, 85, 10, 66, 22, 76, 10, 71, 20, 55, 82, 55, 50, 54, 77, 96, 78}, 91},
		{[]int{58, 47, 7, 29, 28, 40, 24, 37, 56, 47, 59, 44, 9, 46, 19, 88, 90, 52, 8, 52, 83, 88, 2, 11, 19, 24, 81, 64, 70}, 79},
		{[]int{67, 100, 62, 41, 33, 85, 45, 89, 7, 86, 68, 62, 39, 82, 31, 40, 79, 69, 20, 71, 73, 58, 88, 77, 91, 14, 19, 16, 95, 51, 30}, 88},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc: %v", i), func(t *testing.T) {
			log.Printf("Sample Data Table: %v\n", tc.stockPriceYesterday)
			got := getMaxProfit(tc.stockPriceYesterday)
			if got != tc.expectedProfit {
				t.Errorf("Failed: Expected: %v Got: %v", tc.expectedProfit, got)
			}
		})
	}
}
