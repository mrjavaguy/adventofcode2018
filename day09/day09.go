package main

import (
	"container/ring"
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	part1 := Day09Part1("411 players; last marble is worth 72059 points")
	part2 := Day09Part2("411 players; last marble is worth 72059 points")
	fmt.Printf("Day 8 Part 1 %v, Part 2 %v", part1, part2)
}

func Day09Part1(game string) int {
	players, lastMarble := parseGame(game)
	return playGame(players, lastMarble)
}

func Day09Part2(game string) int {
	players, lastMarble := parseGame(game)
	lastMarble *= 100
	return playGame(players, lastMarble)
}

func parseGame(game string) (players int, lastMarble int) {
	re := regexp.MustCompile("(\\d+) players; last marble is worth (\\d+) points")
	matches := re.FindAllStringSubmatch(game, -1)
	players, _ = strconv.Atoi(matches[0][1])
	lastMarble, _ = strconv.Atoi(matches[0][2])
	return
}

func playGame(players, lastMarble int) int {
	circle := ring.New(1)
	circle.Value = 0

	playerScores := make([]int, players)

	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			circle = circle.Move(-8)
			removed := circle.Unlink(1)
			playerScores[i%players] += i + removed.Value.(int)
			circle = circle.Move(1)
		} else {
			circle = circle.Move(1)

			newMarbleLocation := ring.New(1)
			newMarbleLocation.Value = i

			circle.Link(newMarbleLocation)
			circle = circle.Move(1)
		}
	}

	var highScore int
	for _, score := range playerScores {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}
