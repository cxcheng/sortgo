package sortlib

import "fmt"

// List of sort functions
var SortFuncs = []func(*Ctx, *[]Val) {
    Bubblesort, Quicksort, Selectionsort,
}

type Val interface {
    Eq(v Val) bool
    Lt(v Val) bool
    SnapshotString() string
    String() string
}

type Snapsort struct {
    Vals []Val
    highlights map[int]int
}

type Ctx struct {
    Title string
    FmtStr string
    Vals []Val
    Snapshots []Snapsort
    NumberCompares int
    NumberSwaps int
    ExpectedOps int
}

func NewCtx() Ctx {
    s := Ctx {
        Snapshots: make([]Snapsort, 0, 200),
    }
    return s
}

func (s *Ctx) LastSnapshot() Snapsort {
    return s.Snapshots[len(s.Snapshots) - 1]
}

func (s *Ctx) Sort(f func(*Ctx, *[]Val), valsp *[]Val) {
    // add initial snapshot
    s.addSnapshot(*valsp, nil)
    // now sort
    f(s, valsp)
}

func (s *Ctx) Print() {
    fmt.Println(s.Title)
    if s.FmtStr == "" {
        s.FmtStr = "%3s"
    }
    for time, snapshot := range s.Snapshots {
        fmt.Printf("%3d:", time)
        s.PrintSnapshot(snapshot)
    }
    fmt.Printf("Total: %d compares %d swaps, expected %d\n",
        s.NumberCompares, s.NumberSwaps, s.ExpectedOps)
}

func (s *Ctx) addSnapshot(vals []Val, highlights map[int]int) {
    // make a copy of the vals
    n := len(vals)
    vals2 := make([]Val, n, n)
    copy(vals2, vals)
    snapshot := Snapsort{ Vals: vals2, highlights: highlights }

    // copy the highlights
    snapshot.highlights = make(map[int]int)
    for k, v := range highlights {
        snapshot.highlights[k] = v
    }

    // add it to the snapshot list
    s.Snapshots = append(s.Snapshots, snapshot)
}

func (s *Ctx) PrintSnapshot(snapshot Snapsort) {
    const TextReset = "\033[0m"
    const TextBold = "\033[1m"
    const TextRed = "\033[31m"
    const TextBlue = "\033[34m"

    vals := snapshot.Vals
    for i, val := range vals {
        fmt.Print(" ")
        mode, isHighlighted := snapshot.highlights[i]
        if isHighlighted {
            switch mode {
            case 1:
                fmt.Print(TextBold)
            case 2:
                fmt.Print(TextRed)
            case 3:
                fmt.Print(TextBlue)
            }
        }
        if valStr := val.SnapshotString(); valStr == "" {
            fmt.Printf("%2d", i)
        } else {
            fmt.Printf(s.FmtStr, valStr)
        }
        if isHighlighted {
            fmt.Printf(TextReset) // reset to default
        }
    }
    fmt.Println()
}

func (s *Ctx) Gt(v1 Val, v2 Val) bool {
    s.NumberCompares++
    return !v1.Eq(v2) && !v1.Lt(v2)
}

func (s *Ctx) Lt(v1 Val, v2 Val) bool {
    s.NumberCompares++
    return v1.Lt(v2)
}

func (s *Ctx) Swap(valsp *[]Val, i int, j int) {
    s.NumberSwaps++
    (*valsp)[i], (*valsp)[j] = (*valsp)[j], (*valsp)[i]
}
