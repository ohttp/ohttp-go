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
// @subpackage ohttp.response
// @imports    ohttp.util, ohttp.message, ohttp.response.status
// @author     Kerem Güneş <qeremy[at]gmail[dot]com>
package response

import (
    "ohttp/util"
    "ohttp/message"
    "ohttp/response/status"
)

// @object ohttp.response.Response
type Response struct {
    message.Message // extends
    status          *status.Status
}

func Shutup() {}

// Constructor.
//
// @param  ohttp.message.Message
// @return (*ohttp.response.Response)
func New(m message.Message) (*Response) {
    return &Response{
        Message: m,
        status: status.New(0, "", ""),
    }
}

// Get: status.
//
// @return (*ohttp.response.status.Status)
func (this *Response) Status() (*status.Status) {
    return this.status
}

// Set: status.
//
// @param  s string
// @return (*ohttp.response.Response)
// @panics
func (this *Response) SetStatus(s string) (*Response) {
    m, _, err := util.RegExpMatch(s, "^HTTP/\\d+\\.\\d+\\s+(\\d+)\\s+(.+)")
    if err != nil {
        panic(err)
    }

    if len(m) == 3 {
        this.status.SetStatus(s)
        this.status.SetCode(util.Int(m[1]))
        this.status.SetText(util.Trim(m[2], ""))
    }

    return this
}

// Send.
//
// @shablon For server implementation later..
// func (this *Response) Send() (error) {}

// Get: state
//
// @return (bool)
func (this *Response) OK() (bool) {
    return (this.Error().Code() == 0 && this.Error().Text() == "")
}

// Get: as string.
//
// @return (string)
// @implements
func (this *Response) String() (string) {
    return this.ToString(this.TheStatusLine() + util.CRLF)
}

// Get: the status-line
//
// @return (string)
func (this *Response) TheStatusLine() (string) {
    return this.status.Status()
}
