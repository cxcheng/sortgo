package sortlib

import "fmt"
import "math"
import "math/rand"
import "time"

type SortData interface {
    Copy() SortData
    Len() int
    Swap(i int, j int)
    Lt(i int, j int) bool
    Gt(i int, j int) bool
    Print(highlights map[int]int)
}

type ISortData struct {
    max int
    vals []int
}

func NewISortData(num int, max int) ISortData {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    vals := make([]int, num, num)
    for i := 0; i < num; i++ {
        vals[i] = r1.Intn(max)
    }
    return ISortData{ max: max, vals: vals }
}

func (s ISortData) Copy() SortData {
    // create a copy of the ISortData
    t := ISortData{}
    t.max = s.max
    t.vals = make([]int, len(s.vals), len(s.vals))
    copy(t.vals, s.vals)

    // return it as SortData interface
    t2 := new(SortData)
    *t2 = &t
    return *t2
}

func (s ISortData) Len() int {
    return len(s.vals)
}

func (s ISortData) Lt(i int, j int) bool {
    return s.vals[i] < s.vals[j]
}

func (s ISortData) Gt(i int, j int) bool {
    return s.vals[i] > s.vals[j]
}

func (s ISortData) Swap(i int, j int) {
    s.vals[i], s.vals[j] = s.vals[j], s.vals[i]
}

func (s ISortData) Print(highlights map[int]int) {
    const TEXT_RESET = "\033[0m"
    const TEXT_BOLD = "\033[1m"
    const TEXT_RED = "\033[31m"
    const TEXT_BLUE = "\033[34m"

    // the formatting string with number of digits
    fmtStr := fmt.Sprintf("%%%dd", int(math.Log10(float64(s.max))) + 1)

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
        fmt.Printf(fmtStr, vals[i])
        if isHighlighted {
            fmt.Printf(TEXT_RESET) // reset to default
        }
    }
    fmt.Println()
}
