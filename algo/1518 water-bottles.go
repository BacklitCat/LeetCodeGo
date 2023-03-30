package algo

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	haveDrink := numBottles                   // 先喝水，剩下的都是空瓶子
	newNumBottles := numBottles / numExchange // 新瓶子
	for newNumBottles >= 1 {                  // 能兑换
		haveDrink += newNumBottles                      // 喝水
		numBottles -= newNumBottles * (numExchange - 1) // 减去旧瓶子，拿来新瓶子
		fmt.Println(numBottles, newNumBottles, haveDrink)
		newNumBottles = numBottles / numExchange
	}
	//不能兑换了
	return haveDrink
}
