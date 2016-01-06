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
// @subpackage ohttp.util.array.iarray
// @link       https://gobyexample.com/collection-functions
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package iarray

// Index.
//
// @param  arr []int
// @param  s   int
// @return (int)
func Index(arr []int, s int) (int) {
    for i, v := range arr {
        if v == s {
            return i
        }
    }
    return -1
}

// Has.
//
// @param  arr []int
// @param  s   int
// @return (bool)
func Has(arr []int, s int) (bool) {
    return Index(arr, s) > -1
}

// Find.
//
// @param  arr []int
// @param  s   int
// @return (int, bool)
func Find(arr[]int, s int) (int, bool) {
    if i := Index(arr, s); i > -1 {
        return arr[i], true
    }
    return 0, false
}

// Map.
//
// @param  arr []int
// @param  fn  func(int, int) (int)
// @return ([]int)
func Map(arr []int, fn func(int, int) (int)) ([]int) {
    ret := make([]int, len(arr))
    for k, v := range arr {
        ret[k] = fn(v, k)
    }
    return ret
}

// Filter.
//
// @param  arr []int
// @param  fn  func(int, int) (bool)
// @return ([]int)
func Filter(arr []int, fn func(int, int) (bool)) ([]int) {
    if fn == nil {
        fn = func(v, k int) (bool) {
            return v != 0
        }
    }
    ret := make([]int, 0)
    for k, v := range arr {
        if fn(v, k) {
            ret = append(ret, v)
        }
    }
    return ret
}

// Uniq.
//
// @param  arr []int
// @return ([]int)
func Uniq(arr []int) ([]int) {
    ret, f := []int{}, map[int]bool{}
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
// @param  arr *[]int
// @return (int)
func Shift(arr *[]int) (int) {
    v := (*arr)[0]
    *arr = (*arr)[1:len(*arr)]
    return v
}
