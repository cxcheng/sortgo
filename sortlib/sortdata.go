package sortlib

import "fmt"
import "math"
import "math/rand"
import "time"

type SortData interface {
    copy() SortData
    len() int
    swap(i int, j int)
    lt(i int, j int) bool
    gt(i int, j int) bool
    getNumberCompares() int
    getNumberSwaps() int
    print(highlights map[int]int)
}

type SortDataBase struct {
    numberCompares int
    numberSwaps int
}

type ISortData struct {
    SortDataBase
    fmtStr string
    vals []int
}

func NewISortData(num int, max int) ISortData {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    vals := make([]int, num, num)
    for i := 0; i < num; i++ {
        vals[i] = r1.Intn(max)
    }
    // find the largest integer in the data to generate
    // the formatting string with number of digits
    fmtStr := fmt.Sprintf("%%%dd", int(math.Log10(float64(max))) + 1)
    return ISortData{ fmtStr: fmtStr, vals: vals }
}

func (s ISortData) copy() SortData {
    t := ISortData{
        vals: make([]int, 0, len(s.vals)),
    }
    copy(t.vals, s.vals)

    t2 := new(SortData)
    *t2 = &t
    return *t2
}

func (s ISortData) len() int {
    return len(s.vals)
}

func (s ISortData) swap(i int, j int) {
    s.vals[i], s.vals[j] = s.vals[j], s.vals[i]
    s.numberSwaps++
}

func (s ISortData) lt(i int, j int) bool {
    s.numberCompares++
    return s.vals[i] < s.vals[j]
}

func (s ISortData) gt(i int, j int) bool {
    s.numberCompares++
    return s.vals[i] > s.vals[j]
}

func (s ISortData) getNumberCompares() int {
    return s.numberCompares
}

func (s ISortData) getNumberSwaps() int {
    return s.numberSwaps
}

func (s ISortData) print(highlights map[int]int) {
    const TEXT_RESET = "\033[0m"
    const TEXT_BOLD = "\033[1m"
    const TEXT_RED = "\033[31m"
    const TEXT_BLUE = "\033[34m"

    vals := s.vals
    for i := 0; i < len(vals); i++ {
        fmt.Print(" ")
        mode, isHighlighted := highlights[i]
        if isHighlighted {
            switch mode {
            case 1:
                fmt.Print(TEXT_BOLD)
            case 2:
                fmt.Print(TEXT_RED)
            case 3:
                fmt.Print(TEXT_BLUE)
            }
        }
        fmt.Printf(s.fmtStr, vals[i])
        if isHighlighted {
            fmt.Printf(TEXT_RESET) // reset to default
        }
    }
    fmt.Println()
}
