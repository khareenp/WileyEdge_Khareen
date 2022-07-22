//Khareen Proverbs
package assignments

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// Deck is a collection of cards
type Deck struct {
	Cards []Card
}
// Card is what makes up a deck
type Card struct {
	Suite string
	Value string
}


func (d *Deck) shuffle() (err error){
	var old []Card
	old = d.Cards
	var shuffled []Card
	// This should be a relatively fast shuffle, as we always pick a random number
	// within the remaining cards left to be shuffled.  This has the added bonus
	// of allowing a deck to be a 'shoe' like in a casino where many decks are
	// shuffled together and drawn from.

	// For N times (where N is the total number of cards in the deck)
	for i := len(old); i > 0; i-- {
		// Pick an index within the old cards
		nBig, e := rand.Int(rand.Reader, big.NewInt(int64(i)))
		if e != nil {
			panic(e)
		}
		j := nBig.Int64()
		// Append the randomly picked card to the 'shuffled' pile
		shuffled = append(shuffled, old[j])
		// remove the chosen card from the old pile and collapse
		// (length will be decremented)
		old = remove(old, j)
	}
	d.Cards = shuffled
	return nil

}
// remove removes a card at index i from a slice of Cards and collapses the hole
// (length with be decremented)
func remove(slice []Card, i int64) []Card {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}


func GenerateDeck(d *Deck) error {
	suites := []string{"spades", "hearts", "diamonds", "clubs"}
	values := []string{"Ace","two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "Jack", "Queen", "King"}
	//D := Deck{}
		
	// 'i' and 'j' are replaced with underscores, '_' , because we only need the values not index
	//fill the deck
	for _, suite := range suites {
		for _, value := range values {
			d.Cards= append(d.Cards, Card{Suite: suite, Value: value})
		}
	}
	d.shuffle()
	return nil
}

// the second paranthesis are the two types we will be returning
func deal(d *Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}
func BlackJack() {
	Deck:= Deck{}
	GenerateDeck(&Deck)
	Deck.shuffle()
}
