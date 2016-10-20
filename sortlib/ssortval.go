package sortlib

type SVal struct {
    Val string
}

func (s SVal) Eq(v Val) bool {
    if v2, ok := v.(SVal); ok {
        return s.Val == v2.Val
    } else {
        return false
    }
}

func (s SVal) Lt(v Val) bool {
    if v2, ok := v.(SVal); ok {
        return s.Val < v2.Val
    } else {
        return false
    }
}

func (s SVal) SnapshotString() string {
    return ""
}

func (s SVal) String() string {
    return s.Val
}

