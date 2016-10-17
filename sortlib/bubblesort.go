package sortlib

func Bubblesort(ctx *SortCtx, vals []SortVal, snapshots bool) {
    ctx.title = "Bubble Sort"
    n := len(vals)
    ctx.expectedOps = n * n // O(n**2)

    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        for i := 0; i < n - 1; i++ {
            if ctx.Lt(i + 1, i) {
                swappedSomething = true
                ctx.Swap(i, i + 1)
                if snapshots {
                    // increment the number of swaps
                    // highlight the swapped items in red
                    ctx.addSnapshot(map[int]int{i:2, i+1:2,})
                }
            }
        }
    }
}
