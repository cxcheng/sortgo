package sortlib

import "math"

func findPartition(ctx *SortCtx, data SortData,
    lo int, hi int, snapShots bool) int {
    // choose pivot using Hoare partition scheme
    pivotIndex := lo
    i, j, partition_size := lo, hi - 1, hi - lo

    for {
        for data.Lt(i, pivotIndex) {
            // skip if already on the right side (< pivot)
            ctx.numberCompares++
            i++
        }
        for data.Gt(j, pivotIndex) {
            // skip if already on the right side (> pivot)
            ctx.numberCompares++
            j--
        }
        if i < j {
            data.Swap(i, j)
            ctx.numberSwaps++
            if i == pivotIndex {
                pivotIndex = j
            } else if j == pivotIndex {
                pivotIndex = i
            }
            if snapShots {
                highlights := make(map[int]int)
                for m := 0; m < partition_size; m++ {
                    highlights[lo + m] = 3
                }
                // mark the swapped values red and add snapshot
                highlights[i], highlights[j] = 2, 2
                ctx.addSnapshot(data, highlights)
            }
        } else {
            break // we are done when the two indexes touch
        }
    }
    ctx.addSnapshot(data, make(map[int]int))
    return i
}

func quicksortPartition(ctx *SortCtx, data SortData,
    lo int, hi int, snapShots bool) {
   if lo < hi {
       p := findPartition(ctx, data, lo, hi, snapShots)
       quicksortPartition(ctx, data, lo, p, snapShots)
       quicksortPartition(ctx, data, p + 1, hi, snapShots)
   }
}

func Quicksort(ctx *SortCtx, data SortData, snapShots bool) {
    ctx.title = "Quick Sort"
    n := data.Len()
    ctx.expectedOps = int(float64(n) * math.Log(float64(n))) // O(n lg n)
    quicksortPartition(ctx, data, 0, n, snapShots)
}

