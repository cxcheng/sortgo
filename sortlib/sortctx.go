package sortlib

import "fmt"

type SortSnapshot struct {
    data SortData
    highlights map[int]int
}

type SortCtx struct {
    title string
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

func (s *SortCtx) addSnapshot(data SortData, highlights map[int]int) *SortSnapshot {
    snapshot := new(SortSnapshot)
    snapshot.data = data.Copy()

    // copy the highlights
    snapshot.highlights = make(map[int]int)
    for k, v := range highlights {
        snapshot.highlights[k] = v
    }

    // add it to the snapshot list
    s.snapshots = append(s.snapshots, *snapshot)
    return snapshot
}

func (s *SortCtx) SortedData() *SortData {
    if len(s.snapshots) > 0 {
        return &s.snapshots[len(s.snapshots) - 1].data
    } else {
        return nil
    }
}

func (s *SortCtx) Print() {
    fmt.Println(s.title)
    for time, snapshot := range s.snapshots {
        fmt.Printf("%3d:", time)
        snapshot.data.Print(snapshot.highlights)
    }
    fmt.Printf("Total: %d compares %d swaps, expected %d\n",
        s.numberCompares, s.numberSwaps, s.expectedOps)
}

