// @link https://gobyexample.com/collection-functions
package scol

func Index(scol []string, s string) int {
    for i, val := range scol {
        if val == s {
            return i
        }
    }
    return -1
}

func Has(scol []string, s string) bool {
    return Index(scol, s) >= 0
}

func Find(scol []string, s string) (string, bool) {
    i := Index(scol, s)
    if i >= 0 {
        return scol[i], true
    }
    return "", false
}

func FindIndex(scol []string, s int) (string, bool) {
    for i, val := range scol {
        if i == s {
            return val, true
        }
    }
    return "", false
}

func Map(scol []string, fun func(string) string) []string {
    ret := make([]string, len(scol))
    for i, val := range scol {
        ret[i] = fun(val)
    }
    return ret
}

func Filter(scol []string, fun func(string) bool) []string {
    if fun == nil {
        fun = func(s string) bool {
            return s != ""
        }
    }
    ret := make([]string, 0)
    for _, val := range scol {
        if fun(val) {
            ret = append(ret, val)
        }
    }
    return ret
}
