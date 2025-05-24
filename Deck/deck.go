package main

import (
	"fmt"
	"strings"
)

// NOTE....
// returning sonething and Printing something is different ...
// returning means tht particular funtion is returning a particular value...
// but if the value isnt Printed with using Printf statements it will not be shown in the output
type deck []string //deck is string typed

func newDeck() deck {
	card := deck{}
	fmt.Println("deck at the start is %+v", deck{})
	Cardshelf := []string{"Clubs", "Spades", "heart", "diamond"}
	CardValue := []string{"Ace", "Two", "three", "four"}
	for _, shelf := range Cardshelf { //1- Clubs
		for _, value := range CardValue { //1- Ace
			card = append(card, shelf+"of"+value) //1 -  Clubs of Ace
		}
	}
	fmt.Println("cards at the end is %+v", card)
	return card //new deck returns the cards

}
func (d deck) print() { //deck is already declatred in the main.go
	for i, val := range d {
		fmt.Println(i, val)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	//return ([]string(d)) //ClubsofAce ClubsofTwo Clubsofthree Clubsoffour Spadeso "if the return type is []string"
	//return cards
}

//output....
// 0 ClubsofAce
// 1 ClubsofTwo
// 2 Clubsofthree
// 3 Clubsoffour
// 4 SpadesofAce
// 5 SpadesofTwo
