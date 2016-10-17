package sortlib

import "math/rand"
import "strconv"
import "time"

type ISortVal struct {
    Val int
}

func (s ISortVal) Eq(v SortVal) bool {
    if v2, ok := v.(ISortVal); ok {
        return s.Val == v2.Val
    } else {
        return false
    }
}

func (s ISortVal) Gt(v SortVal) bool {
    if v2, ok := v.(ISortVal); ok {
        return s.Val > v2.Val
    } else {
        return false
    }
}

func (s ISortVal) Lt(v SortVal) bool {
    if v2, ok := v.(ISortVal); ok {
        return s.Val < v2.Val
    } else {
        return false
    }
}

func (s ISortVal) SnapshotString() string {
    return strconv.Itoa(s.Val)
}

func (s ISortVal) ValueString() string {
    return strconv.Itoa(s.Val)
}

func RandomISortVals(num int, max int) []SortVal {
    vals := make([]SortVal, num, num)
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    for i := 0; i < num; i++ {
        vals[i] = ISortVal{ Val: r1.Intn(max) }
    }
    return vals
}

func ISortVals(ints []int) []SortVal {
    vals := make([]SortVal, len(ints), len(ints))
    for i, v := range(ints) {
        vals[i] = ISortVal{ Val: v }
    }
    return vals
}
