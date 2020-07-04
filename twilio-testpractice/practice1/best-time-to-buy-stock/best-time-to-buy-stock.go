package main

import "fmt"

func bestTimeToBuy(prices []int){
	if len(prices) < 2{
		return 
	}

	buyidx := 0
	sellidx := 1
	buy := prices[buyidx]
	sell := prices[sellidx]


	for i:=1; i<len(prices); i++{
		if prices[i] > sell && buyidx < i{
			sell = prices[i]
			sellidx =i
		}

		if prices[i]<buy && sellidx > buyidx{
			buy = prices[i]
			buyidx = i
		}

		
	}

	if sell - buy <=0 {
		fmt.Println("Do not buy")
	}
	fmt.Printf("Buy at %v at price %v \n", buyidx, buy)
	fmt.Printf("Sell at %v at price %v \n", sellidx, sell)
	fmt.Printf("Profit %v \n", sell-buy)
}


func main(){
	bestTimeToBuy([]int{7,1,5,3,6,4})
}