package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input");
	// file, _ := os.Open("test");
	    scanner := bufio.NewScanner(file);
    scanner.Split(bufio.ScanLines);

    var sumVec []int;
    var total = 0;
    // var max = 0;
	for scanner.Scan() {
		if scanner.Text() == "" {
            sumVec = append(sumVec,total);
            total = 0;
            continue;
        }
        nr,_ := strconv.Atoi(scanner.Text())
        total += nr;
	}
    sumVec = append(sumVec,total);

    sort.Sort(sort.Reverse(sort.IntSlice(sumVec)));
    num := [3]int{0, 1, 2};
    tot := 0;
    for ind := range num {
         fmt.Printf("value %v at %d\n",sumVec[ind],ind);
         tot += sumVec[ind];
    }
    fmt.Printf("all: - and max: %d and top3 totaled: %d",sumVec[0],tot)

}
