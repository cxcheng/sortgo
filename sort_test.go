package main

import "testing"
import sortlib "github.com/cxcheng/sortgo/sortlib"

func verifySorted(t *testing.T, sorted *sortlib.SortData) {
    // check that results are sorted
    sortedVal := *sorted
    for i := 0; i < sortedVal.Len() - 1; i++ {
        if !sortedVal.Lt(i, i + 1) {
            t.Fail()
        }
    }
}

func TestBubblesort(t *testing.T) {
    var data sortlib.SortData
    var ctx sortlib.SortCtx

    // generate array of random integers
    data = sortlib.NewISortData(10, 999)

    // perform bubble sort
    ctx = sortlib.NewSortCtx()
    sortlib.Bubblesort(&ctx, data, true)

    // check results
    verifySorted(t, ctx.SortedData())
}

func TestQuicksort(t *testing.T) {
    var data sortlib.SortData
    var ctx sortlib.SortCtx

    // generate array of random integers
    data = sortlib.NewISortData(10, 999)

    // perform bubble sort
    ctx = sortlib.NewSortCtx()
    sortlib.Quicksort(&ctx, data, true)

    // check results
    verifySorted(t, ctx.SortedData())
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
