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

// @object ohttp.message.MessageError
type MessageError struct {
    code   int
    text   string
}

// Constructor.
//
// @param  c int
// @param  t string
// @return (*ohttp.message.MessageError)
func NewMessageError(c int, t string) (*MessageError) {
    return &MessageError{
        code: c,
        text: t,
    }
}

// Get: code
//
// @return (int)
func (this *MessageError) Code() (int) {
    return this.code
}

// Get: text
//
// @return (string)
func (this *MessageError) Text() (string) {
    return this.text
}
