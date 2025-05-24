package main

import (
	"fmt"
	"io/ioutil"
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

//WriteFile writes data to a file named by filename. Format...
//func WriteFile(filename string, data []byte, perm fs.FileMode) error

func (d deck) savetofile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 6666) // this will make a new plain text file and print the putput
}
