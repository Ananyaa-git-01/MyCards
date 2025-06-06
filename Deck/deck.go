package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// NOTE....
// returning sonething and Printing something is different ...
// returning means tht particular funtion is returning a particular value...
// but if the value isnt Printed with using Printf statements it will not be shown in the output
type deck []string //deck is string typed

func newDeck() deck {
	card := deck{}
	fmt.Printf("deck at the start is %+v", deck{})
	Cardshelf := []string{"Clubs", "Spades", "heart", "diamond"}
	CardValue := []string{"Ace", "Two", "three", "four"}
	for _, shelf := range Cardshelf { //1- Clubs
		for _, value := range CardValue { //1- Ace
			card = append(card, shelf+"of"+value) //1 -  Clubs of Ace
		}
	}
	fmt.Printf("cards at the end is %+v", card)
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

//to read the file .....
// func ReadFile(filename string) ([]byte, error) {
// 	return os.ReadFile(filename)
// }

func ReadtheFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//option 1 - we can print the error
		//option 2-  we can print the error and quit the program entirely
		fmt.Printf("error in reading the file %+v", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), " ")
	return deck(s)

}

func (d deck) shuffle() {
	for i := range d {

		rand.Seed(time.Now().UnixNano()) // Ensures different shuffle each run

		num := rand.Intn(len(d)) //this will ensure all the numbers are considered for the randomness
		fmt.Printf("the random number generated %+v", num)
		d[i], d[num] = d[num], d[i]

	}
}
