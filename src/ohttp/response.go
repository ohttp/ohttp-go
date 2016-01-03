package ohttp

import (
    "ohttp/message"
    "ohttp/response"
    "ohttp/util/params"
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
