package http

type Client struct {
    request    *Request
    response   *Response
}

func NewClient() (*Client) {
    return &Client{
        request: NewRequest(),
        response: NewResponse(),
    }
}
