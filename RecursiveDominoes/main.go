package main

import (
	"fmt"
)

func main() {
	currentGame := initGame()
	fmt.Println("Game Started Good Luck!!!")
	startGame(currentGame, 0, 0)
	fmt.Println("Good game hope to see you soon")
}

func startGame(currentGame *Game, t1Points int, t2Points int) {
	const maxScore = 200
	team1Won := t1Points >= maxScore
	team2Won := t2Points >= maxScore
	gameEnded := team1Won || team2Won

	if gameEnded {
		var n int
		if team1Won {
			n = 1
		} else {
			n = 2
		}
		endGameMessage(n, t1Points, t2Points)
		return
	} else {
		if !currentGame.isGameNew() {
			endRoundMessage(t1Points, t2Points)
		}
		t1Points, t2Points = nextRound(currentGame, t1Points, t2Points)
	}
	startGame(currentGame, t1Points, t2Points)
}

func nextRound(currentGame *Game, t1Points int, t2Points int) (int, int) {
	var roundPoints int
	gameLocked := currentGame.isGameLocked()
	roundEnded := currentGame.currentPlayerWon()

	if gameLocked {
		currentPlayer := currentGame.currentPlayer()
		player2 := currentGame.nextPlayer()
		player1Points, player2Points := whoWonLockedGame(*currentPlayer, player2)

		// Player 1 losed set Next player as winner
		if player1Points >= player2Points {
			currentGame.selectNexPlayer()
		} // Implicit else Player 1 won

		roundEnded = true
	}

	if !roundEnded { //Keep playing

		currentGame.selectNexPlayer()
		currentGame.currentPlayerPlayOrSkip()
		return nextRound(currentGame, t1Points, t2Points)
	} else {
		//Collect round points and start new round
		currentGame.playerWonRound()
		roundPoints = currentGame.getRoundPoints()

		switch currentGame.didTeam1Win() {
		case true:
			fmt.Println("With ", roundPoints, "for Team 1")
			t1Points += roundPoints
		case false:
			fmt.Println("With ", roundPoints, "for Team 2")
			t2Points += roundPoints
		}
		currentGame.newRound()
	}
	return t1Points, t2Points
}
