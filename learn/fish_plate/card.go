package fush_plate

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numOfDecks        = 2                                      // 使用两副牌
	numOfCards        = 52                                     // 每副牌有52张牌
	numOfPlayers      = 4                                      // 发牌给4个玩家
	numOfCardsPerHand = numOfCards * numOfDecks / numOfPlayers // 每个玩家应该得到的牌数
)

type Card struct {
	Suit  string
	Value string
}

type Deck struct {
	Cards []Card
}

func (d *Deck) Initialize() {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for i := 0; i < numOfDecks; i++ {
		for _, suit := range suits {
			for _, value := range values {
				card := Card{Suit: suit, Value: value}
				d.Cards = append(d.Cards, card)
			}
		}
	}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) Deal(numOfCards int) []Card {
	hand := make([]Card, numOfCards)
	copy(hand, d.Cards[:numOfCards])
	d.Cards = d.Cards[numOfCards:]
	return hand
}

func test() {
	deck := Deck{}
	deck.Initialize()

	deck.Shuffle()
	fmt.Println(deck)
	return
	for i := 0; i < numOfPlayers; i++ {
		hand := deck.Deal(numOfCardsPerHand)
		fmt.Printf("Player %d's hand: %v\n", i+1, hand)
	}
}
