package main

import "fmt"

type Player struct {
    Name    string
    Cards   []Card
    Score   int
    Stanted bool
}

func NewPlayer(name string) Player {
    var cards []Card
    player := Player{Name: name, Cards: cards, Score: 0, Stanted: false}
    return player
}

func DrawCard(player *Player, deck *[]Card) Card {
    card := (*deck)[0]
    *deck = (*deck)[1:]
    player.Cards = append(player.Cards, card)
    return card
}

func PrintDrawCard(player Player, card Card) {
    fmt.Println(player.Name + "の引いたカードは" + CardSuitName(card) + "の" + CardNumberName(card) + "です。")
}
