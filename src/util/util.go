package util

import (
    _fmt "fmt"
)

func Shutup() {}

// Dump.
//
// @param  args... interface{}
// @return (void)
func Dump(args... interface{}) {
    _fmt.Println(args...)
}

// Dump as string.
//
// @param  args... interface{}
// @return (void)
func Dumps(args... interface{}) {
    var format string
    for _, arg := range args {
        _ = arg // silence..
        format += "%+v "
    }
    _fmt.Printf("%s\n", _fmt.Sprintf(format, args...))
}

// Dump as formatted string.
//
// @param  format  string
// @param  args... interface{}
// @return (void)
func Dumpf(format string, args... interface{}) {
    if format == "" {
        for _, arg := range args {
            _ = arg // silence..
            format += "%+v "
        }
    }
    _fmt.Printf("%s\n", _fmt.Sprintf(format, args...))
}

// Map maker.
//
// @return (map[string]interface{})
func Map() (map[string]interface{}) {
    return make(map[string]interface{})
}

// Int map maker.
//
// @return (map[int]string)
func MapInt() (map[int]string) {
    return make(map[int]string)
}

// String map maker.
//
// @return (map[string]string)
func MapString() (map[string]string) {
    return make(map[string]string)
}
