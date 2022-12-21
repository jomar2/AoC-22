package main

import (
	"fmt"
	"regexp"
	// "math"
	"os"
	"sort"
	"strconv"
	"strings"
	// "golang.org/x/exp/slices"
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
    val,_ := strconv.Atoi(strings.Trim(input," "));
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
func atoiStr(input []string) []int64 {
    var output []int64;
    for _,val := range input {
        output = append(output,int64(atoi(val)));
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
func PopInt(list []int) (int,[]int, bool) {
    empty := len(list)<=1;
    if empty {
        x, a := list[0], make([]int,0);
        return x, a ,empty;
    } else {
        x, a := list[0], list[1:];
        return x, a ,empty;
    }
}


type operation func(int64) int64;
type monkey struct {
    id int;
    items *[]int64;
    op operation;
    testValue int64;
    tTarget int;
    fTarget int;
    inspected int64;
}
    // re, _ := regexp.Compile(`[^\w]`)
    // re2, _ := regexp.Compile(`\s{4}`)
    // for _, crates := range input[1:] {
    //     crates = re2.ReplaceAllString(crates, "[x]")
    //     crates = re.ReplaceAllString(crates, "")
    //     }
    // }


func getOp(instr string) func(int64) int64 {
    ins := strings.Fields(instr);
    operand := ins[1];
    switch (operand) {
        case "*": {
            if ins[2] == "old" {
                return func(old int64) int64 {
                    return old*old;
                }
            } else {
                val := int64(atoi(ins[2]));
                return func(old int64) int64 {
                    return old*val;
                }
            }
        }
        case "+": {
            if ins[2] == "old" {
                return func(old int64) int64 {
                    return old+old;
                }
            } else {
                val := int64(atoi(ins[2]));
                return func(old int64) int64 {
                    return old+val;
                }
            }
        }
        case "-": {
            if ins[2] == "old" {
                return func(old int64) int64 {
                    return 0;
                }
            } else {
                val := int64(atoi(ins[2]));
                return func(old int64) int64 {
                    return old-val;
                }
            }
        }
        case "/": {
            if ins[2] == "old" {
                return func(old int64) int64 {
                    return 1;
                }
            } else {
                val := int64(atoi(ins[2]));
                if(val == 0) {
                    break;
                }
                return func(old int64) int64 {
                    return old/val;
                }
            }
        }
    }
    return func(in int64) int64 {return 0};
}

func parseMonkey(inp string) monkey {
    in := strings.Split(inp,"\n");
    re_id, _ := regexp.Compile(`[\d+]`)
    // re_id, _ := regexp.Compile(`[^\w]`)
    id := atoi(re_id.FindString(in[0]));
    // item := atoiStr(strings.Split(strings.SplitN(in[1]," ",1)[0],","));
    item := atoiStr(strings.Split(strings.Split(in[1],":")[1],","));
    operation := strings.SplitN(in[2],"=",2)[1];
    op := getOp(operation);
    // op := getOp("new= old + 2");
    test := int64(atoi(strings.SplitN(in[3],"by",2)[1]));
    tT :=atoi(re_id.FindString(in[4]));
    fT :=atoi(re_id.FindString(in[5]));
    outputM := monkey{id: id,items: &item, op: op,fTarget: fT,tTarget:tT,testValue: test,inspected: 0};
    return outputM;
}

func getMoulus(in *[]monkey) int64 {
    var output int64;
    output = 1;
    for _,tv := range *in {
        output *= int64(tv.testValue);;
    }
    return output;
}
func main() {
    file, _ := os.ReadFile("input");
    // file, _ := os.ReadFile("test");
	// file, _ := os.ReadFile("test");
    // bodySize := 2;
    input := string(file);
    monkeyInput := strings.Split(input,"\n\n");
    g := []monkey{};
    gang := &g;
    rounds := 10000;
    for _,mon := range monkeyInput {
        //     inst := strings.Fields(instr);
        fmt.Printf("inp: \n%s, monk: %v\n",mon,parseMonkey(mon));
        *gang = append(*gang,parseMonkey(mon));
    }
    inspect := make([]int,len(*gang));
    modulus := getMoulus(gang);
    for r := 0;r<rounds;r++ {
        for i,mon := range *gang {
            var val int64;
            for _,old := range *mon.items {
                inspect[i] += 1;
                val = mon.op(old);
                *mon.items = (*mon.items)[1:];
                var target int;
                newval := val % modulus;
                if val % mon.testValue == 0 {
                    target = mon.tTarget;
                } else {
                    target = mon.fTarget;
                }
                *(*gang)[target].items = append(*(*gang)[target].items,newval);
            }
        // fmt.Printf("inspect: %d\n",mon.inspected);
            // mon.items = nil;
        }
    }
    // business := []int{};
    // for _,mon := range *gang {
    //     fmt.Printf("insp: %d\n", mon.inspected);
    //     business = append(business,mon.inspected);
    // };

    for i,v := range inspect {
        fmt.Printf("mon %d: %d\n", i,v);
    }
    sort.Ints(inspect);
    l := len(inspect);
    fmt.Printf("l: %d\n", l);
    ans1 := inspect[l-2]*inspect[l-1];
    //     l := len(inst);
    //     if l == 0{
    //         continue;
    //     }
    //     cmd := inst[0];
    //     val := 0;
    //     if l > 1 {
    //         val = atoi(inst[1]);
    //     }
    //     calc(cmd,val,&puter);
    //    // fmt.Printf("cmd: %s, param: %v\n",cmd,val);
    // baseAmount := 2*size+2*(size-2)
    // Too big: 32390460690
    // too high : 18377326293

    fmt.Printf("ans1: %v, ans2: %v, tot: \n%v",ans1, (*gang)[1].items, gang)
}
