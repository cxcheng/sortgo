// This program runs through a few sort algorithms
// to demonstrate how different sorting algorithms works.
// This was written to get comfortable with Go, and to demonstrate to my son
// the differences between different sorting algorithms.

package main

import sortlib "github.com/cxcheng/sortgo/sortlib"

func main() {
   data := sortlib.GenerateNumbers(10, 999)

   sortData := sortlib.NewSortData(data)
   sortlib.BubbleSort(sortData)
   sortData.Print("Bubble Sort")

   sortData = sortlib.NewSortData(data)
   sortlib.QuickSort(sortData)
   sortData.Print("Quick Sort- pivot in red")
}
