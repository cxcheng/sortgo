package sortlib

import "math/rand"
import "strconv"
import "time"

type IVal struct {
    Val int
}

func (s IVal) Eq(v Val) bool {
    if v2, ok := v.(IVal); ok {
        return s.Val == v2.Val
    } else {
        return false
    }
}

func (s IVal) Lt(v Val) bool {
    if v2, ok := v.(IVal); ok {
        return s.Val < v2.Val
    } else {
        return false
    }
}

func (s IVal) SnapshotString() string {
    return strconv.Itoa(s.Val)
}

func (s IVal) String() string {
    return strconv.Itoa(s.Val)
}

func RandomIVals(num int, max int) []Val {
    vals := make([]Val, num, num)
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    for i := 0; i < num; i++ {
        vals[i] = IVal{ Val: r1.Intn(max) }
    }
    return vals
}

func IVals(ints []int) []Val {
    vals := make([]Val, len(ints), len(ints))
    for i, v := range(ints) {
        vals[i] = IVal{ Val: v }
    }
    return vals
}
