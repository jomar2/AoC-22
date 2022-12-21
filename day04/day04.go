package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	mapset "github.com/deckarep/golang-set/v2"
)
func calc_score_2(input []string) int {
    // pack1 := input[0];
    // pack2 := input[1];
    // pack3 := input[2];
    //
    // pack1 = SwapCaseII(pack1);
    // pack2 = SwapCaseII(pack2);
    // pack3 = SwapCaseII(pack3);
    // set1 := mapset.NewSet[int]();
    // set2 := mapset.NewSet[int]();
    // set3 := mapset.NewSet[int]();
    // for _,v := range pack1 {
    //     value := int(v)-0x40;
    //     if value > 26 {
    //         value -= 6;
    //     }
    //     set1.Add(value);
    // }
    // for _,v := range pack2 {
    //     value := int(v)-0x40;
    //     if value > 26 {
    //         value -= 6;
    //     }
    //     set2.Add(value);
    // }
    // for _,v := range pack3 {
    //     value := int(v)-0x40;
    //     if value > 26 {
    //         value -= 6;
    //     }
    //     set3.Add(value);
    // }
    //
    // setU := set1.Intersect(set2.Intersect(set3));
    // fmt.Printf("value: %v\n",setU);
    // val,_ := setU.Pop();
    return 0;
}

func Split(r rune) bool {
    return r == ',' || r == '-'
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    // fmt.Printf("range: %v\n", a);
    return a
}
func makeRangeStr(min, max string) []int {
    max_i := atoi(max);
    min_i := atoi(min);

    return makeRange(min_i,max_i);

}
func atoi(input string) int {
    val,_ := strconv.Atoi(input);
    return val;
}
func check_task_pair(input string) (bool,bool) {
    round1 := strings.FieldsFunc(input,Split);
    set1 := mapset.NewSet[int]();
    set2 := mapset.NewSet[int]();

    for _,val := range makeRangeStr(round1[0],round1[1]) {
        set1.Add(val);
    }
    for _,val := range makeRangeStr(round1[2],round1[3]) {
        set2.Add(val);
    }
    contains := set1.IsSubset(set2) || set2.IsSubset(set1);
    overlap := len(set1.Intersect(set2).ToSlice()) != 0;

    return contains,overlap;
}

func main() {
	file, _ := os.Open("input");
	// file, _ := os.Open("test")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
    total := 0;
    total2 := 0;
    // var lines []string;
	for scanner.Scan() {
		round_input := scanner.Text();
        contains, overlap :=check_task_pair(round_input);
        if(contains) {
            total++;
        }
        if(overlap) {
            total2++;
        }
	//         // lines = append(lines,round_input);
	//         // total += calc_score(round_input);
	//         // fmt.Printf("source: %v\n",total)
	}
	//
	//     for i,y := 0, 0; y<len(lines)/3;y++ {
	//         total2 += calc_score_2(lines[i:i+3])
	//         i+=3;
	//         // .Printf("scanner %d, ans2: %v",len(lines) , lines[i:i+3])
	//     }
	//
    // total2 += calc_score_2(round_input);
	fmt.Printf("ans1: %d, ans2: %d", total, total2)

}
