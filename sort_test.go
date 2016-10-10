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

func TestBubbleSort(t *testing.T) {
    var data sortlib.SortData
    var ctx sortlib.SortCtx

    // generate array of random integers
    data = sortlib.NewISortData(10, 999)

    // perform bubble sort
    ctx = sortlib.NewSortCtx()
    sortlib.Bubblesort(&ctx, data)

    // check results
    verifySorted(t, ctx.SortedData())
}
