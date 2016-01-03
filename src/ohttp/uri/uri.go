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
// @imports    ohttp.util, ohttp.util.query, ohttp.util.array.sarray
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package uri

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

// Valid port range.
// @const uint
const (
    VALID_PORT_MIN uint = 1
    VALID_PORT_MAX      = 65535
)

// Constructor.
//
// @param  s string
// @param  q interface{}
// @return (*ohttp.Uri)
// @panics
func New(s string, q interface{}) (*Uri) {
    this := &Uri{
        source: s,
    }

    if s != "" {
        p := Parse(util.UrlDecode(s))
        // set scheme
        if s := p["Scheme"]; s != "" {
            this.scheme = s
        }

        // set host
        if s := p["Host"]; s != "" {
            this.host = s
        }

        // set port
        if s := p["Port"]; s != "" {
            s := util.UInt(s)
            if s < VALID_PORT_MIN || s > VALID_PORT_MAX {
                panic("Invalid port given!")
            }
            this.port = s
        }

        // set path
        if s := p["Path"]; s != "" {
            this.path = s
            // set segments @note will be used later for routing operations
            if seg := util.Explode(s, "/", -1); len(seg) > 0 {
                this.segments = util.MapStringSlice(seg)
                for i, se := range seg {
                    se = util.Trim(se, "")
                    if se != "" {
                        this.segments[i] = se
                    }
                }
                // remove empty segments
                this.segments = sarray.Filter(this.segments, nil)
            }
        }

        // set username & password
        if p["Username"] != "" {
            this.username = p["Username"]
            this.password = p["Password"]
        }

        // set query
        if s := p["Query"]; s != "" {
            this.query = query.New(s)
        } else {
            this.query = query.New(q)
        }

        // set fragment
        if s := p["Fragment"]; s != "" {
            this.fragment = s
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

// Parse URL.
//
// @param  u string
// @return (map[string]string)
// @panics
func Parse(u string) (map[string]string) {
    if u == "" {
        panic("No URL given!")
    }

    // add http as default
    if u[:4] != "http" {
        u = "http://"+ u
    }

    var p string
    // add scheme
    p = "(?:(?P<Scheme>https?)://)?"

    // check authorization stuff
    if util.RegExpTest(u, "://(.+?)@") {
        p += "(?:(?P<Username>[^:@]+))?(?:(?P<Password>[^@]+))?"
    }

    // append others
    p += "(?:(?P<Host>[^:/]+))?"      +
         "(?:\\:(?P<Port>\\d+))?"     +
         "(?P<Path>/[^?#]*)?"         +
         "(?:\\?(?P<Query>[^#]+))?"   +
         "(?:\\??#(?P<Fragment>.*))?"

    r, _, err := util.RegExpMatchName(u, p)
    if err != nil {
        return nil
    }

    return r
}
