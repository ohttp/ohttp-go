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
// @author  Kerem Güneş <qeremy[at]gmail[dot]com>
package message

// @object ohttp.method.MessageBodyData
type MessageBodyData struct {
    content       interface{}
    contentType   string
}

// Constructor.
//
// @param  c  string
// @param  ct string
// @return (*ohttp.message.MessageBodyData)
func NewMessageBodyData(c, ct string) (*MessageBodyData) {
    // @todo parse by contentType
    return &MessageBodyData{
            content: c,
        contentType: ct,
    }
}

// Get: content.
//
// @return (interface{})
func (this *MessageBodyData) Content() (interface{}) {
    return this.content
}

// Get: content type.
//
// @return (string)
func (this *MessageBodyData) ContentType() (string) {
    return this.contentType
}
