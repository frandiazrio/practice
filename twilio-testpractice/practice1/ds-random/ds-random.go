package main
import "time"
import "math/rand"
type RandomSet struct {
	mp   map[interface{}]int
	list []interface{}
}

func NewRandomSet() RandomSet {
	return RandomSet{}
}

func (r *RandomSet) Insert(val interface{}) {
	index := len(r.mp)
	mp[val] = index
	r.list = append(r.list, val)
}

func (r *RandomSet) Remove(val interface{}) {
	index, ok := r.mp[val]
	if ok {
		delete(r.mp, val)
		r.list = append(r.list[:index], r.list[index+1:]...)
	}
}

func (r *RandomSet) GetRandom() interface{}{
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(len(r.list))
	return list[index]

}


func main(){

}