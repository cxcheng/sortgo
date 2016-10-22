// This program runs through a few sort algorithms
// to demonstrate how different sorting algorithms works.
// This was written to get comfortable with Go, and to demonstrate to my son
// the differences between different sorting algorithms.

package main

import "fmt"
import "github.com/cxcheng/sortgo/sortlib"

func printVals(label string, vals []sortlib.Val) {
    fmt.Print(label, ":")
    for _, val := range(vals) {
        fmt.Print(" ", val.String())
    }
    fmt.Println()
}

func main() {
    // generate array of random integers and loop through each sort function
    vals := sortlib.RandomIVals(20, 999)
    for _, f := range(sortlib.SortFuncs) {
        // perform sort and print results
        ctx := sortlib.NewCtx()
        vals2 := make([]sortlib.Val, len(vals), len(vals))
        copy(vals2, vals)
        ctx.Sort(f, &vals2)
        ctx.Print()
    }

    // test the strings
    svals := sortlib.SValsFromFile("text/planets.txt")
    for _, f := range(sortlib.SortFuncs) {
        // perform sort and print results
        ctx := sortlib.NewCtx()
        svals2 := make([]sortlib.Val, len(svals), len(svals))
        copy(svals2, svals)
        printVals("Before", svals2)
        ctx.Sort(f, &svals2)
        ctx.Print()
        printVals("After", svals2)
    }
}
