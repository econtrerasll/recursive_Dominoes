package main

import "fmt"

type Game struct {
	LeftSide, RightSide int
	currentPlayerIdx    int
	skippedPlayers      int
	ValidMove           bool
	isNewRound          bool
	players             []Player
}

func initGame() *Game {
	players := initPlayers()
	return &Game{0, 0, 0, 0, false, true, players}
}

func (g *Game) currentPlayerPlayOrSkip() {
	leftSide := g.LeftSide
	rightSide := g.RightSide
	var validMoves []Tiles
	var move Tiles
	newGame := g.isGameNew()

	if !newGame {
		//If game is not new we are in a new round check for valid moves or skipPlayerTurn turn
		for _, chip := range g.currentPlayer().Chips {
			if chip.tail == leftSide || chip.tail == rightSide || chip.head == leftSide || chip.head == rightSide {
				validMoves = append(validMoves, chip)
			}
		}

	} else {
		chip := g.getDoubleSix()
		validMoves = append(validMoves, chip)
	}

	validMovesLenght := len(validMoves)

	if validMovesLenght > 0 {
		if validMovesLenght == 1 {
			move = validMoves[0]
		} else {
			//Randomize the selected chip
			move = shuffleTiles(validMoves)[0]
		}
		g.currentPlayer().dropTile(move)
		g.updateBoard(move)

	} else {
		g.skipPlayerTurn()
	}
}

func (g *Game) playerWonRound() {
	fmt.Println(g.currentPlayer().Name, " Won the round")
}

func (g *Game) selectNexPlayer() {
	// Bump the current Player Index
	g.currentPlayerIdx = nextPlayerIdx(g.currentPlayerIdx)
}

func (g *Game) isGameNew() bool {
	return g.ValidMove == false && g.isNewRound == true
}

func (g *Game) skipPlayerTurn() {
	//Update the skipped players counter and print the MSG
	g.skippedPlayers++
	g.currentPlayer().skip()
}

func (g *Game) getDoubleSix() Tiles {
	var tiles Tiles
	for i, p := range g.players {
		hasDoubleSix, chip := p.hasDoubleSix()
		if hasDoubleSix {
			tiles = chip
			g.currentPlayerIdx = i
			break
		}
	}
	return tiles
}

func (g *Game) newRound() {
	// Deal chips to start new round
	cards := newChips()
	for i := range g.players {
		g.players[i].Chips = cards[i]
	}
	//Reset the values
	g.isNewRound = true
	g.RightSide = 0
	g.LeftSide = 0
	g.skippedPlayers = 0
}

func (g *Game) updateBoard(chip Tiles) {
	// Update both sides of the board on a new game or round
	if g.isGameNew() || g.isNewRound {
		g.RightSide = chip.head
		g.LeftSide = chip.tail
	}
	// on an on going game select the side of the board to be updated
	if !g.isGameNew() {
		// select the side of the board to be updated
		if g.RightSide == chip.head {
			g.RightSide = chip.tail
		} else if g.RightSide == chip.tail {
			g.RightSide = chip.head
		} else if g.LeftSide == chip.head {
			g.LeftSide = chip.tail
		} else if g.LeftSide == chip.tail {
			g.LeftSide = chip.head
		}
	}

	// update last valid player
	g.skippedPlayers = 0
	g.ValidMove = true
	g.isNewRound = false
	fmt.Println("Board state: left side: ", g.LeftSide, "right side: ", g.RightSide)
}

func (g *Game) isGameLocked() bool {
	return g.skippedPlayers == 4
}

func (g *Game) currentPlayerWon() bool {
	// If player has no more tiles left he won
	l := len(g.currentPlayer().Chips)
	return l == 0
}

func (g *Game) didTeam1Win() bool {
	//Odd = Team 1(1/3) Even = Team 2 (2/4)
	oddEven := (g.currentPlayerIdx + 1) % 2
	return !(oddEven == 0)
}

func (g *Game) getRoundPoints() int {
	var result int
	for _, p := range g.players {
		result += p.getPoints()
	}
	return result
}

func (g *Game) currentPlayer() *Player {
	return &g.players[g.currentPlayerIdx]
}

func (g *Game) nextPlayer() Player {
	next := nextPlayerIdx(g.currentPlayerIdx)
	return g.players[next]
}

func whoWonLockedGame(player Player, player2 Player) (int, int) {
	player1Points := player.getPoints()
	player2Points := player2.getPoints()
	fmt.Println("The game was locked by: ", player.Name)
	fmt.Println(player.Name, " has ", player1Points, " points against")
	fmt.Println(player2.Name, " with ", player2Points, " points")
	return player1Points, player2Points
}

func endGameMessage(n int, t1 int, t2 int) {
	fmt.Println("Game has Ended Team ", n, "Won ")
	fmt.Println("Final scoreboard")
	fmt.Println("Team 1 score: ", t1)
	fmt.Println("Team 2 score: ", t2)
}

func endRoundMessage(t1 int, t2 int) {
	fmt.Println("End of round scoreboard")
	fmt.Println("Team 1 score: ", t1)
	fmt.Println("Team 2 score: ", t2)
}

func nextPlayerIdx(n int) int {
	if n == 3 {
		return 0
	}
	return n + 1
}
