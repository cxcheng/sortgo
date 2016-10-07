package sortlib

import (
    "fmt"
    "math"
)

func BubbleSort(s *SortData) {
    n := len(s.data)
    s.expectedOps = n * n // O(n**2)
    for swappedSomething := true; swappedSomething; {
        swappedSomething = false
        redIndexes := make([]int, 2, n)
        for i := 0; i < len(s.data) - 1; i++ {
            s.numberCompares++
            if s.data[i + 1] < s.data[i] {
                swappedSomething = true
                 s.Swap(i, i + 1)
                redIndexes = append([] int{ i, i + 1})
                s.addSnapshot(redIndexes, nil, "")
            }
        }
    }
}

func partition(s *SortData, lo int, hi int) int {
     // choose pivot uses Hoare partition scheme
    data := s.data
    pivot := data[lo]
    i, j, partition_size := lo, hi - 1, hi - lo
    comment := fmt.Sprintf("pivot %d", pivot)
    redIndexes := []int{ lo }
    blueIndexes := make([]int, partition_size, partition_size)
    for m := 0; m < partition_size; m++ {
        blueIndexes[m] = lo + m
    }
    for {
        s.numberCompares++
        for data[i] < pivot {
            // skip if already on the right side (< pivot)
            i++
        }
        for data[j] > pivot {
            // skip if already on the right side (> pivot)
            j--
        }
        if i < j {
            s.Swap(i, j)
            if data[i] == pivot {
                redIndexes[0] = i
            } else if data[j] == pivot {
                redIndexes[0] = j
            }
            k := i
            for k < j {
                if data[k] == pivot {
                    break
                }
                k++
            }
            s.addSnapshot(redIndexes, blueIndexes, comment)
        } else {
            break // we are done when the two indexes touch
        }
    }
    s.addSnapshot(redIndexes, blueIndexes, comment)
    return i
}

func quickSort2(s *SortData, lo int, hi int) {
   if lo < hi {
       p := partition(s, lo, hi)
       quickSort2(s, lo, p)
       quickSort2(s, p + 1, hi)
   }
}

func QuickSort(s *SortData) {
    n := len(s.data)
    s.expectedOps = int(float64(n) * math.Log(float64(n))) // O(n lg n)
    quickSort2(s, 0, n)
}

