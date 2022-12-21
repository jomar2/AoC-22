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
type folder struct{
    parent *folder;
    folders map[string]*folder;
    fileSizes []int;
    folderSize int;
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
func printFolder (inFolder folder, name string, level int) {
    for k,v := range inFolder.folders {
        printFolder(*v,k,level+1);
    }
    for i:=0;i<level;i++ {
        fmt.Printf("  ");
    }
    fmt.Printf("folder: %s , size: %d\n",name,inFolder.folderSize);
}
func calcFolder (inFolder *folder) int {
    fileTotal := 0;
    subFolderTotal := 0;
    for _,val := range inFolder.fileSizes {
        fileTotal += val;
    }
    for _,v := range inFolder.folders {
        subFolderTotal += calcFolder(v);
    }
    total := fileTotal + subFolderTotal;
    inFolder.folderSize = total;
    // fmt.Printf("file: %d , subfold: %d folderTotal: %d\n",fileTotal,subFolderTotal,total);
    return total;
}
func getFolders (inFolder *folder, limit int) int {
    total := 0;
    for _,v := range inFolder.folders {
        total += getFolders(v,limit);
    }
    if inFolder.folderSize < limit {
        total+=inFolder.folderSize;
    }
    return total;
}
func getBiggerFolders (inFolder *folder, limit int) []int {
    total := make([]int,0);

    for _,v := range inFolder.folders {
        total = append(total,getBiggerFolders(v,limit)...);
    }
    if inFolder.folderSize > limit {
        total = append(total,inFolder.folderSize);
    }
    return total;
}

func main() {
    file, _ := os.ReadFile("input");
	// file, _ := os.ReadFile("test");
    input := string(file);
    lines := strings.Split(input,"\n");
    empty := len(lines)==0;

    root := make(map[string]*folder)
    currentDir := "/";
    root[currentDir] = &folder{folders: make(map[string]*folder)}
    currentFolder := root[currentDir];
    var line string;
    for ;!empty;{
        line, lines, empty = Pop(lines);
        cmd := strings.Fields(line);
        fmt.Printf("line : %s\n",line);
        if cmd[0] == "$" {
            switch cmd[1] {
            case "cd":
                dir := cmd[2];
                if dir == "/" {
                    // fmt.Printf("gotoroot");
                    currentDir = "/";
                } else if dir == ".." {
                    // fmt.Printf("cd to parent")
                    currentFolder = currentFolder.parent;
                } else {
                    currentDir = dir;
                    // fmt.Printf("cd to dir %s\n",dir);
                    currentFolder = currentFolder.folders[dir];
                }
            case "ls":
                // fmt.Printf("ls!\n");
                for ;!strings.Contains(lines[0],"$"); {
                    var entry string;
                    entry,lines,empty = Pop(lines);
                    if empty {
                        break;
                    }
                    parts := strings.Fields(entry);
                    if parts[0] == "dir" {
                        dir := parts[1];
                        // fmt.Printf("add dir %s\n", dir);
                        if _,ok := currentFolder.folders[dir]; !ok {
                            // fmt.Printf("nok! %s\n", dir);
                            // currentFolder.folders[dir] = make(folder{parent: currentDir});
                            newFolder := folder{parent: currentFolder, folders: make(map[string]*folder)};
                            currentFolder.folders[dir] = &newFolder;
                        }
                    } else {
                        val := atoi(parts[0]);
                        // fmt.Printf("add file with size %d\n", val);
                        currentFolder.fileSizes = append(currentFolder.fileSizes,val);
                    }
                }
            }
        }
    }
    totalValue := calcFolder(root["/"]);
    spaceFree := 70000000-totalValue;
    spaceNeeded := 30000000-spaceFree;
    // printFolder(*root["/"],"/",0);
    ans1 := getFolders(root["/"],100000);
    tmp := getBiggerFolders(root["/"],spaceNeeded);
    sort.Ints(tmp);
    fmt.Printf("needed: %d, diff after: %d, len: %d,",spaceNeeded, tmp[0]-spaceNeeded, len(tmp))
    fmt.Printf("ans1: %v, ans2: %v, tot: %d", ans1,tmp[0] ,totalValue)
}
