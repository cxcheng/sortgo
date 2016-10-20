package sortlib

func Bubblesort(ctx *Ctx, vals []Val) {
    ctx.Title = "Bubble Sort"
    n := len(vals)
    ctx.ExpectedOps = n * n // O(n**2)

    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        for i := 0; i < n - 1; i++ {
            if ctx.Lt(vals[i + 1], vals[i]) {
                swappedSomething = true
                ctx.Swap(&vals, i, i + 1)
                // increment the number of swaps
                // highlight the swapped items in red
                ctx.addSnapshot(vals, map[int]int{i:2, i+1:2,})
            }
        }
    }
}
