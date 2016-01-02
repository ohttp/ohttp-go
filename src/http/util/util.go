package util

import (
    _fmt "fmt"
    _str "strings"; _strc "strconv"
    _url "net/url"
    _json "encoding/json"
    _re "regexp"
)

const (
    CRLF = "\r\n"
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

// Check is empty.
//
// @param  input interface{}
// @return (bool)
func IsEmpty(input interface{}) (bool) {
    return (input == nil || input == "" || input == 0)
}

// Check empty & set default value.
//
// @param  input        interface{}
// @param  inputDefault interface{}
// @return (bool)
func IsEmptySet(input, inputDefault interface{}) (interface{}) {
    if IsEmpty(input) {
        input = inputDefault
    }
    return input
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

// Bool converter..
//
// @param  input interface{}
// @return (bool)
func Bool(i interface{}) (bool) {
    if r := String(i); r == "true" || r == "1" {
        return true
    }
    return false
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
            inputType := TypeReal(input)
            // check numerics
            if RegExpTest(inputType, "u?int(\\d+)?|float(32|64)") {
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
    return _str.Index(input, search) > -1
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

// Parse URL query.
//
// @param  q string
// @return (map[string]string)
func UrlQueryParse(q string) (map[string]string) {
    ret := MapString()
    if tmp := _str.Split(q, "&"); tmp != nil {
        for _, tm := range tmp {
            if t := _str.SplitN(tm, "=", 2); len(t) == 2 {
                ret[t[0]] = t[1]
            }
        }
    }
    return ret
}

// Unparse URL query.
//
// @param  q map[string]interface{}
// @return (string)
func UrlQueryUnparse(q map[string]interface{}) (string) {
    m := MapStringSlice(0)
    for k, v := range q {
        m = append(m, _joinKeyValue(k, v))
    }
    return Implode(m, "&")
}

// JSON encode.
//
// @param  in interface{}
// @return (string, error)
func JsonEncode(in interface{}) (string, error) {
    out, err := _json.Marshal(in)
    if err != nil {
        return "", _fmt.Errorf("JSON error: %s!", err)
    }
    return string(out), nil
}

// JSON decode.
//
// @param  in  string
// @param  out interface{}
// @return (interface{}, error)
func JsonDecode(in string, out interface{}) (interface{}, error) {
    // simply prevent useless unmarshal error
    if in == "" {
        in = `null`
    }
    err := _json.Unmarshal([]byte(in), &out)
    if err != nil {
        return nil, _fmt.Errorf("JSON error: %s!", err)
    }
    return out, nil
}

// String upper/lower
//
// @param  input string
// @return (string)
func Upper(input string) (string) {
    return _str.ToUpper(input)
}
func Lower(input string) (string) {
    return _str.ToLower(input)
}

// RegExp match.
//
// @param  s  string
// @param  sr string
// @return ([]string, *regexp.Regexp, error)
func RegExpMatch(s, sr string) ([]string, *_re.Regexp, error) {
    re, err := _re.Compile(sr)
    if err != nil {
        return nil, re, err
    }
    return re.FindStringSubmatch(s), re, nil
}

// RegExp match all.
//
// @param  s  string
// @param  sr string
// @return (map[string]string, *regexp.Regexp, error)
func RegExpMatchAll(s, sr string) (map[string]string, *_re.Regexp, error) {
    m, re, err := RegExpMatch(s, sr)
    if err != nil {
        return nil, re, err
    }
    ret := MapString()
    for i, n := range re.SubexpNames() {
        if i != 0 { // pass re input
            ret[n] = m[i]
        }
    }
    return ret, re, nil
}

// RegExp test.
//
// @param  format string
// @param  search string
// @return (bool)
func RegExpTest(input, search string) (bool) {
    re, _ := _re.Compile(search)
    if re == nil {
        return false
    }
    return ("" != re.FindString(input))
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

// Trim.
//
// @param  s  string
// @param  sc string
// @return (string)
func Trim(s, sc string) (string) {
    if sc == "" {
        return _str.TrimSpace(s)
    }
    return _str.Trim(s, sc)
}

// Explode.
//
// @param  i string
// @param  s string
// @param  n int
// @return ([]string)
func Explode(i, s string, n int) ([]string) {
    ret := _str.SplitN(i, s, n)
    if len(ret) < 2 {
        return nil
    }
    return ret
}

// Implode.
//
// @param  i interface{}
// @param  s string
// @param  n int
// @return ([]string)
func Implode(i interface{}, s string) (string) {
    var ret string
    switch iv := i.(type) {
        case []int:
            for _, v := range iv {
                ret += String(v) + s
            }
            ret = ret[: len(ret) -1]
        case []string:
            ret = _str.Join(iv, s)
    }
    return ret
}

// Quote.
//
// @param  input string
// @return (string)
func Quote(input string) (string) {
    return _strc.Quote(input)
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

// Join key => value pairs (only 2-dims arrays).
//
// @param  k string
// @param  v interface{}
// @return (string)
func _joinKeyValue(k string, v interface{}) (string) {
    var s string
    switch v := v.(type) {
        case []int:
            for _, vv := range v {
                s += StringFormat("%s[]=%v&", UrlEncode(k), vv)
            }
            break
        case []string:
            for _, vv := range v {
                s += StringFormat("%s[]=%v&", UrlEncode(k), vv)
            }
            break
        case map[string]interface{}:
            for kk, vv := range v {
                switch vv := vv.(type) {
                    case []int, []string:
                        s += _joinKeyValue(k, vv)
                        break
                    case map[string]interface{}:
                        for kkk, vvv := range vv {
                            s += _joinKeyValue(kkk, vvv)
                        }
                        break
                    default:
                        s += StringFormat("%s[%s]=%v&",
                            UrlEncode(k), UrlEncode(kk), UrlEncode(String(vv)))
                }
            }
            break
        default:
            s = StringFormat("%s=%v", UrlEncode(k), UrlEncode(String(v)))
    }
    return Trim(s, "&")
}
