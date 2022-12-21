package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
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
func moveSnek(snek *snekType, dx,dy, length int) {
    for i:=0;i<length;i++ {
        snek.body[0].x +=dx;
        snek.body[0].y +=dy;
        tailId := len(snek.body)-1;
        for _,j := range makeRange(1,tailId) {
            dx_ := (snek.body[j-1].x-snek.body[j].x);
            dy_ := (snek.body[j-1].y-snek.body[j].y);
            // fmt.Printf("i: %v, i-1:%v, dxdy:%d,%d\n",snek.body[j],snek.body[j-1],dx,dy)
            dx_l := math.Abs(float64(dx_));
            dy_l := math.Abs(float64(dy_));
            for ; dx_l > 1 || dy_l > 1; {
                // fmt.Printf("dx: x:%v,%v, y:%v,%v\n",dx_,dx_l,dy_,dy_l);
            //     if dx_ > 2 || dy_ > 2 {
            //         fmt.Printf("oh... x:%d, y:%d\n",dx_,dy_)
            //     }
               if dx_l >= 1 {
                   snek.body[j].x += dx_/int(dx_l);
                }
                if dy_l >= 1 {
                    snek.body[j].y += dy_/int(dy_l);
                }
            //     // fmt.Printf("dy big\n");
                dx_ = (snek.body[j-1].x-snek.body[j].x);
                dy_ = (snek.body[j-1].y-snek.body[j].y);
                dx_l = math.Abs(float64(dx_));
                dy_l = math.Abs(float64(dy_));
                if j == len(snek.body)-1 {
                    snek.tailPos = append(snek.tailPos,snek.body[tailId]);
                }
                // fmt.Printf("Move: (%dx%d), SnekHead: %v, Tail: %v\n",dx,dy,snek.body[0],snek.body[len(snek.body)-1]);
            }
            snek.tailPos = append(snek.tailPos,snek.body[tailId]);
        }
    }
}
func parseMove(snek *snekType, theMove string) {
    move := strings.Fields(theMove);
    dir:=move[0];
    l := atoi(move[1]);
    switch dir {
    case "R":
        moveSnek(snek,1 ,0 ,l);
    case "L":
        moveSnek(snek,-1 ,0 ,l);
    case "U":
        moveSnek(snek,0 ,1 ,l);
    case "D":
        moveSnek(snek,0 ,-1 ,l);
    }
    // fmt.Printf("Move: %s\n", theMove);
    printMap(snek);
}
type coord struct {
    x int;
    y int;
}
type snekType struct {
    body []coord;
    headPos []coord;
    tailPos []coord;
}
func unique(in []coord) []coord {
    sort.Slice(in[:],func (c1,c2 int) bool {
        if in[c1].x < in[c2].x {
            return true;
        } else {
            return in[c1].y < in[c2].y;
        }
    });
    j := 0
    for i := 1; i < len(in); i++ {
        if in[j] == in[i] {
            continue
        }
        j++
        // preserve the original data
        // in[i], in[j] = in[j], in[i]
        // only set what is required
        in[j] = in[i]
    }
    result := in[:j+1]
    // fmt.Println(result) // [1 2 3 4]
    return result;
}

func printMap(snek *snekType){
    print(snek.body);
    fmt.Printf("\n\n");
}
func print(in []coord) {
    myMap := make([][]string,30)
    for i := range myMap {
        myMap[i] = make([]string, 30)
    }

    for i := range myMap {
        for j := range myMap[i] {
            myMap[i][j] = "-"
        }
    }

    offset := 12;
    for _,c := range in {
        // myMap[c.y+offset][c.x+offset] = fmt.Sprintf("%d",i);
        myMap[c.y+offset][c.x+offset] = "#";
        // fmt.Printf("Dot: %d:%d\n",c.x+offset,c.y+offset)
    }
    myMap[offset][offset] = "s";
    for _,row := range makeRangeRev(0, len(myMap)-1) {
        fmt.Printf("%s\n",myMap[row])
    }

}
func print_str(in string) {
    out := strings.Split(in,",");
    // out := strings.Fields(in);
    fmt.Printf("%v\n",out)
    // print(out);
}

func main() {
    // file, _ := os.ReadFile("input");
    file, _ := os.ReadFile("test2");
    // file, _ := os.ReadFile("test3");
    bodySize := 10;
	// file, _ := os.ReadFile("test");
    // bodySize := 2;
    input := string(file);
    lines := strings.Split(input,"\n");
    // baseAmount := 2*size+2*(size-2);
    snek := snekType{};
    for i:=0;i<bodySize;i++ {
        snek.body = append(snek.body,coord{x:0, y:0});
    }

    for _,nextMove := range lines {
        if len(nextMove)==0 {
            break;
        }
        fmt.Printf("nextMove: %s\n",nextMove);
        parseMove(&snek,nextMove);
        // move := strings.Fields(nextMove);
    }
    set1 := mapset.NewSet[coord]();
    // for _,tailpos := range tailP {
    fmt.Printf("size: %v\n", set1.Cardinality());
    for _,tailpos := range snek.tailPos {
    //     // fmt.Printf("tail: %v\n",tailpos);
        set1.Add(tailpos);
    }
    // tailP := set1.String();

    //2720 too high unique,
    // 2592 tool high, set works on A.
    // 2285 too low
    // 2557 Correct!!
    tailPuniq := unique(snek.tailPos);
    // print(tailPuniq);
    fmt.Printf("ans1: %v, ans2: %v, tot: \n%v",len(tailPuniq) ,set1.Cardinality(),0);// ,len(tailP))
}
