package main

import "fmt"
import "testing"
import "github.com/cxcheng/sortgo/sortlib"

func runSorts(t *testing.T, vals []sortlib.SortVal, snapshots bool) {
    // loop thru each function
    for _, f := range(sortlib.SortFuncs) {
        // perform bubble sort
        ctx := sortlib.NewSortCtx()
        ctx.Sort(f, vals, true)
        // check that results are sorted
        for i := 0; i < len(vals) - 1; i++ {
            if !ctx.Lt(i, i + 1) && !ctx.Eq(i, i + 1) {
                fmt.Printf("Error: %s at %d\n", ctx.Title(), i)
                ctx.Print()
                t.Fail()
            }
        }
    }
}

func TestSortsWithRandom(t *testing.T) {
    // run through the sorting with array of random numbers
    runSorts(t, sortlib.RandomISortVals(20, 999), true)
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
        runSorts(t, sortlib.ISortVals(vals), true)
    }
}

func BenchmarkBubblesort(b *testing.B) {
    // generate array of 1000 random integers
    data := sortlib.RandomISortVals(5000, 999)

    // perform bubble sort
    ctx := sortlib.NewSortCtx()
    b.StartTimer()
    sortlib.Bubblesort(&ctx, data, false)
    b.StopTimer()
}

func BenchmarkQuicksort(b *testing.B) {
    // generate array of 1000 random integers
    data := sortlib.RandomISortVals(5000, 999)

    // perform bubble sort
    ctx := sortlib.NewSortCtx()
    b.StartTimer()
    sortlib.Quicksort(&ctx, data, false)
    b.StopTimer()
}
