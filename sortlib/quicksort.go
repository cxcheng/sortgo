package sortlib

import "math"

func findPartition(ctx *SortCtx, data SortData, lo int, hi int) int {
    // choose pivot using Hoare partition scheme
    pivotIndex := lo
    i, j, partition_size := lo, hi - 1, hi - lo
    highlights := make(map[int]int)
    for m := 0; m < partition_size; m++ {
        highlights[lo + m] = 3
    }
    for {
        for data.lt(i, pivotIndex) {
            // skip if already on the right side (< pivot)
            ctx.numberCompares++
            i++
        }
        for data.gt(j, pivotIndex) {
            // skip if already on the right side (> pivot)
            ctx.numberCompares++
            j--
        }
        if i < j {
            data.swap(i, j)
            ctx.numberSwaps++
            if i == pivotIndex {
                pivotIndex = j
            } else if j == pivotIndex {
                pivotIndex = i
            }
            // k := i
            // for k < j {
            //     if data[k] == pivot {
            //         break
            //     }
            //     k++
            // }
            // mark the swapped values red and add snapshot
            highlights[i], highlights[j] = 2, 2
            ctx.addSnapshot(data, highlights)
        } else {
            break // we are done when the two indexes touch
        }
    }
    ctx.addSnapshot(data, highlights)
    return i
}

func quicksortPartition(ctx *SortCtx, data SortData, lo int, hi int) {
   if lo < hi {
       p := findPartition(ctx, data, lo, hi)
       quicksortPartition(ctx, data, lo, p)
       quicksortPartition(ctx, data, p + 1, hi)
   }
}

func Quicksort(ctx *SortCtx, data SortData) {
    ctx.title = "Quick Sort"
    n := data.len()
    ctx.expectedOps = int(float64(n) * math.Log(float64(n))) // O(n lg n)
    quicksortPartition(ctx, data, 0, n)
}

