package main

import "fmt"

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

//output....
// 0 ClubsofAce
// 1 ClubsofTwo
// 2 Clubsofthree
// 3 Clubsoffour
// 4 SpadesofAce
// 5 SpadesofTwo
