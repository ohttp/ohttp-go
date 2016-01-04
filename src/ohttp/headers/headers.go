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
// @subpackage ohttp.headers
// @imports    ohttp.util, ohttp.util.array.sarray
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package headers

import (
    "ohttp/util"
    "ohttp/util/array/sarray"
)

// @object ohttp.Headers
type Headers struct {
    data map[string]string
}

// Constructor.
//
// @return (*ohttp.http.Headers)
func New() (*Headers) {
    return &Headers{
        data: util.MapString(),
    }
}

// Set header.
//
// @param  k string
// @param  v string
// @return (*ohttp.header.Headers)
func (this *Headers) Set(k, v string) (*Headers) {
    this.data[k] = util.Trim(v, "")
    return this
}


// Set header.
//
// @param  kv map[string]string
// @return (void)
func (this *Headers) SetAll(kv map[string]string) {
    this.data = data
}


// Get header.
//
// @param  k string
// @return (string)
func (this *Headers) Get(k string) (string) {
    if v, ok := this.data[k]; ok {
        return v
    }
    return ""
}

// Get all headers.
//
// @return (map[string]string)
func (this *Headers) GetAll() (map[string]string) {
    return this.data
}

// Parse.
//
// @params hs string
// @return (map[string]string)
func Parse(hs string) (map[string]string) {
    ret := util.MapString()
    if tmp := util.Explode(hs, util.CRLF, -1); tmp != nil {
        // status line (HTTP/1.0 200 OK)
        ret["0"] = sarray.Shift(&tmp)

        for _, tm := range tmp {
            if t := util.Explode(tm, ":", 2); len(t) == 2 {
                ret[util.Trim(t[0], "")] =  util.Trim(t[1], "")
            }
        }
    }

    return ret
}
