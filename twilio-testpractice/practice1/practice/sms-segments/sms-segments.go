package main

import (
	"fmt"
)

func sms_maker(sms string) []string {
	res := []string{}

	if len(sms) <= 160 {
		res = append(res, sms)
		return res
	}

	// 160 - 5 "(%d/%d)" -1 for space
	for len(sms) > 155 {
		segment := sms[:155]
		res = append(res, segment)
		sms = sms[155:]
	}

	res = append(res, sms)
	// for every segment add segment suffix to string
	for i := 0; i < len(res); i++ {
		seg := fmt.Sprintf("(%d/%d)", i+1, len(res))
		res[i] = fmt.Sprintf("%s%s", res[i], seg)
	}

	return res
}

func main() {

	// testing
	message1 := "Hello, this is a test message."
	message2 := "On the far-away island of Sala-ma-Sond, Yertle the Turtle was king of the pond. A nice little pond. It was clean. It was neat. The water was warm. There was plenty to eat. The turtles had everything turtles might need. And they were all happy. Quite happy indeed."
	expected1 := []string{"Hello, this is a test message."}

	fmt.Println(sms_maker(message1))
	fmt.Println(expected1)
	res := sms_maker(message2)
	for _, v := range res {
		fmt.Println(v, len(v))
	}

}
