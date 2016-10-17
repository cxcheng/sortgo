package sortlib

import "fmt"

// List of sort functions
var SortFuncs = []func(*SortCtx, []SortVal, bool) {
    Bubblesort, Quicksort, Selectionsort,
}

type SortVal interface {
    Eq(v SortVal) bool
    Gt(v SortVal) bool
    Lt(v SortVal) bool
    SnapshotString() string
    ValueString() string
}

type SortSnapshot struct {
    vals []SortVal
    highlights map[int]int
}

type SortCtx struct {
    title string
    fmtStr string
    snapshots []SortSnapshot
    numberCompares int
    numberSwaps int
    expectedOps int
}

func NewSortCtx() SortCtx {
    s := SortCtx {
        snapshots: make([]SortSnapshot, 0, 200),
    }
    return s
}

func (s *SortCtx) LastSnapshot() *SortSnapshot {
    if len(s.snapshots) > 0 {
        return &s.snapshots[len(s.snapshots) - 1]
    } else {
        return nil
    }
}

func (s *SortCtx) Sort(f func(*SortCtx, []SortVal, bool), vals []SortVal, snapshots bool) {
    // add initial snapshot
    snapshot := SortSnapshot{ vals: vals, highlights: nil }
    s.snapshots = append(s.snapshots, snapshot)
    // now sort
    f(s, vals, snapshots)
}

func (s *SortCtx) Title() string {
    return s.title
}

func (s *SortCtx) Print() {
    fmt.Println(s.title)
    if s.fmtStr == "" {
        s.fmtStr = "%3s"
    }
    for time, snapshot := range s.snapshots {
        fmt.Printf("%3d:", time)
        s.PrintSnapshot(snapshot)
    }
    fmt.Printf("Total: %d compares %d swaps, expected %d\n",
        s.numberCompares, s.numberSwaps, s.expectedOps)
}

func (s *SortCtx) addSnapshot(highlights map[int]int) *SortSnapshot {
    last := s.LastSnapshot()
    if last != nil {
        len2 := len((*last).vals)
        vals2 := make([]SortVal, len2, len2)
        copy(vals2, (*last).vals)
        snapshot := SortSnapshot{ vals: vals2, highlights: highlights }

        // copy the highlights
        snapshot.highlights = make(map[int]int)
        for k, v := range highlights {
            snapshot.highlights[k] = v
        }

        // add it to the snapshot list
        s.snapshots = append(s.snapshots, snapshot)
        return &snapshot
    } else {
        return nil
    }
}

func (s *SortCtx) PrintSnapshot(snapshot SortSnapshot) {
    const TEXT_RESET = "\033[0m"
    const TEXT_BOLD = "\033[1m"
    const TEXT_RED = "\033[31m"
    const TEXT_BLUE = "\033[34m"

    vals := snapshot.vals
    for i, val := range vals {
        fmt.Print(" ")
        mode, isHighlighted := snapshot.highlights[i]
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
        if valStr := val.SnapshotString(); valStr == "" {
            fmt.Printf("%2d", i)
        } else {
            fmt.Printf(s.fmtStr, valStr)
        }
        if isHighlighted {
            fmt.Printf(TEXT_RESET) // reset to default
        }
    }
    fmt.Println()
}

func (s *SortCtx) Eq(i int, j int) bool {
    snapshot := s.LastSnapshot()
    if snapshot != nil {
        s.numberCompares++
        vals := snapshot.vals
        return vals[i].Eq(vals[j])
    } else {
        return false
    }
}

func (s *SortCtx) Lt(i int, j int) bool {
    snapshot := s.LastSnapshot()
    if snapshot != nil {
        s.numberCompares++
        vals := snapshot.vals
        return vals[i].Lt(vals[j])
    } else {
        return false
    }
}

func (s *SortCtx) Gt(i int, j int) bool {
    snapshot := s.LastSnapshot()
    if snapshot != nil {
        s.numberCompares++
        vals := snapshot.vals
        return vals[i].Gt(vals[j])
    } else {
        return false
    }
}

func (s *SortCtx) Swap(i int, j int) {
    snapshot := s.LastSnapshot()
    if snapshot != nil {
        s.numberSwaps++
        snapshot.vals[i], snapshot.vals[j] = snapshot.vals[j], snapshot.vals[i]
    }
}
