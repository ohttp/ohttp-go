package util

import (
    _fmt "fmt"
    _rex "regexp"
    _str "strings"; _strc "strconv"
    _url "net/url"
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

// Dump types.
//
// @param  args... interface{}
// @return (void)
func Dumpt(args... interface{}) {
    var format string
    for _, arg := range args {
        _ = arg // silence..
        format += "%T "
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

// Get short type.
//
// @param  args.. interface{}
// @return (string)
func Type(args... interface{}) (string) {
    return _str.Trim(TypeReal(args[0]), " *<>{}[]")
}

// Get real type.
//
// @param  args.. interface{}
// @return (string)
func TypeReal(args... interface{}) (string) {
    return _str.Replace(_fmt.Sprintf("%T", args[0]), " ", "", -1)
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

// String converter.
//
// @param  input interface{}
// @return (string)
// @panics
func String(input interface{}) (string) {
    switch input.(type) {
        case nil:
            return ""
        case int, bool, string:
            return _fmt.Sprintf("%v", input)
        default:
            inputType := _fmt.Sprintf("%T", input)
            // check numerics
            if StringSearch(inputType, "u?int(\\d+)?|float(32|64)") {
                return _fmt.Sprintf("%v", input)
            }
            panic("Unsupported input type '"+ inputType +"' given!")
    }
}

// String format.
//
// @param  format string
// @param  args... interface{}
// @return (string)
func StringFormat(format string, args... interface{}) (string) {
    return _fmt.Sprintf(format, args...)
}

// String search.
//
// @param  format string
// @param  search string
// @return (bool)
func StringSearch(input, search string) (bool) {
    re, _ := _rex.Compile(search)
    if re == nil {
        return false
    }

    return ("" != re.FindString(input))
}

// URL encode.
//
// @param  input string
// @return (string)
func UrlEncode(input string) (string) {
    return _url.QueryEscape(input)
}

// URL decode.
//
// @param  input string
// @return (string)
func UrlDecode(input string) (string) {
    input, err := _url.QueryUnescape(input)
    if err != nil {
        return ""
    }

    return input
}

// Parse query.
//
// @param  query string
// @return (map[string]string)
func ParseQuery(query string) (map[string]string) {
    ret := MapString()
    if tmps := _str.Split(query, "&"); tmps != nil {
        for _, tmp := range tmps {
            if t := _str.SplitN(tmp, "=", 2); len(t) == 2 {
                ret[t[0]] = t[1]
            }
        }
    }

    return ret
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

// Map int slice maker.
//
// @param  length interface{}
// @return ([]int)
func MapIntSlice(length interface{}) ([]int) {
    len := _length(length)
    if len != -1 {
        return make([]int, len)
    }
    return []int{}
}

// Map string slice maker.
//
// @param  length interface{}
// @return ([]string)
func MapStringSlice(length interface{}) ([]string) {
    len := _length(length)
    if len != -1 {
        return make([]string, len)
    }
    return []string{}
}

// Detect length.
//
// @param  length interface{}
// @return (int)
// @private
func _length(length interface{}) (int) {
    switch length.(type) {
        case int:
            return length.(int)
        case []int:
            return len(length.([]int))
        case []string:
            return len(length.([]string))
        case []interface{}:
            return len(length.([]interface{}))
        // case:
            // @todo add more cases if needs
    }
    return -1
}
