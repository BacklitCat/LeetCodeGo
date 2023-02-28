package algo

import (
	"log"
	"math"
)

func test121() {
	log.Println(maxProfit([]int{1}))                   //0
	log.Println(maxProfit([]int{7, 6, 4, 3, 1}))       //0
	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))    //5
	log.Println(maxProfit([]int{3, 2, 10, 12}))        //10
	log.Println(maxProfit([]int{3, 14, 3, 2, 10, 12})) //11
	log.Println(maxProfit([]int{3, 14, 2, 2, 10, 12})) //11

}

func maxProfit(prices []int) int {
	minPrice, maxPrice := math.MaxInt, math.MinInt
	profit, res := 0, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			profit, maxPrice = 0, math.MinInt
		}
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
		if maxPrice-minPrice > profit {
			profit = maxPrice - minPrice
			if profit > res {
				res = profit
			}
		}
	}
	return res
}
