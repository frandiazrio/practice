package main

import "time"
import "fmt"

// [["9:00AM" ,"10:00AM"], ["10:00AM","12:00AM"]] --> [["9:00AM", "12:00AM"]]

func main() {
	const layout = "12:30PM"
	res := [][]time.Time{}
	timeSlots := [][]string{{"9:00AM", "10:00AM"}, {"10:00AM", "1:00PM"}}
	for _, slot := range timeSlots {
		beginning, _ := time.Parse(time.Kitchen, slot[0])
		end, _ := time.Parse(time.Kitchen, slot[1])

		if len(res) == 0 {
			res = append(res, []time.Time{beginning, end})
		} else {
			last := res[len(res)-1]

			lastEnd := last[1]

			if lastEnd.Equal(beginning) {
				res[len(res)-1][1] = end
			} else {
				res = append(res, []time.Time{beginning, end})
			}
		}

	}

	fmt.Println(res, len(res))
}
