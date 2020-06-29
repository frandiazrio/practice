// O(n)


package main
import "fmt"
func ksub(arr []int, k int)int{
	kmap := make([]int, k)
	sum := 0 // current sum
	res := 0 //total results

	kmap[0]=1 // A sum of 0 always divisible

	for i:=0; i< len(arr); i++{
		sum = sum + arr[i]
		sum%= k
		if sum < 0{
			sum += k
		} // check for negative indices

		res += kmap[sum]

		kmap[sum]++

	}

	return res

}


func main(){
	fmt.Println(ksub([]int{5, 10, 11, 9, 5}, 5))
	fmt.Println(ksub([]int{-5}, 5))
	fmt.Println(ksub([]int{7,4,-10}, 5))
	fmt.Println(ksub([]int{4,5,0,-2,-3,1}, 5))

	fmt.Println(-9%5)
}