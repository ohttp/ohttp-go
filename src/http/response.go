package http

import (
    "http/message"
    "http/response"
)

func NewResponse() (*response.Response) {
    return response.New(
        message.NewMessage(
            message.TYPE_RESPONSE,
            message.PROTOCOL_VERSION_1_0,
        ),
    )
}
