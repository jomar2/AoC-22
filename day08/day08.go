package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
func makeRangeRev(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = max - i
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
func atoiStrVec(input []string) [][]int {
    var output [][]int;
    for _,inp := range input {
        var tmp []int;
        for _,val := range strings.Split(inp,"") {
            tmp = append(tmp,atoi(val));
        }
        output = append(output,tmp);
    }
    return output;
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

func Pop(lines []string) (string,[]string, bool) {
    empty := len(lines)<=1;
    if empty {
        x, a := lines[0], make([]string,0);
        return x, a ,empty;
    } else {
        x, a := lines[0], lines[1:];
        return x, a ,empty;
    }
}
func checkTree(val,x,y int,input [][]int,maxSize int) (bool, int) {
    if(x==0 || y== 0 || x==maxSize-1 || y == maxSize-1) {
        // fmt.Printf("Edge Value: %dx%d\n",x,y);
        return true,0;
    }
    ok := true;
    var nrTrees []int;
    dirTrees := 0;
    for _,i := range makeRangeRev(0 ,x-1) {
        //fmt.Printf("check Left Value: %dx%d: %d>=%d\n",i,y,input[y][i],val);
        dirTrees++;
        if (input[y][i]>=val){
            ok = false;
            break;
        }
    }
    // fmt.Printf("check Left Value: %dx%d: %d: Trees: %d\n",x,y,input[y][x],dirTrees);
    nrTrees = append(nrTrees,dirTrees);
    dirTrees =0;
    nrOks := 0;
    if (ok) {
        nrOks++;
    }
    ok = true;
    //fmt.Printf("Left Nok!\n");
    for _,i := range makeRange(x+1,maxSize-1) {
        //fmt.Printf("check Value: %dx%d: %d>=%d\n",i,y,input[y][i],val);
        dirTrees++;
        if (input[y][i]>=val){
            ok = false;
            break;
        }
    }
    nrTrees = append(nrTrees,dirTrees);
    dirTrees =0;
    if (ok) {
        nrOks++;
    }
    ok = true;
    // //fmt.Printf("right Nok!\n");
    // ok = true;
    for _,i := range makeRangeRev(0 ,y-1) {
        dirTrees++;
        if (input[i][x]>=val){
            ok = false;
            break;
        }
    }
    nrTrees = append(nrTrees,dirTrees);
    dirTrees =0;
    if (ok) {
        nrOks++;
    }
    ok = true;
    // //fmt.Printf("Up Nok!\n");
    // ok = true;
    for _,i := range makeRange(y+1,maxSize-1) {
        dirTrees++;
        if (input[i][x]>=val){
            ok = false;
            //fmt.Printf("Down Nok!\n");
            break;
        }
    }
    nrTrees = append(nrTrees,dirTrees);
    dirTrees = 0;
    totTrees :=1;
    // fmt.Printf("trees:%v\n",nrTrees);
    for _,trees := range nrTrees {
        totTrees *= trees;
    }

    return nrOks>0, totTrees;
}
func checkArray(input [][]int,maxSize int) (int,int) {
    visible := 0;
    totT := []int{};
    for y,row := range input {
        for x,val := range row {
            vis,nrTrees := checkTree(val,x,y,input,maxSize);
            if vis {
                visible++;
                if nrTrees>0 {
                    totT = append(totT,nrTrees);
                }
            }
        }
    }
    sort.Ints(totT);
    // fmt.Printf("treesList %v", totT);
    return visible,totT[len(totT)-1];
}

func main() {
    file, _ := os.ReadFile("input");
	// file, _ := os.ReadFile("test");
    input := string(file);
    lines := strings.Split(input,"\n");
    size := len(lines[0]);
    // baseAmount := 2*size+2*(size-2);
    data := atoiStrVec(lines);
    ans1,ans2:=checkArray(data,size);
    fmt.Printf("ans1: %v, ans2: %v, tot: \n%v",ans1,ans2,0)
}
