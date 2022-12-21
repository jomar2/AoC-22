package main

import (
	"fmt"
	"math"
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
type coord struct {
    x int;
    y int;
}
type info struct {
    start coord;
    end coord;
    limits coord;
}
func parseMap(in string) (map[coord]int, info) {

    start := coord{};
    end := coord{};
    x_max := 0;
    lines := strings.Split(in,"\n");
    y_max := len(lines)-1;
    output := make(map[coord]int,0);
    for y,row := range lines {
        if x_max == 0 {
            x_max = len(row)-1;
        }
        for x,val := range row {
            if val == 'S' {
                start.x=x;
                start.y=y;
                val = 'a';
            } else if val == 'E' {
                end.x=x;
                end.y=y;
                val = 'z';
            }
            output[coord{x:x,y:y}] = int(byte(val));
        }
    }
    fmt.Printf("\n");
    return output, info{start: start,end:end, limits: coord{x: x_max,y:y_max}};

}
func reconstruct_path(cameFrom map[coord]coord, current coord) []coord {
    total_path := []coord{current};
    c , notDone := cameFrom[current];
    for ;notDone;{
        current := c;
        total_path = append(total_path,current);
        c , notDone = cameFrom[current];
    }
    return total_path
}
func getLowestScore(in map[coord]int) coord {
    lowestScore := math.MaxInt16;
    lowestCoord := coord{};
    for c,val := range in {
        if val < lowestScore {
            lowestScore = val;
            lowestCoord = c;
        }
    }
    return lowestCoord;
}
func getNeighbours(current, maxValue coord) []coord {
    neigh := []coord{};
    if current.x == 0 {
        neigh = append(neigh,coord{x: current.x+1,y:current.y});
    } else if current.x == maxValue.x {
        neigh = append(neigh,coord{x: current.x-1,y:current.y});
    } else {
        neigh = append(neigh,coord{x: current.x+1,y:current.y});
        neigh = append(neigh,coord{x: current.x-1,y:current.y});
    }
    if current.y == 0 {
        neigh = append(neigh,coord{x: current.x,y:current.y+1});
    } else if current.y == maxValue.y {
        neigh = append(neigh,coord{x: current.x,y:current.y-1});
    } else {
        neigh = append(neigh,coord{x: current.x,y:current.y+1});
        neigh = append(neigh,coord{x: current.x,y:current.y-1});
    }
    // neigh_out := []coord{};
    // for _,v := range neigh {
    //     _,alreadyChecked := checked[v]; if !alreadyChecked {
    //         neigh_out = append(neigh,v);
    //     }
    // }
    return neigh;
}
func cost(current,neigh coord,myMap map[coord]int) int {
    diff := myMap[neigh] - myMap[current];
    if diff > 1 {
        return math.MaxInt32;
    } else {
        return 1;
    }
}
func cost_b(current,neigh coord,myMap map[coord]int) int {
    diff := myMap[neigh] - myMap[current];
    if diff > 0 {
        return math.MaxInt32;
    } else {
        return 1;
    }
}
// A* finds a path from start to goal.
// h is the heuristic function. h(n) estimates the cost to reach goal from node n.
func A_Star(values info, heightMap map[coord]int) []coord {
    // The set of discovered nodes that may need to be (re-)expanded.
    // Initially, only the start node is known.
    // This is usually implemented as a min-heap or priority queue rather than a hash-set.
    start := values.start;
    goal := values.end;

    openSet := make(map[coord]int,0);
    openSet[start] = 0;
    checkedSet := openSet;
    // For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
    // to n currently known.
    cameFrom := map[coord]coord{};
    // For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
    gScore := make(map[coord]int,len(heightMap));
    for c := range heightMap {
        gScore[c] = math.MaxInt16;
    }
    gScore[start] = 0;

    // For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
    // how cheap a path could be from start to finish if it goes through n.
    fScore := map[coord]int{};
    // for c := range heightMap {
    //     fScore[c] = math.MaxInt16;
    // }
    fScore[start] = 0;
    rounds := 0;
    lastCoord := coord{};
    for ;len(openSet)>0; {
        // fmt.Printf("did this at least once..\n");
        // fmt.Printf("round :%d , length: %d\n",i,len(openSet));
        // This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
        // current := the node in openSet having the lowest fScore[] value
        current := getLowestScore(fScore);
        lastCoord = current;
        // fmt.Printf("lowest Score coord: %v\n",current);
        if current == goal {
            return reconstruct_path(cameFrom, current)
        }
        delete(openSet,current);
        delete(fScore,current);
        for _,neighbor := range getNeighbours(current,values.limits) {
            // fmt.Printf("checking: %v:%s\n",neighbor,string(heightMap[neighbor]));
            // d(current,neighbor) is the weight of the edge from current to neighbor
            // tentative_gScore is the distance from start to the neighbor through current
            currentCost := cost(current,neighbor,heightMap);
            tentative_gScore := gScore[current] + currentCost;
            if tentative_gScore < gScore[neighbor] {
            // fmt.Printf("adding: %v\n",neighbor);
                // This path to neighbor is better than any previous one. Record it!
                cameFrom[neighbor] = current;
                gScore[neighbor] = tentative_gScore;
                fScore[neighbor] = tentative_gScore + 1;
                // fmt.Printf("costs: %d, %d\n",currentCost,tentative_gScore);
                _,exist := openSet[neighbor]; if !exist {
                    // fmt.Printf("Actually adding: %v score: %d\n",neighbor,tentative_gScore);
                    openSet[neighbor] = 1;
                    checkedSet[neighbor] = 1;
                }
            }
        }
        rounds++;
    }
    fmt.Printf("could not find goal..rounds: %d\n",rounds);
    // Open set is empty but goal was never reached
    return reconstruct_path(cameFrom, lastCoord);
}
func A_Star_StartLevel(values info, heightMap map[coord]int) []coord {
    start := values.start;

    openSet := make(map[coord]int,0);
    openSet[start] = 0;
    output := []coord{};
    gScore := make(map[coord]int,len(heightMap));
    for c := range heightMap {
        gScore[c] = math.MaxInt16;
    }
    gScore[start] = 0;

    fScore := map[coord]int{};
    fScore[start] = 0;
    rounds := 0;
    for ;len(openSet)>0; {
        current := getLowestScore(fScore);
        delete(openSet,current);
        delete(fScore,current);
        for _,neighbor := range getNeighbours(current,values.limits) {
            currentCost := cost_b(current,neighbor,heightMap);
            tentative_gScore := gScore[current] + currentCost;
            if tentative_gScore < gScore[neighbor] {
            // fmt.Printf("adding: %v\n",neighbor);
                // This path to neighbor is better than any previous one. Record it!
                gScore[neighbor] = tentative_gScore;
                fScore[neighbor] = tentative_gScore + 1;
                // fmt.Printf("costs: %d, %d\n",currentCost,tentative_gScore);
                _,exist := openSet[neighbor]; if !exist {
                    // fmt.Printf("Actually adding: %v score: %d\n",neighbor,tentative_gScore);
                    openSet[neighbor] = 1;
                    output = append(output,neighbor);
                }
            }
        }
        rounds++;
    }
    fmt.Printf("could not find goal..rounds: %d -- val: %d\n",rounds,len(output));
    // Open set is empty but goal was never reached
    return output;
}
func printMap(in *map[coord]int,infoP info) {
    tmp := "";
    keys := []coord{};
    for k := range *in {
        keys = append(keys,k);
    }
    sort.SliceStable(keys, func(i, j int) bool{
        return keys[i].y < keys[j].y || (keys[i].y == keys[j].y && keys[i].x < keys[j].x);
    })
    fmt.Printf("%v\n", keys);
    i := 0;
    for _,c := range keys {
        i++;
        fmt.Printf("Val %s ",string((*in)[c]));
        tmp += string((*in)[c]);
        if i%(infoP.limits.x+1) == 0 {
            tmp += "\n";
        }
    }
    fmt.Printf("\n%s",tmp);
}
func main() {
    file, _ := os.ReadFile("input");
    // file, _ := os.ReadFile("test");
    // bodySize := 2;
    input := string(file);
    myMap, infoPoints :=parseMap(input);
    // myMap := strings.Split(input,"\n");
    // bestPath := A_Star(infoPoints,myMap); // Part A
    startPoints := A_Star_StartLevel(infoPoints,myMap);
    lowestCount := math.MaxInt;

    for _,c := range startPoints {
        infoPoints.start = c;
        bestPath := A_Star(infoPoints,myMap);
        if len(bestPath) < lowestCount {
            lowestCount = len(bestPath);
        }
    }
    // fmt.Printf("path: %v\n",bestPath);
    // for n,s := range bestPath {
    //     myMap[s] = int(byte('0'+n%10));
    // }

    // printMap(&myMap,infoPoints);

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
    fmt.Printf("ans1: %v, ans2: %v, tot: \n%v",528 /*len(bestPath)-1*/, lowestCount-1, 0)
}
