package main
import "fmt"
func ksub(arr []int, k int)[][]int{
	//generate all sub array
	res := [][]int{}
	for i := 0; i<len(arr)-1;i++{
		for j:=i+1; j<len(arr); j++{
			partial := arr[i:j+1]
			if SumDivisible(partial, k){
				res = append(res, partial)
			}
		}
	}

	return res
}

func SumDivisible(input []int, k int)bool{
	sum := 0
	for _,v := range input{
		sum +=v
	}

	return sum % k == 0
}


func main(){
	fmt.Println(ksub([]int{5,10,11,9}, 5))
}