package main

import (
	"fmt"
	// "regexp"
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
   // fmt.Printf("min,max: %v,%v\n", min,max);
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
type coord struct {
    x int;
    y int;
}
const Y_SIZE = 200;
const X_SIZE = 1000;
func main() {
    file, _ := os.ReadFile("input");
    // file, _ := os.ReadFile("test");
    // file, _ := os.ReadFile("test2");
    // bodySize := 2;

    input := string(file);
    myMap := [Y_SIZE][X_SIZE]string{};
    lines := strings.Split(input,"\n");
    floorLevel := 0;
    for i,dLines := range lines {
        lineParts := []coord{};
        // dLines := strings.Split(c,"\n");
        if dLines == "" {
            break;
        }
        parts := strings.Split(dLines," -> ");
        for _,p := range parts {
            tmp := strings.Split(p,",");
            c := coord{x:atoi(tmp[0]),y:atoi(tmp[1])};
            if c.y > floorLevel {
                floorLevel = c.y;
            }
            lineParts = append(lineParts,c);
        }
        printLine(lineParts,&myMap);
        // fmt.Printf("\nLeft: \n%v\n",parts[0]);
        // fmt.Printf("Right: \n%v\n\n",(parts[1]));
        fmt.Printf("Done index %d, %v!\n",i,lineParts);
    }
    secondMap := myMap;
    sand_a := addSand(&myMap,coord{x:500,y:0}, floorLevel);
    addFloor(&secondMap,floorLevel+2);
    sand_b := addSand(&secondMap,coord{x:500,y:0}, floorLevel+3);

    // myMap := strings.Split(input,"\n");
    // printMap(&myMap,floorLevel);
    fmt.Printf("\n\n");
    // printMap(&secondMap,floorLevel+2);

    // bestPath := A_Star(infoPoints,myMap); // Part A

    fmt.Printf("ans1: %v, ans2: %v, tot: %v\n",sand_a, sand_b, floorLevel)
}

func addFloor(myMap *[Y_SIZE][X_SIZE]string, floor int) {
    for i:=0;i<len(myMap[0]);i++ {
        myMap[floor][i]="#";
    }
}

func addSand(myMap *[Y_SIZE][X_SIZE]string, start coord, floorLevel int) int{
    sand := 0;
    for done:=false;!done; {
        currPos := start;
        // fmt.Printf("start\n");

        for stationary:=false;!stationary; {
            // fmt.Printf("sand: %v\n",currPos);
            newSpace := myMap[currPos.y+1][currPos.x]
            level := currPos.y+1;
            if level > floorLevel {
                myMap[currPos.y][currPos.x] = "~";
                done = true;
                break;
            }
            if newSpace == "" {
                currPos.y++;
            } else if myMap[currPos.y+1][currPos.x-1] == ""{
                currPos.y++;
                currPos.x--;

            } else if myMap[currPos.y+1][currPos.x+1] == ""{
                currPos.y++;
                currPos.x++;
            } else {
                myMap[currPos.y][currPos.x] = "o";
                // fmt.Printf("Add sand at %v\n",currPos);
                sand++;
                stationary = true;
            }
        }
        if currPos == start {
            done=true;
        }
    }
    return sand;
}

func printMap(myMap *[Y_SIZE][X_SIZE]string,floor int) {
    fmt.Printf("\n\n");
    myMap[0][500] = "+";
    for i,y := range myMap {
        var line string;
        for _,x := range y {
            if x == "" {
                line+=".";
            } else {
                line+=x;
            }
        }
        fmt.Printf("%s\n",line);
        if i > floor {
            break;
        }
    }
    fmt.Printf("\n\n");
}

func printLine(lines []coord, myMap *[Y_SIZE][X_SIZE]string) {
    for i,l := range lines {
        if i == 0 {
            continue;
        }
        // fmt.Printf("line %d\n",i)
        start := lines[i-1];
        end := l
        dist :=0;
        direction := "+x";
        if start.x == end.x {
            dist = end.y-start.y;
            if dist > 0 {
                direction = "+y"
            } else {
                direction = "-y"
                dist *= -1;
            }
        } else {
            dist = end.x-start.x;
            if dist > 0 {
                direction = "+x"
            } else {
                direction = "-x"
                dist *= -1;
            }
        }
        // fmt.Printf("data %d,%v\nstart %v - end %v\n",dist,direction,start,end)

        switch direction {
        case "+x":
            for c := range makeRange(start.x,end.x){
                myMap[start.y][start.x+c] = "#";
            }
        case "-x":
            for c := range makeRange(end.x,start.x) {
                myMap[start.y][end.x+c] = "#";
            }
        case "+y":
            for c := range makeRange(start.y,end.y) {
                myMap[start.y+c][start.x] = "#";
            }
        case "-y":
            for c := range makeRange(end.y,start.y) {
                myMap[end.y+c][start.x] = "#";
            }
        }
    }
}
