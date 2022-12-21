package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

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
func ReverseSlice[T comparable](s []T) {
    sort.SliceStable(s, func(i, j int) bool {
        return i > j;
    })
}

func contains(data []int, elem int) bool {
    for _,val := range data {
        if val == elem {
            return true;
        }
    }
    return false;
}

func FindMarker(input string,chunkSize int) int {
    for ind := 0; ind<len(input)-chunkSize; ind++ {
        fmt.Printf("%s\n",input[ind:ind+chunkSize]);
        var chunk []int;
        unique := true;
        for _,char := range input[ind:ind+chunkSize] {
           if contains(chunk,int(char)) {
               unique = false;
               break;
           }
           chunk = append(chunk,int(char));
           fmt.Printf("%v\n",chunk);
        }
        if unique {
            return ind+chunkSize;
        }
    }
    return 0;
}
func main() {
    // testx := "mjqjpqmgbljsphdztnvjfqwrcgsmlb";
    // testx := "bvwbjplbgvbhsrlpgdmjqwftvncz"
    // testx := "nppdvjthqldpwncqszvftbrmjlhg"
    // testx := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
    // testx := "zcfzfwzzqf+4rljwzlrfnpqdbhtmscgvjw"
    // input := testx;
    file, _ := os.ReadFile("input");
	// file, _ := os.ReadFile("test");
	input := string(file);
    firstMarker := FindMarker(input,4);
    secondMarker := FindMarker(input,14);
    // data = fancy_moves(data,moveslist);
    // ans1 := print_ans1(data);
	fmt.Printf("ans1: %d, ans2: %v", firstMarker,secondMarker)
    // Wrong: RQFCHNCVFM
    // also : RQCHNCVFM
    // 1: correct! :ZBDRNPMVH
    // 2: WDLPFNNNB
}
