package main

type Entry struct {
	start, end int
}

func NumberOfMeetingRooms(meetings [][]int) int {

	rooms := []Entry{}
	for _, m := range meetings {
		if len(rooms) == 0 {
			rooms = append(rooms, Entry{m[0], m[1]})
		} else {
			index, can := canSwap(rooms, m)
			if can {
				rooms[index] = Entry{m[0], m[1]}
				continue
			} else {
				rooms = append(rooms, Entry{m[0], m[1]})
			}

		}
	}

	return len(rooms)

}

func canSwap(rooms []Entry, m []int) (int, bool) {
	for i, r := range rooms {
		if r.end < m[0] {
			return i, true
		}
	}

	return -1, false
}

func main() {

}
