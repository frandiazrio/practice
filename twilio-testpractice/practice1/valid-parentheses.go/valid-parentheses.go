package main
import "fmt"
func IsValid(p string)bool{

	opening:= map[string]struct{} {
		"(": struct{}{},
		"{": struct{}{},
		"[": struct{}{},
	}
	


	closingToOpening := map[string]string {
		")": "(",
		"}": "{",
		"]": "[",
	}

	stack := []string{}
	for _,c := range p{
		
		fmt.Println(stack)
		_, ok := opening[string(c)]
		if ok{
			stack = append(stack, string(c))
			continue
		}

		
		openingBracket := closingToOpening[string(c)]

		if len(stack) != 0{
			if openingBracket != stack[len(stack)-1]{
				return false
			}
			stack = stack[:len(stack)-1]
		}else{
			return false
		}
	}

	if len(stack) != 0{
		return false
	}
	return true

}

func main(){
	fmt.Println(IsValid("((({})))"))
}