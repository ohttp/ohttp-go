package http

import (
    "http/util"
    "http/util/params"
    "http/request"
    "http/response"
    "http/headers"
)

type Client struct {
    request    *request.Request
    response   *response.Response
}

func NewClient(o *params.Params) (*Client) {
    return &Client{
         request: NewRequest(o),
        response: NewResponse(o),
    }
}

func (this *Client) Request() (*request.Request) {
    return this.request
}
func (this *Client) Response() (*response.Response) {
    return this.response
}

func (this *Client) Do(u string, up interface{}, b interface{}, h interface{}) (*response.Response) {
    m, _, err := util.RegExpMatch(u, "^([A-Z]+)\\s+(.+)")
    if len(m) < 3 {
        panic("Usage: <Method GET|POST...> <Scheme http|https>://<Host>/<Path>... !")
    }

    this.request.SetMethod(m[1])
    this.request.SetUri(m[2], up)
    this.request.SetBody(b)
    this.request.SetHeaderAll(h)

    rs, err := this.request.Send()
    if err != nil {
        panic(err)
    }

    rt := util.Explode(rs, util.CRLF + util.CRLF, 2)
    if len(rt) != 2 {
        panic("No valid response returned from server!")
    }

    rh := headers.Parse(rt[0])
    if _, ok := rh["0"]; ok {
        this.response.SetStatus(rh["0"])
    }
    this.response.SetHeaderAll(rh)
    this.response.SetBody(rt[1])

    if sc := this.response.Status().Code(); sc >= 400 {
        st := this.response.Status().Text()
        this.response.SetError(sc, st)
    }

    return this.response
}
func (this *Client) DoFunc(u string, up interface{}, b interface{}, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Do(u, up, b, h)
    fn(this.request, this.response)
}

func (this *Client) Head() {}
func (this *Client) HeadFunc() {}

// ...
