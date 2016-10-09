package sortlib

func Bubblesort(ctx SortCtx, data SortData) {
    n := data.len()
    ctx.title = "Bubble Sort"
    ctx.expectedOps = n * n // O(n**2)
    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        highlights := make(map[int]int)
         for i := 0; i < n - 1; i++ {
            if data.lt(i + 1, i) {
                swappedSomething = true
                data.swap(i, i + 1)
                highlights[i] = 1
                highlights[i + 1] = 1
                ctx.addSnapshot(data, highlights)
            }
        }
    }
}
