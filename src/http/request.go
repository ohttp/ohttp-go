package http

import (
    "http/message"
    "http/request"
)

func NewRequest() (*request.Request) {
    return request.New(
        message.NewMessage(
            message.TYPE_REQUEST,
            message.PROTOCOL_VERSION_1_0,
        ),
    )
}
