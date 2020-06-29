package main

func main() {

	type Item struct {
		value    string
		priority int
		index    int
	}

	type PriorityQueue []*Item

	// Len()
	// Less(i,j int)
	// Swap(i,j int)
	// Push(x interface{})
	// Pop() interface{}



	func (pq PriorityQueue) Len() int { return len(pq) }

	func (pq PriorityQueue) Less(i, j int)bool{
		return pq[i].priority > pq[j].priority
	}


	func (pq PriorityQueue) Swap(i,j int){
		p[i], p[j] = p[j], p[i]
		p[i].index = i
		p[j].index = j
	}




	func (pq *PriorityQueue) Push(x interface{}){
		n := len(*pq)
		item := x.()
	}








}
