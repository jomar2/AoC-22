package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
    "regexp"
)
func calc_score_2(input []string) int {
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
func parse_input(input []string) [][]string {
    nr_bins := len(strings.Fields(input[0]));
    // fmt.Printf("\nnr :%v\n",nr_bins);
    output := make([][]string,nr_bins);
    // var crates []string
    re, _ := regexp.Compile(`[^\w]`)
    re2, _ := regexp.Compile(`\s{4}`)
    for _, crates := range input[1:] {
        crates = re2.ReplaceAllString(crates, "[x]")
        // fmt.Printf("\n%v",crates);
        crates = re.ReplaceAllString(crates, "")
        for bin_nr := 0;bin_nr < nr_bins; bin_nr++ {
            crate := string(crates[bin_nr])
            // fmt.Printf("ehh :%v at nr: %d\n",crate,bin_nr);
            if crate != "x" {
                output[bin_nr] = append(output[bin_nr], crate);
            }
        }
        // tmp := strings.Fields(crates);
        // fmt.Printf("\ncrate :\n%v",output);

    }
    count := 0;
    for _,bin := range output {
        count += len(bin)
    }
    fmt.Printf("count:\n%v\n", count)
	// fmt.Printf("\nsize:\n%v", len(bins))
	// fmt.Printf("\ncrates:\n%v", crates);
    // for _,bin := range bins {
    //     ind,_ : = strconv.Atoi(bin);
    //     output[ind] = ;
    // }
    // round1 := strings.FieldsFunc(input,Split);
    // set1 := mapset.NewSet[int]();
    // set2 := mapset.NewSet[int]();
    //
    // for _,val := range makeRangeStr(round1[0],round1[1]) {
    //     set1.Add(val);
    // }
    // for _,val := range makeRangeStr(round1[2],round1[3]) {
    //     set2.Add(val);
    // }
    // contains := set1.IsSubset(set2) || set2.IsSubset(set1);
    // overlap := len(set1.Intersect(set2).ToSlice()) != 0;
    //
    return output;
}
func ReverseSlice[T comparable](s []T) {
    sort.SliceStable(s, func(i, j int) bool {
        return i > j;
    })
}
func print_ans1(data [][]string) string{
    ans1 := "";
    for _,crates := range data {
        top_crate := crates[len(crates)-1]
        // fmt.Printf("bin %d : %s \n",bin,top_crate)
        ans1+=top_crate;
    }
    return ans1;
}
func fancy_moves(start [][]string, moveslist []string) [][]string {
    re, _ := regexp.Compile(`\D`);
    for _,moves := range moveslist {
        // fmt.Printf("start move: %v\n",moves)
        movesStr := strings.Fields(re.ReplaceAllString(moves," "));
        amount,_ := strconv.Atoi(movesStr[0]);
        from,_ := strconv.Atoi(movesStr[1]);
        to,_ := strconv.Atoi(movesStr[2]);
        from -= 1;
        to -=1;
        // fmt.Printf("move:\n%v: %v\n", movesStr,start)
        // crates := start[from][len(start[from])-amount:];
        var crates []string;
        crates, start[from] = start[from][len(start[from])-amount:], start[from][:len(start[from])-amount];
        // ReverseSlice(crates); // Enable for part 1!
        start[to] = append(start[to],crates...);
        count := 0;
        for _,bin := range start {
            count += len(bin)
        }
        // fmt.Printf("count:\n%d\n", count)
        // fmt.Printf("moving crate(s):\n%v\n", crates)
        // fmt.Printf("tree:\n%v", start)
    }
    return start;
}
func main() {
	file, _ := os.ReadFile("input");
	// file, _ := os.ReadFile("test");
	input := string(file);
    inputParsed := strings.Split(input,"\n\n");
    start := strings.Split(inputParsed[0],"\n");
    moveslist := strings.Split(inputParsed[1],"\n");
    moveslist = moveslist[:len(moveslist)-1];
    ReverseSlice(start);
	fmt.Printf("start:\n%v\n", inputParsed[0])
	// fmt.Printf("\nstartInv:\n%v", start)
	// fmt.Printf("\nmoves:\n%v\n", len(moveslist))
    // var lines []string;
	// for scanner.Scan() {
		// input := scanner.Text();
    data := parse_input(start);
    data = fancy_moves(data,moveslist);
    ans1 := print_ans1(data);
	fmt.Printf("ans1: %s, ans2: %v", ans1, data)
    // Wrong: RQFCHNCVFM
    // also : RQCHNCVFM
    // 1: correct! :ZBDRNPMVH
    // 2: WDLPFNNNB
}
