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
// @subpackage ohttp.params
// @imports    ohttp.util
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package params

import (
    "ohttp/util"
)

// @object ohttp.params.Params
type Params map[string]interface{}

// Constructor.
//
// @return (ohttp.params.Params)
func New() (*Params) {
    return &Params{}
}

// Set.
//
// @param  k string
// @param  v interface{}
// @return (ohttp.params.Params)
func (this *Params) Set(k string, v interface{}) (*Params) {
    (*this)[k] = v
    return this
}

// Get.
//
// @param  k string
// @return (interface{})
func (this *Params) Get(k string) (interface{}) {
    return (*this)[k]
}

// Get: int.
//
// @param  k string
// @return (int)
func (this *Params) GetInt(k string) (int) {
    return util.Int((*this)[k])
}

// Get: uint.
//
// @param  k string
// @return (uint)
func (this *Params) GetUInt(k string) (uint) {
    return util.UInt((*this)[k])
}

// Get: string.
//
// @param  k string
// @return (string)
func (this *Params) GetString(k string) (string) {
    return util.String((*this)[k])
}

// Get: bool.
//
// @param  k string
// @return (bool)
func (this *Params) GetBool(k string) (bool) {
    return util.Bool((*this)[k])
}

// Check empty.
//
// @return (bool)
func (this *Params) Empty() (bool) {
    for k, v := range (*this) {
        _ = k; _ = v
        return false
    }
    return true
}

// Stringfy.
//
// @return (int)
func (this *Params) String() (string) {
    m := util.Map()
    for k, v := range (*this) {
        m[k] = v
    }
    return util.UrlQueryUnparse(m)
}
