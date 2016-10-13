package main

import "fmt"
import "testing"
import "github.com/cxcheng/sortgo/sortlib"

func runSorts(t *testing.T, data sortlib.SortData, snapshots bool) {
    // loop thru each function
    for _, f := range(sortlib.SortFuncs) {
        // perform bubble sort
        ctx := sortlib.NewSortCtx()
        ctx.Sort(f, data.Copy(), true)
        // check that results are sorted
        sortedVal := *ctx.SortedData()
        for i := 0; i < sortedVal.Len() - 1; i++ {
            if !sortedVal.Lt(i, i + 1) && !sortedVal.Eq(i, i + 1) {
                fmt.Printf("Error: %s\n", ctx.Title())
                sortedVal.Print(map[int]int{i:2, i+1:2,})
                ctx.Print()
                t.Fail()
            }
        }
    }
}

func TestSortsWithRandom(t *testing.T) {
    // run through the sorting with array of random numbers
    runSorts(t, sortlib.NewISortData(20, 999), true)
}

func TestSortsWithFixed(t *testing.T) {
    valsArr := [...][]int {
        [] int {},
        [] int { 2 },
        []int {
           367, 367, 572, 752, 175, 813, 950, 734, 594, 402, 589, 668, 365, 889, 716, 814, 526, 434, 755, 810,
        },
        []int {
           367, 367, 572, 752, 175, 813, 950, 734, 594, 402, 589, 668, 365, 367, 889, 716, 814, 526, 434, 755, 810,
        },
    }
    for _, vals := range(valsArr) {
        // run through the sorting with array of pre-defined numbers that were known to cause issues
        runSorts(t, sortlib.NewISortDataVals(vals), true)
    }
}

func BenchmarkBubblesort(b *testing.B) {
    var data sortlib.SortData
    var ctx sortlib.SortCtx

    // generate array of 1000 random integers
    data = sortlib.NewISortData(5000, 999)

    // perform bubble sort
    ctx = sortlib.NewSortCtx()
    b.StartTimer()
    sortlib.Bubblesort(&ctx, data, false)
    b.StopTimer()
}

func BenchmarkQuicksort(b *testing.B) {
    var data sortlib.SortData
    var ctx sortlib.SortCtx

    // generate array of 1000 random integers
    data = sortlib.NewISortData(5000, 999)

    // perform bubble sort
    ctx = sortlib.NewSortCtx()
    b.StartTimer()
    sortlib.Quicksort(&ctx, data, false)
    b.StopTimer()
}
