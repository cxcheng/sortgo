package sortlib

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

const TEXT_RESET = "\033[0m"
const TEXT_BOLD = "\033[1m"
const TEXT_RED = "\033[31m"
const TEXT_BLUE = "\033[34m"

type SortSnapshot struct {
    data []int
    redIndexes []int
    blueIndexes []int
    comment string
}

func (s *SortSnapshot) Print(time int, numberFmt string) {
    fmt.Printf("%3d:", time)
    for i := 0; i < len(s.data); i++ {
        fmt.Printf(" ")
        red, blue := false, false
        for _, s := range s.redIndexes {
           if s == i {
               red = true
               break
           }
        }
        for _, s := range s.blueIndexes {
           if s == i {
               blue = true
               break
           }
        }
        if red || blue {
            fmt.Printf(TEXT_BOLD) // bold
            // red has priority over blue
            if (red) {
                fmt.Printf(TEXT_RED)
            } else if (blue) {
                fmt.Printf(TEXT_BLUE)
            }
        }
        fmt.Printf(numberFmt, s.data[i])
        if red || blue {
            fmt.Printf(TEXT_RESET) // reset to default
        }
    }
    if len(s.comment) > 0 {
        fmt.Printf("   %s", s.comment)
    }
    fmt.Println()
}

type SortData struct {
    data []int
    snapshots []SortSnapshot
    numberFmt string
    numberCompares int
    numberSwaps int
    expectedOps int
}

func GenerateNumbers(num int, max int) []int {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    data := make([]int, num, num)
    for i := 0; i < num; i++ {
        data[i] = r1.Intn(max)
    }
    return data
}

func NewSortData(data []int) *SortData {
    s := &SortData {
        data: make([]int, len(data), len(data) * len(data)),
        snapshots: make([]SortSnapshot, 0, len(data) * len(data)),
        numberSwaps: 0, numberCompares: 0,
    }
    copy(s.data, data) // make a copy
    // add first snapshot for original
    s.addSnapshot( make([]int, 0, 0),  make([]int, 0, 0), "")
    // find the largest integer in the data to generate
    // the formatting string with number of digits
    largest := 0
    for _, n := range(data) {
        if n > largest {
            largest = n
        }
    }
    s.numberFmt = fmt.Sprintf("%%%dd", int(math.Log10(float64(largest))) + 1)
    return s
}

func (s *SortData) addSnapshot(
    redIndexes []int, blueIndexes []int, comment string) *SortSnapshot {
    snapshot := &SortSnapshot {
        data: make([]int, len(s.data), len(s.data)),
        redIndexes: redIndexes, blueIndexes: blueIndexes,
        comment: comment,
    }
    // make a copy of the data, but only keep references to indexes
    copy(snapshot.data, s.data)
    s.snapshots = append(s.snapshots, *snapshot)
    return snapshot
}

func (s *SortData) Print(name string) {
    fmt.Println(name)
    for time, snapshot := range s.snapshots {
        snapshot.Print(time, s.numberFmt)
    }
    fmt.Printf("Total: %d compares %d swaps, expected %d\n",
        s.numberCompares, s.numberSwaps, s.expectedOps)
}

func (s *SortData) Swap(i int, j int) {
   s.data[i], s.data[j] = s.data[j], s.data[i]
   s.numberSwaps++
}

