package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	mapset "github.com/deckarep/golang-set/v2"
)
func calc_score_2(input []string) int {
    pack1 := input[0];
    pack2 := input[1];
    pack3 := input[2];

    pack1 = SwapCaseII(pack1);
    pack2 = SwapCaseII(pack2);
    pack3 = SwapCaseII(pack3);
    set1 := mapset.NewSet[int]();
    set2 := mapset.NewSet[int]();
    set3 := mapset.NewSet[int]();
    for _,v := range pack1 {
        value := int(v)-0x40;
        if value > 26 {
            value -= 6;
        }
        set1.Add(value);
    }
    for _,v := range pack2 {
        value := int(v)-0x40;
        if value > 26 {
            value -= 6;
        }
        set2.Add(value);
    }
    for _,v := range pack3 {
        value := int(v)-0x40;
        if value > 26 {
            value -= 6;
        }
        set3.Add(value);
    }

    setU := set1.Intersect(set2.Intersect(set3));
    fmt.Printf("value: %v\n",setU);
    val,_ := setU.Pop();
    return val;
}
func calc_score(input string) int {
    round1 := input[:len(input)/2];
    round2 := input[len(input)/2:];
    round1 = SwapCaseII(round1);
    round2 = SwapCaseII(round2);
    // set1 := mapset.NewSet[string]();
    // set2 := mapset.NewSet[string]();
    // for _,v := range round1 {
    //     set1.Add(string(v));
    // }
    // for _,v := range round2 {
    //     set2.Add(string(v));
    // }
    // setU := set1.Intersect(set2);
    set1 := mapset.NewSet[int]();
    set2 := mapset.NewSet[int]();
    for _,v := range round1 {
        value := int(v)-0x40;
        if value > 26 {
            value -= 6;
        }
        set1.Add(value);
    }
    for _,v := range round2 {
        value := int(v)-0x40;
        if value > 26 {
            value -= 6;
        }
        set2.Add(value);
    }
    setU := set1.Intersect(set2);
    val,_ := setU.Pop();
    return val;
}
func SwapCaseII(str string) string {
    return strings.Map(SwapRune, str)
}

// rune is variable-length and can be made up of one or more bytes.
// rune literals are mapped to their unicode codepoint.
// For example, a rune literal 'a' is a number 97.
// 32 is the offset of the uppercase and lowercase characters.
// So if you add 32 to 'A', you get 'a' and vice versa.
func SwapRune(r rune) rune {
	switch {
	case 'a' <= r && r <= 'z':
		return r - 'a' + 'A'
	case 'A' <= r && r <= 'Z':
		return r - 'A' + 'a'
	default:
		return r
	}
}

func main() {
	file, _ := os.Open("input");
	// file, _ := os.Open("test")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner2 := bufio.NewScanner(file)
	scanner2.Split(bufio.ScanLines)
    total := 0;
    var lines []string;
    total2 := 0;
	for scanner.Scan() {
		round_input := scanner.Text();
        lines = append(lines,round_input);
        total += calc_score(round_input);
        // fmt.Printf("source: %v\n",total)
	}

    for i,y := 0, 0; y<len(lines)/3;y++ {
        total2 += calc_score_2(lines[i:i+3])
        i+=3;
        // .Printf("scanner %d, ans2: %v",len(lines) , lines[i:i+3])
    }

    // total2 += calc_score_2(round_input);
	fmt.Printf("ans1: %d, ans2: %d", total, total2)

}
