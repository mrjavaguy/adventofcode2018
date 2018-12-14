package main

import (
	"adventofcode2018/input"
	"fmt"
	"sort"
	"strings"
)

type cart struct {
	position         complex128
	direction        complex128
	intersectionMode complex128
	crashed          bool
}

type tracks map[complex128]rune

var cartMap map[complex128]*cart = map[complex128]*cart{}

func buildTrack(lines []string) (tracks tracks, carts []*cart) {
	tracks = map[complex128]rune{}
	carts = []*cart{}
	for y, s := range lines {
		for x, r := range s {
			position := complex(float64(x), float64(y))
			i := strings.IndexRune("^>v<", r)
			if i >= 0 {
				var direction complex128
				var trackType rune
				switch i {
				case 0:
					direction = -1i
					trackType = '|'
				case 1:
					direction = 1
					trackType = '-'
				case 2:
					direction = 1i
					trackType = '|'
				case 3:
					direction = -1
					trackType = '-'
				}
				cart := &cart{direction: direction, position: position, intersectionMode: -1i}
				carts = append(carts, cart)
				cartMap[position] = cart
				tracks[position] = trackType
			} else if r != ' ' {
				tracks[position] = r
			}
		}
	}
	return
}

func changeDirectionMaybe(cart *cart, trackType rune) {
	switch trackType {
	case '\\':
		cart.direction = complex(imag(cart.direction), real(cart.direction))
	case '/':
		cart.direction = complex(-imag(cart.direction), -real(cart.direction))
	case '+':
		cart.direction *= cart.intersectionMode
		switch cart.intersectionMode {
		case -1i:
			cart.intersectionMode = 1
		case 1:
			cart.intersectionMode = 1i
		case 1i:
			cart.intersectionMode = -1i
		}
	}
}

func Day13Part1(lines []string) (x, y int) {
	x = -1
	y = -1
	tracks, carts := buildTrack(lines)

	for {
		sort.Slice(carts, func(i, j int) bool {
			return (imag(carts[i].position) < imag(carts[j].position)) || ((imag(carts[i].position) == imag(carts[j].position)) && (real(carts[i].position) < real(carts[j].position)))
		})
		for _, c := range carts {
			delete(cartMap, c.position)
			c.position += c.direction

			if _, exist := cartMap[c.position]; exist {
				return int(real(c.position)), int(imag(c.position))
			}
			cartMap[c.position] = c
			trackType := tracks[c.position]
			changeDirectionMaybe(c, trackType)
		}
	}
}

func Day13Part2(lines []string) (x, y int) {
	x = -1
	y = -1
	cartMap = map[complex128]*cart{}
	tracks, carts := buildTrack(lines)
	for len(cartMap) > 1 {
		sort.Slice(carts, func(i, j int) bool {
			return (imag(carts[i].position) < imag(carts[j].position)) || ((imag(carts[i].position) == imag(carts[j].position)) && (real(carts[i].position) < real(carts[j].position)))
		})
		for _, c := range carts {
			if c.crashed {
				continue
			}
			delete(cartMap, c.position)
			c.position += c.direction

			if crash, exist := cartMap[c.position]; exist {
				delete(cartMap, crash.position)
				fmt.Println("Crash at:", real(c.position), imag(c.position), len(cartMap))
				crash.position, c.position = 0, 0
				crash.crashed, c.crashed = true, true
				continue
			}
			cartMap[c.position] = c
			trackType := tracks[c.position]
			changeDirectionMaybe(c, trackType)
		}
	}

	for p := range cartMap {
		return int(real(p)), int(imag(p))
	}

	return
}

func main() {
	lines, _ := input.FileToLines("day13/input13.txt")
	x, y := Day13Part1(lines)
	fmt.Println(x, y)
	x, y = Day13Part2(lines)
	fmt.Println(x, y)
}
