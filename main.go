package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    deck := NewDeck()
    ShuffleDeck(deck)

    player := NewPlayer("あなた")
    dealer := NewPlayer("ディーラー")

    fmt.Println("/* ========== ブラックジャック ========== */")

    PrintDrawCard(player, PlayCard(&player, &deck))
    PrintDrawCard(player, PlayCard(&player, &deck))

    PrintDrawCard(dealer, PlayCard(&dealer, &deck))
    PlayCard(&dealer, &deck)
    fmt.Println("ディーラーは2枚目のカードを引きました。")

    PrintScore(player)

    for {
        if player.Stanted == true && dealer.Stanted == true {
            break
        }

        fmt.Println("/* ------------------------------------ */")

        if player.Stanted == false {
            fmt.Println("カードを引きますか？")
            fmt.Println("引く場合は「Y」を、引かない場合はそれ以外を入力してください。")

            stdin := bufio.NewScanner(os.Stdin)
            stdin.Scan()
            input := stdin.Text()

            if input == "Y" {
                card := PlayCard(&player, &deck)
                PrintDrawCard(player, card)
                if player.Score > 21 {
                    fmt.Println("バーストしました！")
                    break
                }
            } else {
                player.Stanted = true
                fmt.Println("パスしました。")
            }
            PrintScore(player)
        }

        if dealer.Stanted == false && player.Score < 22 {
            if dealer.Score > 18 {
                dealer.Stanted = true
                fmt.Println("ディーラーはパスしました。")
            } else {
                PlayCard(&dealer, &deck)
                fmt.Println("ディーラーはカードを引きました。")
            }
        }
    }
    PrintScore(player)
    PrintScore(dealer)

    switch {
    case player.Score > 21, player.Score == 21 && dealer.Score == 21:
        fmt.Println("あなたの負けです！")
    case dealer.Score > 21, player.Score > dealer.Score:
        fmt.Println("あなたの勝ちです！")
    case player.Score == dealer.Score:
        fmt.Println("引き分けです！")
    default:
        fmt.Println("あなたの負けです！")
    }

    fmt.Println("/* ==================================== */")
}

func PlayCard(player *Player, deck *[]Card) Card {
    card := DrawCard(player, deck)
    CalcScore(player)
    return card
}

// A(1) = 1, J(11)〜K(13) = 10, それ以外は数字どおり
func CardScore(card Card) int {
    score := 0
    switch card.Number {
    case 11, 12, 13:
        score = 10
    default:
        score = card.Number
    }
    return score
}

func CalcScore(player *Player) {
    score := 0
    n := len(player.Cards)
    for i := 0; i < n; i++ {
        score += CardScore(player.Cards[i])
    }
    player.Score = score
}

func PrintScore(player Player) {
    fmt.Println(player.Name + "の得点は" + fmt.Sprintf("%v", player.Score) + "点です。")
}
