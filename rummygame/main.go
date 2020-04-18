package main

import (
	"math/rand"
	"time"
)

// Implementation of rummy game in GOlang

// club - kilwar- 13 - c1","c2","c3","c4","c5...
// diam - chaukat - 13 - d1","d2","d3","d4...
// Hearts - badam - 13 - h1","h2","h3","h4...
// Spades - ispik - 13 - s1","s2","s3","s4...
// J - joker - 2
// G,Q,K - guddu,queen,king

func shuffle(cards [54]string) [54]string {
	// FIrst we implement Fisher-yates algorithm, to get maximum randomness
	rand.Seed(time.Now().UnixNano())
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

func makeJoker(cards [54]string) string {
	rand.Seed(time.Now().UnixNano())
	joker := cards[rand.Intn(54)]
	if joker != "j" {
		return joker
	} else {
		return "cA"
	}
}

func main() {
	// Generate Cards
	cards := [54]string{
		"cA", "c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "c10", "cG", "cQ", "cK",
		"dA", "d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "d10", "dG", "dQ", "dK",
		"hA", "h2", "h3", "h4", "h5", "h6", "h7", "h8", "h9", "h10", "hG", "hQ", "hK",
		"sA", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10", "sG", "sQ", "sK",
		"j", "j",
	}
	// shuffle cards
	shuffledCards := shuffle(cards)
	// DIstribute the cards to two users
	user1_cards := []string{}
	user2_cards := []string{}
	remaining_deck := []string{}

	for i, card := range shuffledCards {
		if i <= 25 {
			if i%2 == 0 {
				user1_cards = append(user1_cards, card)
			} else {
				user2_cards = append(user2_cards, card)
			}
		} else {
			remaining_deck = append(remaining_deck, card)
		}
	}
	// Make a Joker from the cat
	joker := makeJoker(shuffledCards)

	// Game setup is now completed... Players can start playing

}
