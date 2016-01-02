package http

import (
    "http/message"
    "http/response"
    "http/util/params"
)

func NewResponse(o *params.Params) (*response.Response) {
    return response.New(
        *message.NewMessage(
            message.TYPE_RESPONSE,
            message.PROTOCOL_VERSION_1_0,
            o,
        ),
    )
}
