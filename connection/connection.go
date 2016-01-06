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
// @subpackage ohttp.connection
// @imports    fmt, net, crypto.tls
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package connection

import (
    _fmt "fmt"
    _net "net"
    _tls "crypto/tls"
)

import (
    "ohttp/uri"
)

// Schemes.
// @const string
const (
    SCHEME_HTTP     = "http"
    SCHEME_HTTPS    = "https"
)

// Ports.
// @const uint
const (
    PORT_HTTP  uint = 80
    PORT_HTTPS      = 443
)

// Dial via HTTP or HTTPS scheme.
//
// @param  u *ohttp.uri.Uri
// @return (net.Conn, error)
// @panics
func Dial(u *uri.Uri) (_net.Conn, error) {
    h, s, p := u.Host(), u.Scheme(), u.Port()

    if p == PORT_HTTP {
        return DialHttp(_fmt.Sprintf("%s:%v", h, p))
    } else if p == PORT_HTTPS {
        return DialHttps(_fmt.Sprintf("%s:%v", h, p))
    } else if p != 0 {
        return DialHttp(_fmt.Sprintf("%s:%v", h, p))
    }

    if s == SCHEME_HTTP {
        return DialHttp(_fmt.Sprintf("%s:%s", h, s))
    } else if s == SCHEME_HTTPS {
        return DialHttps(_fmt.Sprintf("%s:%s", h, s))
    } else if s != "" {
        return DialHttp(_fmt.Sprintf("%s:%s", h, s))
    }

    panic("Scheme and/or port are required!")
}

// Dial via HTTP scheme.
//
// @param  u string
// @return (net.Conn, error)
func DialHttp(u string) (_net.Conn, error) {
    return _net.Dial("tcp", u)
}

// Dial via HTTPS scheme.
//
// @param  u string
// @return (*_tls.Conn, error)
func DialHttps(u string) (*_tls.Conn, error) {
    return _tls.Dial("tcp", u, nil)
}
