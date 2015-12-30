// @link https://gobyexample.com/collection-functions
package icol

func Index(icol []int, s int) int {
    for i, val := range icol {
        if val == s {
            return i
        }
    }
    return -1
}

func Has(icol []int, s int) bool {
    return Index(icol, s) >= 0
}

func Find(icol[]int, s int) (int, bool) {
    i := Index(icol, s)
    if i >= 0 {
        return icol[i], true
    }
    return 0, false
}

func Map(icol []int, fun func(int) int) []int {
    ret := make([]int, len(icol))
    for i, val := range icol {
        ret[i] = fun(val)
    }
    return ret
}

func Filter(icol []int, fun func(int) bool) []int {
    if fun == nil {
        fun = func(s int) bool {
            return s != 0
        }
    }
    ret := make([]int, 0)
    for _, val := range icol {
        if fun(val) {
            ret = append(ret, val)
        }
    }
    return ret
}
