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
// @subpackage ohttp.util.query
// @imports    ohttp.util, ohttp.util.params
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package query

import (
    "ohttp/util"
    "ohttp/util/params"
)

// @object ohttp.query.Query
type Query struct {
    params   *params.Params
}

func Shutup() {}

// Constructor.
//
// @param  p interface{}
// @return (*ohttp.query.Query)
func New(p interface{}) (*Query) {
    this := &Query{}
    this.params = params.New()
    if _, ok := p.(params.Params); ok {
        for k, v := range p.(params.Params) {
            this.params.Set(k, v)
        }
    } else if _, ok := p.(*params.Params); ok {
        for k, v := range *p.(*params.Params) {
            this.params.Set(k, v)
        }
    } else if _, ok := p.(map[string]interface{}); ok {
        for k, v := range p.(map[string]interface{}) {
            this.params.Set(k, v)
        }
    } else if p, ok := p.(string); ok {
        for k, v := range util.UrlQueryParse(p) {
            this.params.Set(k, v)
        }
    }
    return this
}

// Set.
//
// @param  k string
// @param  v interface{}
// @return (*ohttp.query.Query)
func (this *Query) Set(k string, v interface{}) (*Query) {
    this.params.Set(k, v)
    return this
}

// Get.
//
// @param  k string
// @return v interface{}
func (this *Query) Get(k string) (interface{}) {
    return this.params.Get(k)
}

// Get params object.
//
// @return (*ohttp.util.params.Params)
func (this *Query) Params() (*params.Params) {
    return this.params
}

// Get as string.
//
// @return (string)
func (this *Query) String() (string) {
    if this.params != nil {
        return this.params.String()
    }
    return ""
}
