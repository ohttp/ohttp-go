// @link https://gobyexample.com/collection-functions
package sarray

func Index(arr []string, s string) (int) {
    for i, v := range arr {
        if v == s {
            return i
        }
    }
    return -1
}

func Has(arr []string, s string) (bool) {
    return Index(arr, s) > -1
}

func Find(arr []string, s string) (string, bool) {
    if i := Index(arr, s); i > -1 {
        return arr[i], true
    }
    return "", false
}

func FindIndex(arr []string, s int) (string, bool) {
    for i, v := range arr {
        if i == s {
            return v, true
        }
    }
    return "", false
}

func Map(arr []string, fun func(string, int) (string)) ([]string) {
    ret := make([]string, len(arr))
    for k, v := range arr {
        ret[k] = fun(v, k)
    }
    return ret
}

func Filter(arr []string, fun func(string, int) (bool)) ([]string) {
    if fun == nil {
        fun = func(v string, k int) (bool) {
            return v != ""
        }
    }
    ret := make([]string, 0)
    for k, v := range arr {
        if fun(v, k) {
            ret = append(ret, v)
        }
    }
    return ret
}
