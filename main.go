// This program runs through a few sort algorithms
// to demonstrate how different sorting algorithms works.
// This was written to get comfortable with Go, and to demonstrate to my son
// the differences between different sorting algorithms.

package main

import sortlib "github.com/cxcheng/sortgo/sortlib"

func main() {
    var data sortlib.SortData
    var ctx sortlib.SortCtx

    // generate array of random integers
    data = sortlib.NewISortData(10, 999)

    // perform bubble sort
    ctx = sortlib.NewSortCtx()
    sortlib.Bubblesort(&ctx, data)
    ctx.Print()

    // perform quicksort
    ctx = sortlib.NewSortCtx()
    sortlib.Quicksort(&ctx, data)
    ctx.Print()
}
