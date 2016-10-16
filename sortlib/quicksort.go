package sortlib

import "math"

func findPartition(ctx *SortCtx, vals []SortVal, lo int, hi int, snapshots bool) int {
    // choose pivot using Hoare partition scheme
    pivotIndex := lo
    i, j, partition_size := lo, hi - 1, hi - lo
    for  {
        for ctx.lt(i, pivotIndex) {
            // skip if already on the right side (< pivot)
            ctx.numberCompares++
            i++
        }
        for ctx.gt(j, pivotIndex) {
            // skip if already on the right side (> pivot)
            ctx.numberCompares++
            j--
        }
        if ctx.eq(i, j) {
           // we are done partitioning if i and j are the same
            break
        }  else if i < j {
            ctx.swap(i, j)
            if i == pivotIndex {
                pivotIndex = j
            } else if j == pivotIndex {
                pivotIndex = i
            }
            if snapshots {
                // mark the partition in bold
                highlights := make(map[int]int)
                for m := 0; m < partition_size; m++ {
                    highlights[lo + m] = 1
                }
                // overlay with the swapped values in red
                highlights[i], highlights[j] = 2, 2
                // overlay with pivot
                highlights[pivotIndex] = 3
                // submit snapshot
                ctx.addSnapshot(highlights)
            }
        } else {
            break // we are done when the two indexes touch
        }
    }
    return i
}

func quicksortPartition(ctx *SortCtx, vals []SortVal, lo int, hi int, snapshots bool) {
   if lo < hi {
       p := findPartition(ctx, vals, lo, hi, snapshots)
       quicksortPartition(ctx, vals, lo, p, snapshots)
       quicksortPartition(ctx, vals, p + 1, hi, snapshots)
   }
}

func Quicksort(ctx *SortCtx, vals []SortVal, snapshots bool) {
    ctx.title = "Quick Sort"
    n := len(vals)
    ctx.expectedOps = int(float64(n) * math.Log(float64(n))) // O(n lg n)
    quicksortPartition(ctx, vals, 0, n, snapshots)
}

