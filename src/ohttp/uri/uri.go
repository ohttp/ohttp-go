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
// @subpackage ohttp.uri
// @imports    net.url, strings
// @imports    ohttp.util, ohttp.util.query, ohttp.util.array.sarray
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package uri

import (
    _url "net/url"
    _str "strings"
)

import (
    "ohttp/util"
    "ohttp/util/query"
    "ohttp/util/array/sarray"
)

// @object ohttp.Uri
type Uri struct {
    source     string
    scheme     string
    host       string
    port       uint
    path       string
    username   string
    password   string
    query      *query.Query
    fragment   string
    segments   []string
}

// Constructor.
//
// @param  s string
// @param  q interface{}
// @return (*ohttp.Uri)
func New(s string, q interface{}) (*Uri) {
    this := &Uri{
        source: s,
    }

    if s != "" {
        s, _ := _url.Parse(util.UrlDecode(s))
        // set scheme
        if ss := s.Scheme; ss != "" {
            this.scheme = ss
        }

        // set host
        if sh := s.Host; sh != "" {
            this.host = sh
            // set port yourself (cos url.Parse doesn't provide it..)
            if tmp := _str.Split(sh, ":"); len(tmp) == 2 {
                this.host = tmp[0]
                this.port = util.UInt(tmp[1])
            }
        }

        // set path
        if sp := s.Path; sp != "" {
            this.path = sp
            // set segments @note will be used later for routing operations
            if seg := _str.Split(sp, "/"); len(seg) > 0 {
                this.segments = util.MapStringSlice(seg)
                for i, se := range seg {
                    se = _str.TrimSpace(se)
                    if se != "" {
                        this.segments[i] = se
                    }
                }
                // remove empty segments
                this.segments = sarray.Filter(this.segments, nil)
            }
        }

        // set username & password
        if s.User != nil {
            this.username = s.User.Username()
            this.password, _ = s.User.Password()
        }

        // set query
        if sq := s.RawQuery; sq != "" {
            this.query = query.New(sq)
        } else {
            this.query = query.New(q)
        }

        // set fragment
        if sf := s.Fragment; sf != "" {
            this.fragment = sf
        }
    }

    return this
}

// Get: source.
//
// @return (string)
func (this *Uri) Source() (string) {
    return this.source
}

// Get: scheme.
//
// @return (string)
func (this *Uri) Scheme() (string) {
    return this.scheme
}

// Get: host.
//
// @return (string)
func (this *Uri) Host() (string) {
    return this.host
}

// Get: port.
//
// @return (uint)
func (this *Uri) Port() (uint) {
    return this.port
}

// Get: path.
//
// @return (string)
func (this *Uri) Path() (string) {
    return this.path
}

// Get: username.
//
// @return (string)
func (this *Uri) Username() (string) {
    return this.username
}

// Get: password.
//
// @return (string)
func (this *Uri) Password() (string) {
    return this.password
}

// Get: query.
//
// @return (*ohttp.util.query.Query)
func (this *Uri) Query() (*query.Query) {
    return this.query
}

// Get: fragment.
//
// @return (string)
func (this *Uri) Fragment() (string) {
    return this.fragment
}

// Get: segments.
//
// @return ([]string)
func (this *Uri) Segments() ([]string) {
    return this.segments
}

// Get: segment.
//
// @return (string)
func (this *Uri) Segment(i int) (string) {
    if se, ok := sarray.FindIndex(this.segments, i); ok {
        return se
    }
    return ""
}

// Get: authorization.
//
// @return (string)
func (this *Uri) Authorization() (string) {
    if this.username != "" && this.password != "" {
        return this.username +":"+ this.password
    } else if this.username != "" {
        return this.username
    }
    return ""
}
