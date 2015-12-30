package util

import (
    _fmt "fmt"
    _strc "strconv"
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

// Int converter..
//
// @param  input interface{}
// @return (int)
func Int(input interface{}) (int) {
    if number := Number(input, "int"); number != nil {
        return number.(int)
    }
    return 0
}

// UInt converter..
//
// @param  input interface{}
// @return (uint)
func UInt(input interface{}) (uint) {
    if number := Number(input, "uint"); number != nil {
        return number.(uint)
    }
    return 0
}

// Number converter..
//
// @param  input     interface{}
// @param  inputType string
// @return (interface{})
func Number(input interface{}, inputType string) (interface{}) {
    if input != nil {
        number, err := _strc.Atoi(String(input))
        if err == nil {
            switch inputType {
                // signed
                case    "int": return int(number)
                case   "int8": return int8(number)
                case  "int16": return int16(number)
                case  "int32": return int32(number)
                case  "int64": return int64(number)
                // unsigned
                case   "uint": return uint(number)
                case  "uint8": return uint8(number)
                case "uint16": return uint16(number)
                case "uint32": return uint32(number)
                case "uint64": return uint64(number)
                // float
                case "float32": return float32(number)
                case "float64": return float64(number)
            }
        }
    }
    return nil
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
