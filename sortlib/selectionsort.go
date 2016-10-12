package sortlib

func Selectionsort(ctx *SortCtx, data SortData, snapshots bool) {
    ctx.title = "Selection Sort"
    n := data.Len()
    ctx.expectedOps = n * n // O(n**2)

    var i, j int
    /* advance the position through the entire array */
    /*   (could do j < n-1 because single element is also min element) */
    for j = 0; j < n - 1; j++ {
        /* find the min element in the unsorted a[j .. n-1] */

        // assume the min is the first element
        iMin := j
        // test against elements after j to find the smallest
        for i = j + 1; i < n; i++ {
            // if this element is less, then it is the new minimum
            ctx.numberCompares++
            if data.Lt(i, iMin) {
                // found new minimum; remember its index
                iMin = i
            }
        }
        if iMin != j {
            data.Swap(j, iMin)
            ctx.numberSwaps++
            if snapshots {
                // increment the number of swaps
                // highlight the swapped items in red
                ctx.addSnapshot(data, map[int]int{j:2, iMin:2,})
            }
        }
    }
}

