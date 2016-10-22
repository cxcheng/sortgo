package main

import "fmt"
import "testing"
import "github.com/cxcheng/sortgo/sortlib"

func runSorts(t *testing.T, vals []sortlib.Val) {
    // loop thru each function
    for _, f := range(sortlib.SortFuncs) {
        // perform bubble sort
        ctx := sortlib.NewCtx()
        ctx.Sort(f, &vals)
        // check that results are sorted
        for i := 0; i < len(vals) - 1; i++ {
            if ctx.Lt(vals[i + 1], vals[i]) {
                fmt.Printf("Error: %s at %d, %s %s\n", ctx.Title, i,
                    vals[i], vals[i + 1])
                fmt.Println()
                ctx.Print()
                t.Fail()
            }
        }
    }
}

func TestSortsWithRandom(t *testing.T) {
    // run through the sorting with array of random numbers
    runSorts(t, sortlib.RandomIVals(20, 999))
}

func TestSortsWithFixed(t *testing.T) {
    valsArr := [...][]int {
        []int {},
        []int { 2 },
        []int { 10, 9, 8, 7, 7, 7, 7, 3, 2, 1 },
        []int {
           367, 367, 572, 752, 175, 813, 950, 734, 594, 402, 589, 668, 365, 889, 716, 814, 526, 434, 755, 810,
        },
        []int {
           367, 367, 572, 752, 175, 813, 950, 734, 594, 402, 589, 668, 365, 367, 889, 716, 814, 526, 434, 755, 810,
        },
    }
    for _, vals := range(valsArr) {
        // run through the sorting with array of pre-defined numbers that were known to cause issues
        runSorts(t, sortlib.IVals(vals))
    }
}

func TestSortsWithFixedStr(t *testing.T) {
    valsArr := [...][]string {
        []string {},
        []string { "abc" },
        []string { "abc", "def", "jkl" },
    }
    for _, vals := range(valsArr) {
        // run through the sorting with array of pre-defined numbers that were known to cause issues
        runSorts(t, sortlib.SVals(vals))
    }
}
func BenchmarkBubblesort(b *testing.B) {
    // generate array of 1000 random integers
    vals := sortlib.RandomIVals(5000, 999)

    // perform bubble sort
    ctx := sortlib.NewCtx()
    b.StartTimer()
    ctx.Sort(sortlib.Bubblesort, &vals)
    b.StopTimer()
}

func BenchmarkQuicksort(b *testing.B) {
    // generate array of 1000 random integers
    vals := sortlib.RandomIVals(5000, 999)

    // perform bubble sort
    ctx := sortlib.NewCtx()
    b.StartTimer()
    ctx.Sort(sortlib.Quicksort, &vals)
    b.StopTimer()
}
