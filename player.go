package main

import (
	"math/rand"
)

// Player adalah struct yang menyimpan informasi tentang pemain
type Player struct {
	dice    []int // dice merupakan array yang menyimpan nilai-nilai dadu yang dimiliki pemain
	score   int   // score merupakan skor yang dimiliki pemain
	id      int   // id merupakan nomor identitas pemain
	stopped bool  // stopped menandakan apakah pemain sudah berhenti bermain atau belum
}

// NewPlayer adalah fungsi yang membuat objek pemain baru dengan id dan numDice dadu yang diberikan
func NewPlayer(id, numDice int) Player {
	dice := make([]int, numDice)
	return Player{dice: dice, score: 0, id: id, stopped: false}
}

// ThrowDice adalah fungsi yang mengeluarkan nilai dadu untuk setiap dadu yang dimiliki pemain
func (p *Player) ThrowDice() {
	for i := range p.dice {
		p.dice[i] = rand.Intn(6) + 1
	}
}

// EvaluateDice adalah fungsi yang mengevaluasi dadu-dadu yang dimiliki pemain, dengan aturan sebagai berikut:
// - Jika ada dadu yang nilainya 6, maka dadu tersebut dianggap tidak ada dan skor pemain bertambah 1
// - Jika ada dadu yang nilainya 1, maka dadu tersebut diberikan kepada pemain berikutnya, dan tidak ada pengaruh terhadap skor pemain
func (p *Player) EvaluateDice(players []Player) {
	for i := 0; i < len(p.dice); i++ {
		if p.dice[i] == 6 {
			p.dice = append(p.dice[:i], p.dice[i+1:]...)
			p.score++
			i--
		} else if p.dice[i] == 1 {
			nextPlayerIdx := (p.id % len(players)) + 1
			players[nextPlayerIdx-1].dice = append(players[nextPlayerIdx-1].dice, 1)
			p.dice = append(p.dice[:i], p.dice[i+1:]...)
			i--
			//fmt.Printf("\t\tPemain #%d memberikan dadu bernilai 1 kepada pemain #%d\n", p.id, players[nextPlayerIdx-1].id)
		}
	}
}

// HasDice adalah fungsi yang mengembalikan true jika pemain memiliki dadu yang masih tersisa, dan false jika tidak
func (p *Player) HasDice() bool {
	return len(p.dice) > 0
}
