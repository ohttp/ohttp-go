package http

import (
    "http/util/params"
    "http/message"
    "http/request"
)

func NewRequest(o *params.Params) (*request.Request) {
    return request.New(
        message.NewMessage(
            message.TYPE_REQUEST,
            message.PROTOCOL_VERSION_1_0,
            o,
        ),
    )
}
