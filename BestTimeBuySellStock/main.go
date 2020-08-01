package main

import "fmt"
import "math"

/**

Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i; and a non-negative integer fee representing a transaction fee.

You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)

Return the maximum profit you can make.



*/

//Return the maximum profit

func BestTime(prices []int, fee int) int {
	if len(prices) == 0 {
		return math.MinInt32
	}
	hold := 0         //keep the money
	buy := -prices[0] // bought first stock

	for i := 1; i < len(prices); i++ {
		hold := Max(hold, prices[i]+hold-fee) // sell our stock
		buy := Max(buy, hold-prices[i])       // buy the stock
	}

	return hold

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("vim-go")
}
