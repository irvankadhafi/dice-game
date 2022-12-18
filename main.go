package main

import (
	"fmt"
	"math/rand"
	"time"
)

// PlayGame adalah fungsi yang memainkan permainan dadu dengan daftar pemain yang diberikan.
// Permainan akan terus berlangsung hingga hanya tersisa 1 pemain yang masih memiliki dadu.
func PlayGame(players []Player) {
	for {
		// Lempar dadu untuk semua pemain
		for i := range players {
			players[i].ThrowDice()
		}

		fmt.Println("===============================")

		// Cetak hasil lemparan dadu untuk setiap pemain
		fmt.Println("Giliran lempar dadu:")
		for i := range players {
			fmt.Printf("Pemain #%d (%d): ", players[i].id, players[i].score)
			for j := range players[i].dice {
				fmt.Printf("%d ", players[i].dice[j])
			}
			fmt.Println()
		}

		// Evaluasi dadu untuk semua pemain
		fmt.Println("Setelah evaluasi:")
		for i := range players {
			players[i].EvaluateDice(players)
			fmt.Printf("Pemain #%d (%d): ", players[i].id, players[i].score)
			if len(players[i].dice) == 0 {
				fmt.Println("_ (Berhenti bermain karena tidak memiliki dadu)")
				players[i].stopped = true
			} else {
				for j := range players[i].dice {
					fmt.Printf("%d ", players[i].dice[j])
				}
				fmt.Println()
			}
		}

		// Periksa apakah hanya tersisa 1 pemain yang masih memiliki dadu
		remainingPlayers := 0
		for i := range players {
			if players[i].HasDice() {
				remainingPlayers++
			}
		}
		if remainingPlayers <= 1 {
			break
		}
	}

	// Cetak pemenangnya
	winner := 0
	highestScore := 0
	for i := range players {
		if players[i].score > highestScore {
			winner = i
			highestScore = players[i].score
		}
	}

	playersWithDice := 0
	for _, p := range players {
		if p.HasDice() {
			playersWithDice = p.id
		}
	}

	// Print the result of the game
	fmt.Println("Game berakhir karena hanya pemain #", players[playersWithDice].id, "yang memiliki dadu.")
	fmt.Println("Game dimenangkan oleh pemain #", players[winner].id, "karena memiliki poin lebih banyak dari pemain lainnya.")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var numPlayers, numDice int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scanln(&numPlayers)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scanln(&numDice)

	// Create the players
	players := make([]Player, numPlayers)
	for i := range players {
		players[i] = NewPlayer(i+1, numDice)
	}

	PlayGame(players)
}
