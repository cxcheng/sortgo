package sortlib

import "bufio"
import "log"
import "os"

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

func SVals(strs []string) []Val {
    vals := make([]Val, len(strs), len(strs))
    for i, v := range(strs) {
        vals[i] = SVal{ Val: v }
    }
    return vals
}

func SValsFromFile(path string) []Val {
    vals := make([]Val, 0, 1000)
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
        return vals
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if str := scanner.Text(); len(str) > 0 {
            vals = append(vals, SVal{ Val: str})
        }
    }
    return vals
}