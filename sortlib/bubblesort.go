package sortlib

func Bubblesort(ctx *SortCtx, data SortData) {
    dataCopy := data.copy()
    n := dataCopy.Len()

    ctx.title = "Bubble Sort"
    ctx.expectedOps = n * n // O(n**2)
    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        highlights := make(map[int]int)
        for i := 0; i < n - 1; i++ {
            ctx.numberCompares++
            if dataCopy.Lt(i + 1, i) {
                swappedSomething = true
                dataCopy.swap(i, i + 1)
                ctx.numberSwaps++
                // increment the number of swaps
                // highlight the swapped items in red
                highlights[i], highlights[i + 1] = 2, 2
                ctx.addSnapshot(dataCopy, highlights)
            }
        }
    }
}
