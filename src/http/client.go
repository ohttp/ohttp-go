package http

import (
    "http/request"
    "http/response"
)

type Client struct {
    request    *request.Request
    response   *response.Response
}

func NewClient() (*Client) {
    return &Client{
        request: NewRequest(),
        response: NewResponse(),
    }
}
