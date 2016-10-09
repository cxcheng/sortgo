package sortlib

import "fmt"

type SortSnapshot struct {
    data SortData
    highlights map[int]int
}

type SortCtx struct {
    title string
    snapshots []SortSnapshot
    expectedOps int
}

func NewSortCtx() SortCtx {
    s := SortCtx {
        snapshots: make([]SortSnapshot, 0, 200),
    }
    return s
}

func (s SortCtx) addSnapshot(
    data SortData, highlights map[int]int) *SortSnapshot {
    snapshot := new(SortSnapshot)
    snapshot.data = data.copy()

    // copy the highlights
    snapshot.highlights = make(map[int]int)
    for k, v := range highlights {
        snapshot.highlights[k] = v
    }

    // add it to the snapshot list
    s.snapshots = append(s.snapshots, *snapshot)
    return snapshot
}

func (s SortCtx) Print() {
    numberCompares := 0
    numberSwaps := 0
    fmt.Println(s.title)
    for time, snapshot := range s.snapshots {
        fmt.Printf("%3d:", time)
        snapshot.data.print(snapshot.highlights)
        numberCompares += snapshot.data.getNumberCompares()
        numberSwaps += snapshot.data.getNumberSwaps()
    }
    fmt.Printf("Total: %d compares %d swaps, expected %d\n",
        numberCompares, numberSwaps, s.expectedOps)
}

