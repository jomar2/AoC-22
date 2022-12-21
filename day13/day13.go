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

    // re2, _ := regexp.Compile(`\s{4}`)
    // for _, crates := range input[1:] {
    //     crates = re2.ReplaceAllString(crates, "[x]")
    //     crates = re.ReplaceAllString(crates, "")
    //     }
    // }
func trimOnce(in string) string {
    return "";

}

type returnType struct {
    done bool;
    correct bool;
}
func checkPair(left,right string) returnType {
    // Leftlist := (strings.SplitN(left,"]",2));
    fmt.Printf("Start Right: %s, Left: %s\n",right,left);
    re, _ := regexp.Compile(`\[(.*)\]$`)
    leftC := re.FindStringSubmatch(left);
    rightC := re.FindStringSubmatch(right);
    if len(leftC)==0 {
        return returnType{done:true,correct: true};
    } else if len(rightC)==0 {
        return returnType{done:true,correct: false};
    }
    leftVec := leftC;
    rightVec := rightC;
    leftChunk := leftC[1];
    rightChunk := rightC[1];
    returnVal := returnType{done:false};
    empty :=false;
    for ;!returnVal.done && !empty; {
        fmt.Printf("Loop %v:%v\n",leftChunk,rightChunk)
        // if !isLeftList {
            leftVec = strings.SplitN(leftChunk,",",2);
        // }
        leftVal := leftVec[0];
        if len(leftVec) > 1 {
            leftChunk = leftVec[1];
        } else {
            leftChunk = "";
            empty = true;
        }
        fmt.Printf("Lefts: %v -- %v\n",leftVal,leftChunk);
        // if !isRightList {
            rightVec = strings.SplitN(rightChunk,",",2);
        // }
        rightVal := rightVec[0];
        if len(rightVec) > 1 {
            rightChunk = rightVec[1];
        } else {
            rightChunk = "";
            empty = true;
        }
        isLeftList := re.MatchString(leftVal) || string(leftVal[0]) == "[";
        isRightList := re.MatchString(rightVal) || string(rightVal[0]) == "[";
        fmt.Printf("Rights: %v -- %v\n",rightVal,rightChunk);
        if isLeftList && isRightList {
            returnVal = checkPair(leftVal,rightVal);
        }
        fmt.Printf("Start Loop %v:%v, %v:%v\n",leftVal,rightVal,leftVec,rightVec)
        if (!isRightList && isLeftList) {
            fmt.Printf("Making right to a list\n");
            returnVal = checkPair(leftChunk,"["+rightVal+"]");
        } else if (isRightList && !isLeftList) {
            fmt.Printf("Making left to a list\n");
            returnVal = checkPair("["+leftVal+"]",rightChunk);
        } else if !isLeftList && isRightList {
            fmt.Printf("!left && right: Not Correct\n");
            returnVal = returnType{done:true,correct:false};
        }
        if len(leftVal)==0 {
            returnVal = returnType{done:true,correct: true};
        } else if len(rightVal)==0 {
           returnVal = returnType{done:true,correct: false};
        }
        if atoi(leftVal) == atoi(rightVal) {
            fmt.Printf("The Same, Continue\n");
            continue;
        } else if atoi(leftVal) < atoi(rightVal) {
            fmt.Printf("left < right: Correct\n");
            returnVal = returnType{done:true,correct:true};
        } else {
            returnVal = returnType{done:true,correct: false};
        }
        if leftVal == "" {
            returnVal = returnType{done:true,correct:true};
        }


        // if returnVal.done {
        //     return returnVal;
        // }
        // isLeftList = re.MatchString(leftVec[1]);
        // isRightList = re.MatchString(rightVec[1]);
        // if !isLeftList {
        //     leftVec = strings.SplitN(leftVal[1],",",2);
        // }
        // if !isRightList {
        //     rightVec = strings.SplitN(rightVal[1],",",2);
        // }
        // leftValue := leftVec[1];
        // rightValue := rightVec[1];
        //
        // if len(leftValue)>=len(rightValue) {
        //     return returnType{done:true,correct:false};
        // } else {
        //     // if rl == 0 {
        //     //     return returnType{done:true,correct:false};
        //     // }
        //     rightValue = rightVec[i];
        // }
        // fmt.Printf("current Loop %v, %v\n",leftValue,rightValue);
        //
        // if isLeftList && isRightList {
        //     leftVal := re.FindStringSubmatch(leftValue);
        //     rightVal := re.FindStringSubmatch(rightValue);
        // } else if !isLeftList && !isRightList {
        //     // leftVal := atoiStr(leftVec);
        //     fmt.Printf("current Inner %v, %v\n",leftVec,rightVec);
        //
        //     // fmt.Printf("Compare?\nLeftlist: %v\nRightList: %v\n",leftVal,rightVal)
    }
    // fmt.Printf("Leftlist: %v\nRightList: %v\n",leftVal,rightVal)
    return returnType{done:false,correct:true};
}
func main() {
    // file, _ := os.ReadFile("input");
    // file, _ := os.ReadFile("test");
    file, _ := os.ReadFile("test2");
    // bodySize := 2;
    input := string(file);
    chunks := strings.Split(input,"\n\n");
    correctId := []int{};
    for i,c := range chunks {
        parts := strings.Split(c,"\n");
        // fmt.Printf("\n\nLeft: \n%v\n",parts[0]);
        // fmt.Printf("Right: \n%v\n\n",(parts[1]));
        ans := checkPair(parts[0],parts[1]);
        fmt.Printf("Done index %d!\n",i);
        if ans.correct {
            // fmt.Printf("add!");
            correctId = append(correctId,i+1);
        }
    }
    sum :=0;
    for _,ind := range correctId{
        sum+=ind;
    }
    // myMap := strings.Split(input,"\n");
    // bestPath := A_Star(infoPoints,myMap); // Part A

    fmt.Printf("ans1: %v, ans2: %v, tot: \n%v",sum, correctId, 0)
}
