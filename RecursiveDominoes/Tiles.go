package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tiles struct {
	head int
	tail int
}

func (c *Tiles) toString() string {
	return fmt.Sprintf("[  %d  |  %d  ]", c.head, c.tail)
}

// Returns a 2d array of tiles (easier for assigning them to player)
func newChips() [][]Tiles {
	var q []Tiles
	//Create Chips in order
	for i := 0; i <= 6; i++ {
		for j := 0; j <= 6; j++ {
			q = appendIfMissing(q, Tiles{i, j})
		}
	}
	// randomize Chips
	sorted := make([]Tiles, len(q))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(q))
	for i, randIndex := range perm {
		sorted[i] = q[randIndex]
	}
	// 4 Tiles stacks 1 per player
	chipStacks := make([][]Tiles, 4)
	for i := 0; i <= 3; i++ {
		chipStacks[i] = make([]Tiles, 7)
		for j := 0; j <= 6; j++ {
			chipStacks[i][j] = sorted[0]
			sorted = sorted[1:]
		}
	}
	return chipStacks
}

func appendIfMissing(slice []Tiles, c Tiles) []Tiles {
	for _, item := range slice {
		if item.head == c.head && item.tail == c.tail ||
			item.head == c.tail && item.tail == c.head {
			return slice
		}
	}
	return append(slice, c)
}

func shuffleTiles(q []Tiles) []Tiles {
	sorted := make([]Tiles, len(q))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(q))
	for i, randIndex := range perm {
		sorted[i] = q[randIndex]
	}
	return sorted
}
