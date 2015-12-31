// @link https://gobyexample.com/collection-functions
package sarray

func Index(arr []string, is string) (int) {
    for i, v := range arr {
        if v == is {
            return i
        }
    }
    return -1
}

func Has(arr []string, is string) (bool) {
    return Index(arr, is) > -1
}

func Find(arr []string, is string) (string, bool) {
    if i := Index(arr, is); i > -1 {
        return arr[i], true
    }
    return "", false
}

func FindIndex(arr []string, is int) (string, bool) {
    for i, v := range arr {
        if i == is {
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

func Uniq(arr []string) ([]string) {
    ret, f := []string{}, map[string]bool{}
    for k := range arr {
        if f[arr[k]] != true {
            f[arr[k]] = true
            ret = append(ret, arr[k])
        }
    }
    return ret
}
