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
// @subpackage ohttp.util.list.slist
// @link       https://gobyexample.com/collection-functions
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package slist

// Index.
//
// @param  arr []string
// @param  s   string
// @return (int)
func Index(arr []string, is string) (int) {
    for i, v := range arr {
        if v == is {
            return i
        }
    }
    return -1
}

// Has.
//
// @param  arr []string
// @param  s   string
// @return (bool)
func Has(arr []string, is string) (bool) {
    return Index(arr, is) > -1
}

// Find.
//
// @param  arr []string
// @param  s   string
// @return (string, bool)
func Find(arr []string, is string) (string, bool) {
    if i := Index(arr, is); i > -1 {
        return arr[i], true
    }
    return "", false
}

// Find by index.
//
// @param  arr []string
// @param  is  int
// @return (string, bool)
func FindIndex(arr []string, is int) (string, bool) {
    for i, v := range arr {
        if i == is {
            return v, true
        }
    }
    return "", false
}

// Map.
//
// @param  arr []string
// @param  fn  func(string, int) (string)
// @return ([]string)
func Map(arr []string, fun func(string, int) (string)) ([]string) {
    ret := make([]string, len(arr))
    for k, v := range arr {
        ret[k] = fun(v, k)
    }
    return ret
}

// Filter.
//
// @param  arr []string
// @param  fn  func(string, int) (bool)
// @return ([]string)
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

// Uniq.
//
// @param  arr []string
// @return ([]string)
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

// Shift.
//
// @param  arr *[]string
// @return (string)
func Shift(arr *[]string) (string) {
    v := (*arr)[0]
    *arr = (*arr)[1 : len(*arr)]
    return v
}
