package sortlib

func Bubblesort(ctx *SortCtx, data SortData, snapShots bool) {
    ctx.title = "Bubble Sort"
    n := data.Len()
    ctx.expectedOps = n * n // O(n**2)

    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        for i := 0; i < n - 1; i++ {
            var highlights map[int]int
            if snapShots {
                highlights = make(map[int]int)
            }
            ctx.numberCompares++
            if data.Lt(i + 1, i) {
                swappedSomething = true
                data.Swap(i, i + 1)
                ctx.numberSwaps++
                if snapShots {
                    // increment the number of swaps
                    // highlight the swapped items in red
                    highlights[i], highlights[i + 1] = 2, 2
                    ctx.addSnapshot(data, highlights)
                }
            }
        }
    }
}
