//Khareen Proverbs
package assignments

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck is a collection of cards
type deck []string

// Card is what makes up a deck
type Card struct {
	Suite string
	Value string
}


func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

}


// func GenerateDeck() Deck {
// 	suites := []string{"spades", "hearts", "diamonds", "clubs"}
// 	values := []string{"Ace","two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "Jack", "Queen", "King"}
// 	deck := Deck{}
		
// 	// 'i' and 'j' are replaced with underscores, '_' , because we only need the values not index
// 	for _, suite := range suites {
// 		for _, value := range values {
// 			deck.Cards = append(deck.Cards, Card{Suite: suite, Value: value})
// 		}
// 	}
// 	return deck
// }

func GenerateDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// 'i' and 'j' are replaced with underscores, '_' , to let go know we don't need these variables
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}


// func Shuffle(a []int) {
// 	t := time.Now()
// 	rand.Seed(int64(t.Nanosecond()))
// 	for i := range a {
// 		j := rand.Intn(i + 1)
// 		a[i], a[j] = a[j], a[i]
// 	}
// }

// the second paranthesis are the two types we will be returning
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func BlackJack() {
	cards:= GenerateDeck()
	cards.shuffle()
	fmt.Printf("cards: %v\n", cards)
	fmt.Println(deal(cards,2))
	// fmt.Println(dealCards(shuffledCards, 5))

}
