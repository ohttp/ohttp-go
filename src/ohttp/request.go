package ohttp

import (
    "ohttp/message"
    "ohttp/request"
    "ohttp/util/params"
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
