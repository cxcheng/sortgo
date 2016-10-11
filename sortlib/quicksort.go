package sortlib

import "math"

func findPartition(ctx *SortCtx, data SortData, lo int, hi int, snapshots bool) int {
    // choose pivot using Hoare partition scheme
    pivotIndex := lo
    i, j, partition_size := lo, hi - 1, hi - lo
    for  {
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
        if data.Eq(i, j) {
            // we are done partitioning if i and j are the same
            break
        }  else if i < j {
            data.Swap(i, j)
            ctx.numberSwaps++
            if i == pivotIndex {
                pivotIndex = j
            } else if j == pivotIndex {
                pivotIndex = i
            }
            if snapshots {
                // mark the partition in blue
                highlights := make(map[int]int)
                for m := 0; m < partition_size; m++ {
                    highlights[lo + m] = 3
                }
                // overlay with the swapped values in red
                highlights[i], highlights[j] = 2, 2
                // overlay with pivot
                highlights[pivotIndex] = 1
                // submit snapshot
                ctx.addSnapshot(data, highlights)
            }
        } else {
            break // we are done when the two indexes touch
        }
    }
    return i
}

func quicksortPartition(ctx *SortCtx, data SortData, lo int, hi int, snapshots bool) {
   if lo < hi {
       p := findPartition(ctx, data, lo, hi, snapshots)
       quicksortPartition(ctx, data, lo, p, snapshots)
       quicksortPartition(ctx, data, p + 1, hi, snapshots)
   }
}

func Quicksort(ctx *SortCtx, data SortData, snapshots bool) {
    ctx.title = "Quick Sort"
    n := data.Len()
    ctx.expectedOps = int(float64(n) * math.Log(float64(n))) // O(n lg n)
    quicksortPartition(ctx, data, 0, n, snapshots)
}

