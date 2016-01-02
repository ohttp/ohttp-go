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
    return Index(arr, s) > -1
}

func Find(arr[]int, s int) (int, bool) {
    if i := Index(arr, s); i > -1 {
        return arr[i], true
    }
    return 0, false
}

func Map(arr []int, fun func(int, int) (int)) ([]int) {
    ret := make([]int, len(arr))
    for k, v := range arr {
        ret[k] = fun(v, k)
    }
    return ret
}

func Filter(arr []int, fun func(int, int) (bool)) ([]int) {
    if fun == nil {
        fun = func(v, k int) (bool) {
            return v != 0
        }
    }
    ret := make([]int, 0)
    for k, v := range arr {
        if fun(v, k) {
            ret = append(ret, v)
        }
    }
    return ret
}

func Uniq(arr []int) ([]int) {
    ret, f := []int{}, map[int]bool{}
    for k := range arr {
        if f[arr[k]] != true {
            f[arr[k]] = true
            ret = append(ret, arr[k])
        }
    }
    return ret
}

func Shift(arr *[]int) (int) {
    v := (*arr)[0]
    *arr = (*arr)[1 : len(*arr)]
    return v
}
