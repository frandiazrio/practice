package main

import "fmt"

func fib(n int)[]int{
	if n < 0{
		return nil
	}else if n == 0{
		return []int{0}
	}else if n == 1{
		return []int{0,1}
	}else{
		f := make([]int, n)
		f[0]=0
		f[1]=1
		for i := 2; i<n; i++{
			f[i] = f[i-1] + f[i-2]
		}
		fmt.Println(f)
		return f
	}

	
}



func main(){
	fib(20)
}
