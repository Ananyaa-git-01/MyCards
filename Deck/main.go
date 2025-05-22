package main

import "fmt"

func main() {
	//cards := []string{"hello , ", newCards()}
	//cards := deck{"hello , ", newCards()} //this is the same as the line above because in the same package we did type deck []string
	cards := newDeck() //clubs of ace
	//cards = append(cards, "six of spade ")
	fmt.Println("Cards in the main func %+v", cards)
	cards.print() //cards ko one by one print krega
	//dot means ... cards ko "print" kro

	//print is a function in the deck.go which does the same as a for loop
	// for i, value := range cards {
	// 	fmt.Println(i, value)
	// }
	hands, remaining := deal(cards, 5) //deck( "kya print krna hai", "kitna print krna hai" )
	hands.print()
	remaining.print()
}

// func newCards() string {
// 	//return "Ace of spade ,"
// }
