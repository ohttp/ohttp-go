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

func Map(arr []string, fun func(string) string) ([]string) {
    ret := make([]string, len(arr))
    for i, v := range arr {
        ret[i] = fun(v)
    }
    return ret
}

func Filter(arr []string, fun func(string) bool) ([]string) {
    if fun == nil {
        fun = func(s string) (bool) {
            return s != ""
        }
    }
    ret := make([]string, 0)
    for _, v := range arr {
        if fun(v) {
            ret = append(ret, v)
        }
    }
    return ret
}
