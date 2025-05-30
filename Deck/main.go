package main

func main() {
	//cards := []string{"hello , ", newCards()}
	//cards := deck{"hello , ", newCards()} //this is the same as the line above because in the same package we did type deck []string
	cards := newDeck() //clubs of ace
	//fmt.Printf("before shuffle %+v \n", cards)
	cards.shuffle()
	//fmt.Printf("after shuffle %+v \n", cards)
	cards.print()
	//cards := ReadtheFile("My_cards") // to read the file
	//cards = append(cards, "six of spade ")

	//fmt.Println("Cards in the main func %+v", cards)
	//fmt.Println(cards.toString()) //cards ko one by one print krega
	//dot means ... cards ko "print" kro
	// cards.savetofile("My_cards")
	//print is a function in the deck.go which does the same as a for loop
	// for i, value := range cards {
	// 	fmt.Println(i, value)
	// }
	// hands, remaining := deal(cards, 5) //deck( "kya print krna hai", "kitna print krna hai" )
	// hands.print()
	// remaining.print()
	// gift := "Hello"
	// fmt.Println(([]byte(gift))) // this will print the type converted values.
}

// func newCards() string {
// 	//return "Ace of spade ,"
// }
