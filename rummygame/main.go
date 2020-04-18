package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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

func RemoveCard(cards []string, card string) []string {
	// Removes the card which the user wants to give away
	var index int
	for i, v := range cards {
		if v == card {
			index = i
		}
	}
	cards = append(cards[:index], cards[index+1:]...)
	return cards
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func sortCards(cards []string) []string {
	// SOrt the cards according to the color
	sorted_h := []string{}
	sorted_d := []string{}
	sorted_c := []string{}
	sorted_s := []string{}
	sorted_j := []string{}
	sorted_deck := make([]string, 13)
	// fmt.Println(cards)
	for _, card := range cards {
		// Break the cards in 4 different arrays according to the color
		if string(card[0]) == "h" {
			sorted_h = append(sorted_h, card)
		} else if string(card[0]) == "c" {
			sorted_c = append(sorted_c, card)
		} else if string(card[0]) == "d" {
			sorted_d = append(sorted_d, card)
		} else if string(card[0]) == "s" {
			sorted_s = append(sorted_s, card)
		} else if string(card[0]) == "j" {
			sorted_j = append(sorted_j, card)
		}
	}
	// fmt.Println(reflect.TypeOf(sorted_h))
	// fmt.Println(sorted_d, sorted_s, sorted_j, sorted_d, sorted_deck)
	// // Reconstruct the cards
	sorted_deck = append(sorted_h)
	sorted_deck = append(sorted_deck, sorted_c...)
	sorted_deck = append(sorted_deck, sorted_d...)
	sorted_deck = append(sorted_deck, sorted_s...)
	sorted_deck = append(sorted_deck, sorted_j...)
	return sorted_deck
}

func main() {
	// Generate Cards

	var rummyDeclared bool = false
	cards := [54]string{
		"cA", "c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "c10", "cG", "cQ", "cK",
		"dA", "d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "d10", "dG", "dQ", "dK",
		"hA", "h2", "h3", "h4", "h5", "h6", "h7", "h8", "h9", "h10", "hG", "hQ", "hK",
		"sA", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10", "sG", "sQ", "sK",
		"j1", "j2",
	}
	// shuffle cards
	shuffledCards := shuffle(cards)
	// DIstribute the cards to two users
	user1_cards := []string{}
	user2_cards := []string{}
	remaining_deck := []string{}
	usedCardsDeck := []string{}
	var RummyClaimeduser int
	for i, card := range shuffledCards {
		if i <= 25 {
			if i%2 == 0 {
				user1_cards = append(user1_cards, card)
				// sort the cards color wise just like any human does
			} else {
				user2_cards = append(user2_cards, card)
				// sort the cards color wise just like any human does
			}
		} else {
			remaining_deck = append(remaining_deck, card)
		}
	}
	// Make a Joker from the cat
	joker := makeJoker(shuffledCards)

	// Game setup is now completed... Players can start playing
	user1_cards = sortCards(user1_cards)
	user2_cards = sortCards(user2_cards)
	// Which PLayer will play next....will change to 2 after player 1 plays and viz
	// find_type(user1_cards)
	var playerturn int = 1
	var upperCard string = ""
	var cardToDrop string = ""
	// GameLoop
	for rummyDeclared != true {
		fmt.Println("Joker Card :-", joker)
		upperCard = remaining_deck[0]

		var userInput string = ""
		switch playerturn {
		case 1:
			fmt.Println("Upper Card :-", upperCard)
			// player one logic
			fmt.Println("Player 1 turn ...")
			fmt.Println("[PLAYER 1] ->", user1_cards)
			fmt.Print("[!] Press u if you want to take the upper card press d if you want to take from deck  \n CMD >")
			fmt.Scanln(&userInput)
			switch userInput {
			case "u":
				user1_cards = append(user1_cards, upperCard)
				fmt.Println("[PLAYER 1] ", user1_cards)
				fmt.Println("[!] You have to Drop one card, enter the card name to drop")
				fmt.Scanln(&cardToDrop)
				user1_cards = RemoveCard(user1_cards, cardToDrop)
				// Again sort the cards,
				user1_cards = sortCards(user1_cards)
				upperCard = cardToDrop

				// Since the user picked up the upper card, remove the upper card from
				go RemoveCard(remaining_deck, upperCard)

				// Add that card to usedcards deck
				usedCardsDeck = append(usedCardsDeck, upperCard)
				fmt.Println("[PLAYER 1]", user1_cards)
				fmt.Println("[+] Press p to pass to next player")
				var pass string
				fmt.Scanln(&pass)
				_ = pass
			case "d":
				// user didnt pick up the upper card, so add it to the used card deck.
				usedCardsDeck = append(usedCardsDeck, upperCard)
				// then remove the card from the main deck
				remaining_deck = RemoveCard(remaining_deck, upperCard)
				// make the next card as upper card
				upperCard = remaining_deck[0]
				fmt.Println("card_picked from dec", upperCard)
				user1_cards = append(user1_cards, upperCard)
				fmt.Println("[PLAYER 1] ", user1_cards)
				fmt.Println("[!] You have to Drop one card, enter the card name to drop")
				fmt.Scanln(&cardToDrop)
				user1_cards = RemoveCard(user1_cards, cardToDrop)
				// Again sort the cards,
				user1_cards = sortCards(user1_cards)
				upperCard = cardToDrop
				fmt.Println("[PLAYER 1]", user1_cards)
				fmt.Println("[+] Press p to pass to next player")
				var pass string
				fmt.Scanln(&pass)
				_ = pass
			case "rummy":
				rummyDeclared = true
				RummyClaimeduser = 1

			}

			// CLear the screen for next player
			clear()
			playerturn = 2
		case 2:
			// player two logic
			fmt.Println("Upper Card :-", upperCard)
			fmt.Println("Player 2 turn ...")
			fmt.Println("[PLAYER 2] ->", user2_cards)
			fmt.Print("[!] Press u if you want to take the upper card/ press d if you want to take from deck \n CMD >")
			fmt.Scanln(&userInput)
			switch userInput {
			case "u":
				user2_cards = append(user2_cards, upperCard)
				fmt.Println("[PLAYER 2] ", user2_cards)
				fmt.Println("[!] You have to Drop one card, enter the card name to drop")

				// WORK FROM HERE
				fmt.Scanln(&cardToDrop)
				user2_cards = RemoveCard(user2_cards, cardToDrop)
				// Again sort the cards,
				user2_cards = sortCards(user2_cards)
				upperCard = cardToDrop

				// Since the user picked up the upper card, remove the upper card from
				RemoveCard(remaining_deck, upperCard)

				// Add that card to usedcards deck
				usedCardsDeck = append(usedCardsDeck, upperCard)
				fmt.Println("[PLAYER 2]", user2_cards)
				fmt.Println("[+] Press p to pass to next player")
				var pass string
				fmt.Scanln(&pass)
				_ = pass
			case "d":
				// user didnt pick up the upper card, so add it to the used card deck.
				usedCardsDeck = append(usedCardsDeck, upperCard)
				// then remove the card from the main deck
				remaining_deck = RemoveCard(remaining_deck, upperCard)
				// make the next card as upper card
				upperCard = remaining_deck[0]
				fmt.Println("card_picked from dec", upperCard)
				user2_cards = append(user2_cards, upperCard)
				fmt.Println("[PLAYER 2] ", user2_cards)
				fmt.Println("[!] You have to Drop one card, enter the card name to drop")
				fmt.Scanln(&cardToDrop)
				user2_cards = RemoveCard(user2_cards, cardToDrop)
				// Again sort the cards,
				user2_cards = sortCards(user2_cards)
				upperCard = cardToDrop
				fmt.Println("[PLAYER 2]", user2_cards)
				fmt.Println("[+] Press p to pass to next player")
				var pass string
				fmt.Scanln(&pass)
				_ = pass
			case "rummy":
				rummyDeclared = true
				RummyClaimeduser = 2

			}
			clear()
			playerturn = 1
		}
	}
	// RUmmy has been claimed, display to both
	clear()
	var userAccept string
	switch RummyClaimeduser {
	case 1:
		fmt.Println("Rummy Claimed by player 1")
		fmt.Println("Player 1 CARDS :-")
		fmt.Println("[PLAYER 1] - ", user1_cards)
		fmt.Println("PLAYER 2, Do you accept player 1 RUMMY ? [y/n]")
		fmt.Scanln(&userAccept)
		if userAccept == "y" {
			fmt.Println("User 1 won")
		} else {
			fmt.Println("oops...")
		}
	case 2:
		fmt.Println("Rummy Claimed by player 2")
		fmt.Println("Player 2 CARDS :-")
		fmt.Println("[PLAYER 2] - ", user1_cards)
		fmt.Println("PLAYER 1, Do you accept player 2 RUMMY ? [y/n]")
		fmt.Scanln(&userAccept)
		if userAccept == "y" {
			fmt.Println("User 2 won")
		} else {
			fmt.Println("oops...")
		}
	}

}
