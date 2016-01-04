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
// @subpackage ohttp.request
// @imports    fmt, bufio
// @imports    ohttp.util, ohttp.message, ohttp.uri, ohttp.useragent, ohttp.connection
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package request

import (
    _fmt "fmt"
    _bio "bufio"
)

import (
    "ohttp/util"
    "ohttp/message"
    "ohttp/uri"
    "ohttp/useragent"
    "ohttp/connection"
)

// @object ohttp.request
type Request struct {
    message.Message // extends
    method          string
    uri             *uri.Uri
}

func Shutup() {}

// Constructor.
//
// @param  m *ohttp.message.Message
// @return (*ohttp.request.Request)
func New(m message.Message) (*Request) {
    return &Request{
        Message: m,
    }
}

// Get: method.
//
// @return (string)
func (this *Request) Method() (string) {
    return this.method
}

// Get: uri.
//
// @return (*ohttp.uri.Uri)
func (this *Request) Uri() (*uri.Uri) {
    return this.uri
}

// Set: method.
//
// @param  m string
// @return (*ohttp.request.Request)
func (this *Request) SetMethod(m string) (*Request) {
    this.method = util.Upper(m)
    return this
}

// Set: uri.
//
// @param  u  string
// @param  up interface{}
// @return (*ohttp.request.Request)
func (this *Request) SetUri(u string, up interface{}) (*Request) {
    this.uri = uri.New(u, up)
    return this
}


// Send.
//
// @return (string, error)
func (this *Request) Send() (string, error) {
    debug := this.Options().GetBool("debug")

    this.SetHeader("Host", this.uri.Host())
    this.SetHeader("Connection", "close")
    if this.Header("User-Agent") == "" {
        this.SetHeader("User-Agent", _fmt.Sprintf("%s/v%s (+%s)",
            useragent.OH_NAME, useragent.OH_VERSION, useragent.OH_LINK))
    }

    link, err := connection.Dial(this.uri)
    if err != nil {
        return "", err
    }
    defer link.Close()

    var rs, rr string
    rs += this.TheRequestLine()
    for k, v := range this.HeaderAll() {
        if v != "" {
            rs += _fmt.Sprintf("%s: %s%s", k, v, util.CRLF)
        }
    }
    rs += util.CRLF
    rs += this.Body().Content()

    _fmt.Fprint(link, rs)

    r := _bio.NewReader(link)
    // status-line
    sl, err := r.ReadString('\n')
    if sl == "" {
        return "", err
    }
    rr += sl

    for {
        b := make([]byte, 1024)
        if bl, _ := r.Read(b); bl == 0 {
            break // eof
        }
        rr += util.Trim(string(b), "\x00")
    }

    link.Close()

    if debug == true {
        util.Dump(rs)
        util.Dump(rr)
    }

    return rr, nil
}

// Get: state
//
// @return (bool)
// @implements
func (this *Request) OK() (bool) {
    return (this.Error().Code() == 0 && this.Error().Text() == "")
}

// Get: as string.
//
// @return (string)
// @implements
func (this *Request) String() (string) {
    return this.ToString(this.TheRequestLine())
}

// Get: the request-line
//
// @return (string)
func (this *Request) TheRequestLine() (string) {
    rm  := this.Method()
    rp  := "/"
    rpv := this.ProtocolVersion()
    if s := this.uri.Path(); s != "" {
        rp = s
    }
    if s := this.uri.Query().String(); s != "" {
        rp += "?"+ s
    }
    return _fmt.Sprintf("%s %s HTTP/%s%s", rm, rp, rpv, util.CRLF)
}
