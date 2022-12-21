package main

import (
	"fmt"
	"math"
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

func makeRange(min, max int64) []int64 {
   // fmt.Printf("min,max: %v,%v\n", min,max);
    a := make([]int64, max-min+1)
    for i := range a {
        a[i] = int64(int(min) + i);
    }
   // fmt.Printf("range: %v\n", a);
    return a
}
func makeRangeSym(start, length int64) []int64 {
   // fmt.Printf("min,max: %v,%v\n", min,max);
    start_val := start;
    a := []int64{};
    // fmt.Printf("R: start %d, l: %d\n",start,length);

    for i := length*-1; i<length+1;i++ {
        // fmt.Printf("ehh..: %v\n",start_val+i);
        a = append(a,start_val+i);
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
func makeRangeStr(min, max string) []int64 {
    max_i := atoi(max);
    min_i := atoi(min);

    return makeRange(min_i,max_i);

}
func atoi(input string) int64 {
    val,_ := strconv.Atoi(strings.Trim(input," "));
    return int64(val);
}
func atoiStrVec(input []string) [][]int {
    var output [][]int;
    for _,inp := range input {
        var tmp []int;
        for _,val := range strings.Split(inp,"") {
            tmp = append(tmp,int(atoi(val)));
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

    // for _, crates := range input[1:] {
    //     crates = re2.ReplaceAllString(crates, "[x]")
    //     crates = re.ReplaceAllString(crates, "")
    //     }
    // }
type coord struct {
    x int64;
    y int64;
}
type sensor struct {
    pos coord;
    beacon coord;
    distance int;
}

func main() {
    re, _ := regexp.Compile(`-?\d+`)
    // file, _ := os.ReadFile("input");
    file, _ := os.ReadFile("test");
    // file, _ := os.ReadFile("test2");
    // row :=2000000;
    // row := 10;

    input := string(file);
    lines := strings.Split(input,"\n");
    sensors := []sensor{}
    beacons := map[coord]string{};
    sensorsCoord := map[coord]string{};
    max := coord{};
    min := coord{};
    for i,dLines := range lines {
        parts := re.FindAllString(dLines,5);
        if len(parts)<4 {
            break;
        }
        x := atoi(parts[0]);
        x_b := atoi(parts[2]);
        y := atoi(parts[1]);
        y_b := atoi(parts[3]);
        distance := int(math.Abs(float64(x_b-x))+math.Abs(float64(y_b-y)));
        // updateLimits(x,y,x_b,y_b,&max,&min);
        updateLimits(x,y,int64(distance),&max,&min);
    
        sensors = append(sensors, sensor{
            pos: coord{x: x, y: y},
            beacon: coord{x: x_b, y: y_b},
            distance: distance});
            beacons[coord {x: x_b,y: y_b}] = "B";
            sensorsCoord[coord {x: x,y: y}] = "S";
        fmt.Printf("Done index %d, %v, len %d\n",i,parts,len(parts));
    }
    range_x := max.x-min.x;
    range_y := max.y-min.y;
    fmt.Printf("ranges: x: %d,y: %d\n",range_x,range_y);

    // emptyArea := make([][]byte,range_y);
    // for i := range emptyArea {
    //     emptyArea[i] = make([]byte,range_x);
    // }
    possible := map[coord]string{};

    for i := 0; i<20; i++ {
        possibleCoord(sensors,sensorsCoord,beacons, max,min,i,&possible);
    }
    // possiblebeacons(
    // printMap(sensors,beacons,emptyRow,max,min);
    // printRow(sensors,beacons,emptyRow,max,min,row);
    sum := 0;
    // var testval int64 = 10;
    // Part A:
    // emptyRow := noBeaconArea(sensors,sensorsCoord,beacons, max,min,row);
    // for _,v := range emptyRow {
    //         if v == "#" {
    //             // fmt.Printf("empty coord: %v\n",i);
    //             sum++;
    //         }
    // }

    // Too low: 4850556
    // Too Low: 4850457
    // Too Low: 4425457
    // Correct A: 5688618
    fmt.Printf("ans1: %v, ans2: %v, tot: \n%v",sum, len(possible),possible)
}

func noBeaconArea(sensors []sensor,beacons,sensorsCoord map[coord]string,max,min coord,row int) map[coord]string{
    empty := map[coord]string{};
    start_x := min.x;
    y := int64(row);
    for i := start_x; i< max.x; i++ {
        for _,s := range sensors {
            distance := int(math.Abs(float64(i-s.pos.x))+math.Abs(float64(y-s.pos.y)));
            p := coord{x:i,y:y};
            _,beacon := beacons[p];
            _,sensor := sensorsCoord[p];
            if distance<=s.distance && !beacon && !sensor {
                empty[p] = "#";
                // fmt.Printf("Ok on %v\n,distance: %d, beacon distance: %d\n",p,distance,s.distance);
                break;
            }
        }
    }
    return empty;
}

func possibleCoord(sensors []sensor,beacons,sensorsCoord map[coord]string,max,min coord,row int, output *map[coord]string) {
    start_x := int64(0);
    y := int64(row);
    for i := start_x; i< 20; i++ {
        p := coord{x:i,y:y};
        for _,s := range sensors {
            distance := int(math.Abs(float64(i-s.pos.x))+math.Abs(float64(y-s.pos.y)));
            _,beacon := beacons[p];
            _,sensor := sensorsCoord[p];
            if distance<=s.distance && !beacon && !sensor {
                // fmt.Printf("Ok on %v\n,distance: %d, beacon distance: %d\n",p,distance,s.distance);
                continue;
            }
        }
        (*output)[coord{x:i,y:y}] = "P";
    }
}


func updateLimits(x, y, distance int64, max, min *coord) {
    if x+distance>max.x {
        max.x = x+distance;
    }
    if y+distance>max.y {
        max.y = y+distance;
    }
    if x-distance<min.x {
        min.x = x-distance;
    }
    if y-distance<min.y {
        min.y = y-distance;
    }
}

func printMap(sensors []sensor, beacons map[coord]string,emptyArea map[coord]string,max,min coord) {
    var offset_x int64;
    var offset_y int64;
    if min.x < 0 {
        offset_x = min.x*-1;
    }
    if min.y < 0 {
        offset_y = min.y*-1;
    }

    myMap := make([][]string,offset_y+max.y+1);
    for i := range myMap {
        myMap[i] = make([]string,offset_x+max.x+1);
        for c := range myMap[i] {
            myMap[i][c] = ".";
        }
    }
    fmt.Printf("Mah Map : %dx%d\noffset: x:%d,y:%d\n",offset_x+max.x+1,offset_y+max.y+1,offset_x,offset_y);
    for s := range emptyArea {

        if s.x < min.x || s.y < min.y || s.x >= max.x || s.y >= max.y {
            continue;
        }
        // fmt.Printf("crasing at : %v\n", s);
        myMap[offset_y+s.y][offset_x+s.x] = "#";
    }
    for _,s := range sensors {
        myMap[offset_y+s.pos.y][offset_x+s.pos.x] = "S";
    }
    for s := range beacons {
        myMap[offset_y+s.y][offset_x+s.x] = "B";
    }
    fmt.Printf("\nmin: x:%d,y:%d  -- max: x:%d, y:%d\n",min.x,min.y,max.x,max.y);
    for _,row := range myMap {
        fmt.Printf("%s\n",row);
    }
    // fmt.Printf("%v\n",myMap);

}


func printRow(sensors []sensor, beacons map[coord]string,emptyArea map[coord]string,max,min coord,row int) {
    var offset_x int64;
    var offset_y int64;
    if min.x < 0 {
        offset_x = min.x*-1;
    }

    myMap := make([]string,offset_x+max.x+1);
    for c := range myMap {
        myMap[c] = ".";
    }
    fmt.Printf("Mah Map : %dx%d\noffset: x:%d,y:%d\n",offset_x+max.x+1,offset_y+max.y+1,offset_x,offset_y);
    for s := range emptyArea {
        if s.x < min.x || s.x >= max.x {
            continue;
        }
        // fmt.Printf("crasing at : %v\n", s);
        myMap[offset_x+s.x] = "#";
    }
    for _,s := range sensors {
        myMap[offset_x+s.pos.x] = "S";
    }
    for s := range beacons {
        myMap[offset_x+s.x] = "B";
    }
    fmt.Printf("\nmin: x:%d,y:%d  -- max: x:%d, y:%d\n",min.x,min.y,max.x,max.y);
    for _,row := range myMap {
        fmt.Printf("%s",row);
    }
    // fmt.Printf("%v\n",myMap);

}


