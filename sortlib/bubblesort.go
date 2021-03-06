package sortlib

func Bubblesort(ctx *Ctx, valsp *[]Val) {
    ctx.Title = "Bubble Sort"
    n := len(*valsp)
    ctx.ExpectedOps = n * n // O(n**2)

    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        for i := 0; i < n - 1; i++ {
            if ctx.Lt((*valsp)[i + 1], (*valsp)[i]) {
                swappedSomething = true
                ctx.Swap(valsp, i, i + 1)
                // increment the number of swaps
                // highlight the swapped items in red
                ctx.addSnapshot(*valsp, map[int]int{i:2, i+1:2,})
            }
        }
    }
}
