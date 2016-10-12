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
    data = sortlib.NewISortData(20, 999)

    // loop thru each function
    for _, f := range(sortlib.SortFuncs) {
        // perform sort and print results
        ctx = sortlib.NewSortCtx()
        ctx.Sort(f, data.Copy(), true)
        ctx.Print()
    }
}
