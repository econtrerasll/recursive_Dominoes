package main

import "fmt"

type Player struct {
	Name  string
	Chips []Tiles
}

func (p *Player) getPoints() int {
	var roundScore int
	for _, c := range p.Chips {
		roundScore += c.head + c.tail
	}
	return roundScore
}

func (player *Player) removePlayerChip(chip Tiles) {
	var result []Tiles
	for _, ch := range player.Chips {
		if !(ch == chip) {
			result = append(result, ch)
		}
	}
	player.Chips = result
}

func (p *Player) skip() {
	fmt.Println(p.Name, "skipped his turn!!! Chips left:", p.Chips)
}

func (p *Player) dropTile(move Tiles) {
	var result []Tiles
	for _, ch := range p.Chips {
		if ch != move {
			result = append(result, ch)
		}
	}
	p.Chips = result
	fmt.Println(p.Name, "placed: ", move.toString(), "on the board Chips left on hand: ", p.Chips)
}

func (p *Player) hasDoubleSix() (bool, Tiles) {
	for _, chip := range p.Chips {
		//chip.head == 6 && chip.tail == 6
		dSix := Tiles{6, 6}
		if chip == dSix {
			return true, chip
		}
	}
	return false, Tiles{0, 0}
}

func initPlayers() []Player {
	var players []Player
	var names = []string{"Ramon", "Miguel", "Pedro", "Juan"}
	cards := newChips()
	for i, name := range names {
		p := Player{name, cards[i]}
		players = append(players, p)
		fmt.Println(p.Name, "started this round with the following Chips: ", p.Chips)
	}
	return players
}
