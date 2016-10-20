package sortlib

import "math"

func quicksortPartition(ctx *Ctx, valsp *[]Val, lo int, hi int,
    count int) {
    size := hi - lo
    if size > 1 {
        pivot := (*valsp)[lo]
        left := lo
        right := hi - 1
        for left <= right {
            for ctx.Lt((*valsp)[left], pivot) {
                // skip if already on the right side (< pivot)
                left++
            }
            for ctx.Gt((*valsp)[right], pivot) {
                // skip if already on the right side (> pivot)
                right--
            }
            if left <= right {
                ctx.Swap(valsp, left, right)

                // mark the partition in bold
                highlights := make(map[int]int)
                for m := 0; m < size; m++ {
                    // check if pivot and mark it red
                    idx := lo + m
                    if (*valsp)[idx].Eq(pivot) {
                        highlights[idx] = 2
                    } else {
                        highlights[idx] = 3
                    }
                }
                // submit snapshot
                ctx.addSnapshot(*valsp, highlights)

                left++
                right--
           }
        }
        quicksortPartition(ctx, valsp, lo, right + 1, count + 1)
        quicksortPartition(ctx, valsp, left, hi, count + 1)
    }
}

func Quicksort(ctx *Ctx, vals []Val) {
    ctx.Title = "Quick Sort"
    n := len(vals)
    ctx.ExpectedOps = int(float64(n) * math.Log(float64(n))) // O(n lg n)
    quicksortPartition(ctx, &vals, 0, n, 0)
}

