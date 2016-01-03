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
//
// @package ohttp
// @imports ohttp.http.params, ohttp.message, ohttp.request
// @author  Kerem Güneş <qeremy[at]gmail[dot]com>
package ohttp

import (
    "ohttp/util/params"
    "ohttp/message"
    "ohttp/request"
)

// Constructor.
//
// @param  o ohttp.params.Params
// @return (*ohttp.request.Request)
func NewRequest(o *params.Params) (*request.Request) {
    return request.New(
        *message.NewMessage(
            message.TYPE_REQUEST,
            message.PROTOCOL_VERSION_1_0,
            o,
        ),
    )
}
