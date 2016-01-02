package http

import (
    "http/util/params"
    "http/message"
    "http/response"
)

func NewResponse(o *params.Params) (*response.Response) {
    return response.New(
        message.NewMessage(
            message.TYPE_RESPONSE,
            message.PROTOCOL_VERSION_1_0,
            o,
        ),
    )
}
