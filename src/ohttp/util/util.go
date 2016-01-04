// Copyright (c) 2015-2016 Kerem Güneş
//   <http://qeremy.com>
//
// GNU General Public License v3.0
//   <http://www.gnu.org/licenses/gpl-3.0.txt>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// @package    ohttp
// @subpackage ohttp.util
// @imports    fmt, strings, strconv, encoding.json, regexp, errors
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package util

import (
    _fmt "fmt"
    _str "strings"; _strc "strconv"
    _url "net/url"
    _json "encoding/json"
    _re "regexp"
    _err "errors"
)

// CRLF.
// @const string
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
    var f string
    for _, arg := range args {
        _ = arg // silence..
        f += "%+v "
    }
    _fmt.Printf("%s\n", _fmt.Sprintf(f, args...))
}

// Dump types.
//
// @param  args... interface{}
// @return (void)
func Dumpt(args... interface{}) {
    var f string
    for _, arg := range args {
        _ = arg // silence..
        f += "%T "
    }
    _fmt.Printf("%s\n", _fmt.Sprintf(f, args...))
}

// Dump as formatted string.
//
// @param  f       string
// @param  args... interface{}
// @return (void)
func Dumpf(f string, args... interface{}) {
    _fmt.Printf("%s\n", _fmt.Sprintf(f, args...))
}

// Error.
//
// @param  f   string
// @param  err error
// @return (error)
func Error(f string, err error) (error) {
    if err == nil {
        return _err.New(f)
    }
    return _fmt.Errorf(f, err)
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
// @param  i interface{}
// @return (bool)
func IsEmpty(i interface{}) (bool) {
    return (i == nil || i == "" || i == 0)
}

// Check empty & set default value.
//
// @param  i  interface{}
// @param  id interface{} Default value.
// @return (interface{})
func IsEmptySet(i, id interface{}) (interface{}) {
    if IsEmpty(i) {
        i = id
    }
    return i
}

// Number converter..
//
// @param  i  interface{}
// @param  it string
// @return (interface{})
func Number(i interface{}, it string) (interface{}) {
    if i != nil {
        n, err := _strc.Atoi(String(i))
        if err == nil {
            switch it {
                // signed
                case    "int": return int(n)
                case   "int8": return int8(n)
                case  "int16": return int16(n)
                case  "int32": return int32(n)
                case  "int64": return int64(n)
                // unsigned
                case   "uint": return uint(n)
                case  "uint8": return uint8(n)
                case "uint16": return uint16(n)
                case "uint32": return uint32(n)
                case "uint64": return uint64(n)
                // float
                case "float32": return float32(n)
                case "float64": return float64(n)
            }
        }
    }
    return nil
}

// Int converter..
//
// @param  i interface{}
// @return (int)
func Int(i interface{}) (int) {
    if n := Number(i, "int"); n != nil {
        return n.(int)
    }
    return 0
}

// UInt converter..
//
// @param  i interface{}
// @return (uint)
func UInt(i interface{}) (uint) {
    if n := Number(i, "uint"); n != nil {
        return n.(uint)
    }
    return 0
}

// Bool converter..
//
// @param  i interface{}
// @return (bool)
func Bool(i interface{}) (bool) {
    if r := String(i); r == "true" || r == "1" {
        return true
    }
    return false
}

// String converter.
//
// @param  i interface{}
// @return (string)
// @panics
func String(i interface{}) (string) {
    switch i.(type) {
        case nil:
            return ""
        case int, bool, string:
            return _fmt.Sprintf("%v", i)
        default:
            it := TypeReal(i)
            // check numerics
            if RegExpTest(it, "u?int(\\d+)?|float(32|64)") {
                return _fmt.Sprintf("%v", i)
            }
            panic("Unsupported input type '"+ it +"' given!")
    }
}

// String format.
//
// @param  f string
// @param  args... interface{}
// @return (string)
func StringFormat(f string, args... interface{}) (string) {
    return _fmt.Sprintf(f, args...)
}

// String search.
//
// @param  s  string
// @param  ss string
// @return (bool)
func StringSearch(s, ss string) (bool) {
    return _str.Index(s, ss) > -1
}

// URL encode.
//
// @param  s string
// @return (string)
func UrlEncode(s string) (string) {
    return _url.QueryEscape(s)
}

// URL decode.
//
// @param  s string
// @return (string)
func UrlDecode(s string) (string) {
    s, err := _url.QueryUnescape(s)
    if err != nil {
        return ""
    }
    return s
}

// Parse URL query.
//
// @param  q string
// @return (map[string]string)
func UrlQueryParse(q string) (map[string]string) {
    r := MapString()
    if tmp := _str.Split(q, "&"); len(tmp) >= 2 {
        for _, tm := range tmp {
            if t := _str.SplitN(tm, "=", 2); len(t) == 2 {
                r[t[0]] = t[1]
            }
        }
    }
    return r
}

// Unparse URL query.
//
// @param  q map[string]interface{}
// @return (string)
func UrlQueryUnparse(q map[string]interface{}) (string) {
    r := MapStringSlice(0)
    for k, v := range q {
        r = append(r, _joinKeyValue(k, v))
    }
    return Implode(r, "&")
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
// @param  s string
// @return (string)
func Upper(s string) (string) {
    return _str.ToUpper(s)
}
func Lower(s string) (string) {
    return _str.ToLower(s)
}

// RegExp test.
//
// @param  s  string
// @param  sr string
// @return (bool)
func RegExpTest(s, sr string) (bool) {
    re, _ := _re.Compile(sr)
    if re == nil {
        return false
    }
    return ("" != re.FindString(s))
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

// RegExp match name.
//
// @param  s  string
// @param  sr string
// @return (map[string]string, *regexp.Regexp, error)
func RegExpMatchName(s, sr string) (map[string]string, *_re.Regexp, error) {
    m, re, err := RegExpMatch(s, sr)
    if err != nil {
        return nil, re, err
    }
    r := MapString()
    for i, name := range re.SubexpNames() {
        if i != 0 { // pass re input
            r[name] = m[i]
        }
    }
    return r, re, nil
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
// @param  i interface{}
// @return ([]int)
func MapIntSlice(i interface{}) ([]int) {
    l := _length(i)
    if l != -1 {
        return make([]int, l)
    }
    return []int{}
}

// Map string slice maker.
//
// @param  i interface{}
// @return ([]string)
func MapStringSlice(i interface{}) ([]string) {
    l := _length(i)
    if l != -1 {
        return make([]string, l)
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
    r := _str.SplitN(i, s, n)
    if len(r) < 2 {
        return nil
    }
    return r
}

// Implode.
//
// @param  i interface{}
// @param  s string
// @param  n int
// @return ([]string)
func Implode(i interface{}, s string) (string) {
    var r string
    switch iv := i.(type) {
        case []int:
            for _, v := range iv {
                r += String(v) + s
            }
            r = r[: len(r) -1]
        case []string:
            r = _str.Join(iv, s)
    }
    return r
}

// Quote.
//
// @param  s string
// @return (string)
func Quote(s string) (string) {
    return _strc.Quote(s)
}

// Detect length.
//
// @param  l interface{}
// @return (int)
// @private
func _length(l interface{}) (int) {
    switch l.(type) {
        case int:
            return l.(int)
        case []int:
            return len(l.([]int))
        case []string:
            return len(l.([]string))
        case []interface{}:
            return len(l.([]interface{}))
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
