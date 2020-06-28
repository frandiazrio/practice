package main

import "fmt"
type Queue struct{
	stack []int
	aux_stack [] int
}

func (q *Queue) Init(){
	q.stack = []int{}
	q.aux_stack = []int{}
}

func (q *Queue) Insert(val int){
	
	q.stack = append(q.stack, val)

}

func (q *Queue) Poll()int{
	//
	for len(q.stack) > 1{
		take := q.stack[len(q.stack)-1]
		q.aux_stack = append(q.aux_stack, take)
		q.stack = q.stack[:len(q.stack)-1]
	} 
	res := q.stack[0]
	q.stack = []int{}
	for len(q.aux_stack) >0{
		take := q.aux_stack[len(q.aux_stack)-1]
		q.stack = append(q.stack, take)
		q.aux_stack = q.aux_stack[:len(q.aux_stack)-1]
	}
	

	q.aux_stack = []int{}
	return res
}


func main(){
	var q Queue
	q.Init()
	for i:=0; i< 10; i++{
		q.Insert(i)
	}

	fmt.Println(q)

	q.Poll()
	fmt.Println(q)
}