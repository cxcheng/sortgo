package sortlib

func Bubblesort(ctx *SortCtx, data SortData, snapshots bool) {
    ctx.title = "Bubble Sort"
    n := data.Len()
    ctx.expectedOps = n * n // O(n**2)

    var highlights map[int]int
    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        for i := 0; i < n - 1; i++ {
            if snapshots {
                // initialize for the new rounds
                highlights = make(map[int]int)
            }
            ctx.numberCompares++
            if data.Lt(i + 1, i) {
                swappedSomething = true
                data.Swap(i, i + 1)
                ctx.numberSwaps++
                if snapshots {
                    // increment the number of swaps
                    // highlight the swapped items in red
                    highlights[i], highlights[i + 1] = 2, 2
                    ctx.addSnapshot(data, highlights)
                }
            }
        }
    }
}
