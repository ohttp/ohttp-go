package params

import (
    "util"
)

func New(argv... interface{}) (map[string]interface{}) {
    argc := len(argv)
    if argc % 2 != 0 {
        panic("Wrong param count (key1, value1, key2, value2)!")
    }

    ret := util.Map()
    for i := 1; i < argc; i += 2 { // tricky?
        if key, ok := argv[i - 1].(string); ok && key != "" {
            ret[key] = argv[i]
            continue
        }
        panic("Each param key must be string and not empty!")
    }

    return ret
}

func NewString(argv... interface{}) (map[string]string) {
    argc := len(argv)
    if argc % 2 != 0 {
        panic("Wrong param count (key1, value1, key2, value2)!")
    }

    ret := util.MapString()
    for i := 1; i < argc; i += 2 { // tricky?
        if key, ok := argv[i - 1].(string); ok && key != "" {
            ret[key] = util.String(argv[i])
            continue
        }
        panic("Each param key must be string and not empty!")
    }

    return ret
}
