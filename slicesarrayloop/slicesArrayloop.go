package main

import (
	"fmt"
)

func main() {
	cards := []string{"hello , ", newCards()}
	cards = append(cards, "six of spade ")
	for i, value := range cards { //why do we use := again, when we already have declared cards the first time?
		//it is because with every iteration the new value is assigned and we throw away the previous value declaration
		fmt.Println(i, value)
	}
	//fmt.Println(cards)

}
func newCards() string {
	return "Ace of spade ,"
}
