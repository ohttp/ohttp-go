// @link https://gobyexample.com/collection-functions
package iarray

func Index(arr []int, s int) int {
    for i, val := range arr {
        if val == s {
            return i
        }
    }
    return -1
}

func Has(arr []int, s int) bool {
    return Index(arr, s) >= 0
}

func Find(arr[]int, s int) (int, bool) {
    i := Index(arr, s)
    if i >= 0 {
        return arr[i], true
    }
    return 0, false
}

func Map(arr []int, fun func(int) int) []int {
    ret := make([]int, len(arr))
    for i, val := range arr {
        ret[i] = fun(val)
    }
    return ret
}

func Filter(arr []int, fun func(int) bool) []int {
    if fun == nil {
        fun = func(s int) bool {
            return s != 0
        }
    }
    ret := make([]int, 0)
    for _, val := range arr {
        if fun(val) {
            ret = append(ret, val)
        }
    }
    return ret
}
