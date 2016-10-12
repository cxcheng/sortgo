package sortlib

func Bubblesort(ctx *SortCtx, data SortData, snapshots bool) {
    ctx.title = "Bubble Sort"
    n := data.Len()
    ctx.expectedOps = n * n // O(n**2)

    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        for i := 0; i < n - 1; i++ {
            ctx.numberCompares++
            if data.Lt(i + 1, i) {
                swappedSomething = true
                data.Swap(i, i + 1)
                ctx.numberSwaps++
                if snapshots {
                    // increment the number of swaps
                    // highlight the swapped items in red
                    ctx.addSnapshot(data, map[int]int{i:2, i+1:2,})
                }
            }
        }
    }
}
