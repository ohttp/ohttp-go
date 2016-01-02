package http

import (
    "http/message"
    "http/request"
    "http/util/params"
)

func NewRequest(o *params.Params) (*request.Request) {
    return request.New(
        *message.NewMessage(
            message.TYPE_REQUEST,
            message.PROTOCOL_VERSION_1_0,
            o,
        ),
    )
}
