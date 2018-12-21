package main

import "testing"

var testCases = []string{
	"^WNE$",
	"^ENWWW(NEEE|SSE(EE|N))$",
	"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$",
	"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$",
	"^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$",
}

var expected = []int{
	3,
	10,
	18,
	23,
	31,
}

func TestDay20Part1(t *testing.T) {
	for i := range testCases {
		actual := Day20Part1(testCases[i])
		if actual != expected[i] {
			t.Errorf("Part1 was incorrect, with %v got: %d, want: %d.", testCases[i], actual, expected[i])
		}
	}
}
