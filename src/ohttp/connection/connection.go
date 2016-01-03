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

// @package ohttp
// @imports fmt, net, crypto.tls
// @author  Kerem Güneş <qeremy[at]gmail[dot]com>
package connection

import (
    _fmt "fmt"
    _net "net"
    _tls "crypto/tls"
)

// Ports.
// @const uing
const (
    PORT_HTTP  uint = 80
    PORT_HTTPS      = 443
)

func Dial(h string, p uint) (_net.Conn, error) {
    if p == PORT_HTTP {
        return DialHttp(h, PORT_HTTP)
    } else if p == PORT_HTTPS {
        return DialHttps(h, PORT_HTTPS)
    } else {
        p = PORT_HTTP
    }
    return DialHttp(h, p)
}

// Dial via HTTP scheme.
//
// @param  h string
// @param  p uint
// @return (net.Conn, error)
func DialHttp(h string, p uint) (_net.Conn, error) {
    return _net.Dial("tcp", _fmt.Sprintf("%s:%v", h, p))
}

// Dial via HTTPS scheme.
//
// @param  h string
// @param  p uint
// @return (net.Conn, error)
func DialHttps(h string, p uint) (_tls.Conn, error) {
    return _tls.Dial("tcp", _fmt.Sprintf("%s:%v", h, p), nil)
}
