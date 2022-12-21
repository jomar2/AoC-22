package main

import (
	"fmt"
	// "math"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
	// mapset "github.com/deckarep/golang-set/v2"
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

func checkCycles(puter *computer){
    cycl := puter.cycles%40;
    reg := puter.register;
    // fmt.Printf("Cycles: %d, register: %d\n",cycl,reg);

    if (cycl== reg || cycl == reg-1 || cycl == reg+1) {
        // fmt.Printf("---Cycles: %d, register: %d\n",cycl,reg);
        puter.screen+="#";
    } else {
        puter.screen+=" ";
    }
    if cycl == 39 {
        puter.screen +="\n";
    }
    if slices.Contains(puter.strValues[:],puter.cycles) {
        strength := puter.cycles*puter.register;
        puter.strengthSum += strength;
        fmt.Printf("Cycles: %d, strength: %d\n",puter.cycles, strength);
    }
}

func calc(cmd string, value int,puter *computer){
    ADDTIME := 1;
    checkCycles(puter);
    switch (cmd) {
        case "addx": {
            for t:=0;t<ADDTIME;t++{
                puter.cycles++;
                checkCycles(puter);
            }
            puter.cycles++;
            puter.register += value;
        }
        case "noop": {
            puter.cycles++;
        }
    }
}

type computer struct {
    register int;
    cycles int;
    strengthSum int;
    strValues []int;
    screen string;
}

func main() {
    file, _ := os.ReadFile("input");
    // file, _ := os.ReadFile("test");
	// file, _ := os.ReadFile("test");
    // bodySize := 2;
    input := string(file);
    lines := strings.Split(input,"\n");
    puter := computer{register:1, cycles:0, strValues: []int{20,60,100,140,180,220}};
    for _,instr := range lines {
        inst := strings.Fields(instr);
        l := len(inst);
        if l == 0{
            continue;
        }
        cmd := inst[0];
        val := 0;
        if l > 1 {
            val = atoi(inst[1]);
        }
        calc(cmd,val,&puter);
       // fmt.Printf("cmd: %s, param: %v\n",cmd,val);
    }
    // baseAmount := 2*size+2*(size-2);
    fmt.Printf("ans1: %v, ans2: %v, tot: \n%v",puter.strengthSum , 0, puter.screen);// ,len(tailP))
}
