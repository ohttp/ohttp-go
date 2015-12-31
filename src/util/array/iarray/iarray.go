// @link https://gobyexample.com/collection-functions
package iarray

func Index(arr []int, s int) (int) {
    for i, v := range arr {
        if v == s {
            return i
        }
    }
    return -1
}

func Has(arr []int, s int) (bool) {
    return Index(arr, s) >= 0
}

func Find(arr[]int, s int) (int, bool) {
    if i := Index(arr, s); i >= 0 {
        return arr[i], true
    }
    return 0, false
}

func Map(arr []int, fun func(int) int) ([]int) {
    ret := make([]int, len(arr))
    for i, v := range arr {
        ret[i] = fun(v)
    }
    return ret
}

func Filter(arr []int, fun func(int) bool) ([]int) {
    if fun == nil {
        fun = func(s int) (bool) {
            return s != 0
        }
    }
    ret := make([]int, 0)
    for _, v := range arr {
        if fun(v) {
            ret = append(ret, v)
        }
    }
    return ret
}
