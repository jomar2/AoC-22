package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func calc_score_2nd(opp, me string) int {
	match := me + opp
	score := 0
    // fmt.Printf("me: %s\n", me)
	switch me {
	case "X":
		score += 0
	case "Y":
		score += 3
	case "Z":
		score += 6
	}
	// fmt.Printf("match: %s\n", match)
    // X,A = Rock Y,B Paper, Z,C Scissor
	switch match {
	case "XA":
		score += 3
	case "XB":
		score += 1
	case "XC":
		score += 2
	case "YA":
		score += 1
    case "YB":
        score += 2
    case "YC":
        score += 3
    case "ZA":
        score += 2
    case "ZB":
        score += 3
    case "ZC":
        score += 1
	}
	// fmt.Printf("match score: %d\n", score)
	return score;
}

func calc_score(opp, me string) int {
	match := me + opp
	score := 0
    // fmt.Printf("me: %s\n", me)
	switch me {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}
	// fmt.Printf("match: %s\n", match)
    // X,A = Rock Y,B Paper, Z,C Scissor
	switch match {
	case "XA":
		score += 3
	case "XB":
		score += 0
	case "XC":
		score += 6
	case "YA":
		score += 6
    case "YB":
        score += 3
    case "YC":
        score += 0
    case "ZA":
        score += 0
    case "ZB":
        score += 6
    case "ZC":
        score += 3
	}
	// fmt.Printf("match score: %d\n", score)
	return score;
}

func main() {
	file, _ := os.Open("input");
	// file, _ := os.Open("test")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	score := 0
	score2 := 0
	for scanner.Scan() {
		round_input := scanner.Text()
		round := strings.Fields(round_input)
		score += calc_score(round[0], round[1])
		score2 += calc_score_2nd(round[0], round[1])
        // fmt.Printf("opp: %s, me: %s, Score: %d\n",round[0],round[1], score)
	}

	fmt.Printf("ans1: %d, ans2: %d", score, score2)

}
