package sortlib

type SSortVal struct {
    Val string
}

func (s SSortVal) Eq(v SortVal) bool {
    if v2, ok := v.(SSortVal); ok {
        return s.Val == v2.Val
    } else {
        return false
    }
}

func (s SSortVal) Gt(v SortVal) bool {
    if v2, ok := v.(SSortVal); ok {
        return s.Val > v2.Val
    } else {
        return false
    }
}

func (s SSortVal) Lt(v SortVal) bool {
    if v2, ok := v.(SSortVal); ok {
        return s.Val < v2.Val
    } else {
        return false
    }
}

func (s SSortVal) Stringify() string {
    return s.Val
}

