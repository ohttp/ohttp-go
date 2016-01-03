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

// @package ohttp.message
// @imports ohttp.util, ohttp.util.params, ohttp.headers
// @author  Kerem Güneş <qeremy[at]gmail[dot]com>
package message

import (
    "ohttp/util"
    "ohttp/util/params"
    "ohttp/headers"
)

// @object ohttp.message.Message
type Message struct {
    type_           uint
    protocolVersion string
    headers         *headers.Headers
    body            *MessageBody
    bodyData        *MessageBodyData // parsed
    options         *params.Params
    error           *MessageError
    MessageOKInterface
    MessageStringInterface
}

// @object ohttp.message.MessageOKInterface
type MessageOKInterface interface {
    OK() (bool)
}

// @object ohttp.message.MessageStringInterface
type MessageStringInterface interface {
    String() (string)
}

// Message types.
// @const uint
const (
    TYPE_REQUEST  uint = 1
    TYPE_RESPONSE      = 2
)

// Message protocols.
// @const string
const (
    PROTOCOL_VERSION_1_0 = "1.0"
    PROTOCOL_VERSION_1_1 = "1.1"
    PROTOCOL_VERSION_2_0 = "2.0"
)

func Shutup() {}

// Constructor.
//
// @param  t  uint
// @param  pv string
// @param  o  *ohttp.params.Params
// @return (*ohttp.message.Message)
func NewMessage(t uint, pv string, o *params.Params) (*Message) {
    return &Message{
        type_: t,
        protocolVersion: pv,
        headers: headers.New(),
        body: NewMessageBody("", "", 0),
        bodyData: NewMessageBodyData(nil, ""),
        error: NewMessageError(0, ""),
        options: o,
    }
}

// Get: type.
//
// @return (uint)
func (this *Message) Type() (uint) {
    return this.type_
}

// Get: protocol version.
//
// @return string
func (this *Message) ProtocolVersion() (string) {
    return this.protocolVersion
}

// Get: header.
//
// @return string
func (this *Message) Header(k string) (string) {
    return this.headers.Get(k)
}

// Get: all headers.
//
// @return map[string]string
func (this *Message) HeaderAll() (map[string]string) {
    return this.headers.GetAll()
}

// Get: body.
//
// @return (*ohttp.message.MessageBody)
func (this *Message) Body() (*MessageBody) {
    return this.body
}

// Get: body data.
//
// @return (*ohttp.message.MessageBodyData)
func (this *Message) BodyData() (*MessageBodyData) {
    return this.bodyData
}

// Get: options.
//
// @return (*ohttp.params.Params)
func (this *Message) Options() (*params.Params) {
    if this.options == nil {
        this.options = params.New()
    }
    return this.options
}

// Get: error.
//
// @return (*ohttp.message.MessageError)
func (this *Message) Error() (*MessageError) {
    return this.error
}

// Set: type.
//
// @param  t uint
// @return (*ohttp.message.Message)
func (this *Message) SetType(t uint) (*Message) {
    this.type_ = t
    return this
}

// Set: protocol version.
//
// @param  pv string
// @return (*ohttp.message.Message)
func (this *Message) SetProtocolVersion(pv string) (*Message) {
    this.protocolVersion = pv
    return this
}

// Set: header.
//
// @param  k string
// @param  v string
// @return (*ohttp.message.Message)
func (this *Message) SetHeader(k, v string) (*Message) {
    this.headers.Set(k, v)
    return this
}

// Set: all headers.
//
// @param  kv interface{}
// @return (*ohttp.message.Message)
func (this *Message) SetHeaderAll(kv interface{}) (*Message) {
    if kv, _ := kv.(map[string]string); kv != nil {
        this.headers.SetAll(kv)
    }
    return this
}

// Set: error.
//
// @param  ec uint
// @param  et string
// @return (void)
func (this *Message) SetError(ec int, et string) {
    this.error.code = ec
    this.error.text = et
}

// Set: type.
//
// @param  b interface{}
// @return (*ohttp.message.Message)
func (this *Message) SetBody(b interface{}) (*Message) {
    if b == nil {
        return this
    }

    if this.type_ == TYPE_REQUEST {
        var c string
        ct := this.Header("Content-Type")
        switch b := b.(type) {
            case string:
                if util.StringSearch(ct, "application/json") {
                    c = util.Quote(b)
                } else {
                    c = util.String(b)
                }
            default:
                if util.StringSearch(ct, "application/json") {
                    b, err := util.JsonEncode(b)
                    if err != nil {
                        panic(err)
                    }
                    c = b
                } else {
                    c = util.String(b)
                }
        }
        cl := len(c)
        // @overwrite
        this.body = NewMessageBody(c, ct, cl)
        this.SetHeader("Content-Length", util.String(cl))
    } else if this.type_ == TYPE_RESPONSE {
        // @overwrite
        this.body = NewMessageBody(
            util.String(b),
            this.Header("Content-Type"),
            util.Int(this.Header("Content-Length")),
        )
    }

    return this
}

// Stringify.
//
// @param  sl string
// @return (void)
func (this *Message) ToString(sl string) (string) {
    s := sl
    for k, v := range this.HeaderAll() {
        if k == "0" { // response only
            continue
        }
        if (v != "") {
            s += util.StringFormat("%s: %s%s", k, v, util.CRLF)
        }
    }

    s += util.CRLF
    if this.body != nil {
        s += this.body.Content()
    }

    return s
}
