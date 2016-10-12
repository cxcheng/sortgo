package main

import "fmt"
import "testing"
import "github.com/cxcheng/sortgo/sortlib"

func verifySorted(t *testing.T, ctx sortlib.SortCtx) {
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

func TestSorts(t *testing.T) {
    var data sortlib.SortData
    var ctx sortlib.SortCtx

    // generate array of random integers
    data = sortlib.NewISortData(20, 999)

    // loop thru each function
    for _, f := range(sortlib.SortFuncs) {
        // perform bubble sort
        ctx = sortlib.NewSortCtx()
        ctx.Sort(f, data.Copy(), true)
        // check results
        verifySorted(t, ctx)
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
