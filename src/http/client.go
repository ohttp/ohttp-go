package http

import (
    "http/headers"
    "http/request"
    "http/request/method"
    "http/response"
    "http/util"
    "http/util/params"
)

type Client struct {
    request    *request.Request
    response   *response.Response
}

func Shutup() {}

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

func (this *Client) Do(u string, up, b, h interface{}) (*response.Response) {
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
func (this *Client) DoFunc(u string, up, b, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Do(u, up, b, h)
    fn(this.request, this.response)
}

func (this *Client) Options(u string, up, h interface{}) (*response.Response) {
    return this.Do(method.OPTIONS +" "+ u, up, nil, h)
}
func (this *Client) OptionsFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Options(u, up, h)
    fn(this.request, this.response)
}

func (this *Client) Head(u string, up, h interface{}) (*response.Response) {
    return this.Do(method.HEAD +" "+ u, up, nil, h)
}
func (this *Client) HeadFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Head(u, up, h)
    fn(this.request, this.response)
}

func (this *Client) Get(u string, up, h interface{}) (*response.Response) {
    return this.Do(method.GET +" "+ u, up, nil, h)
}
func (this *Client) GetFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Get(u, up, h)
    fn(this.request, this.response)
}

func (this *Client) Post(u string, up, b, h interface{}) (*response.Response) {
    return this.Do(method.POST +" "+ u, up, b, h)
}
func (this *Client) PostFunc(u string, up, b, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Post(u, up, b, h)
    fn(this.request, this.response)
}

func (this *Client) Put(u string, up, b, h interface{}) (*response.Response) {
    return this.Do(method.PUT +" "+ u, up, b, h)
}
func (this *Client) PutFunc(u string, up, b, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Put(u, up, b, h)
    fn(this.request, this.response)
}

func (this *Client) Patch(u string, up, b, h interface{}) (*response.Response) {
    return this.Do(method.PATCH +" "+ u, up, b, h)
}
func (this *Client) PatchFunc(u string, up, b, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Patch(u, up, b, h)
    fn(this.request, this.response)
}

func (this *Client) Delete(u string, up, h interface{}) (*response.Response) {
    return this.Do(method.DELETE +" "+ u, up, nil, h)
}
func (this *Client) DeleteFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Delete(u, up, h)
    fn(this.request, this.response)
}

func (this *Client) Trace(u string, up, h interface{}) (*response.Response) {
    return this.Do(method.TRACE +" "+ u, up, nil, h)
}
func (this *Client) TraceFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Trace(u, up, h)
    fn(this.request, this.response)
}

func (this *Client) Connect(u string, up, h interface{}) (*response.Response) {
    return this.Do(method.CONNECT +" "+ u, up, nil, h)
}
func (this *Client) ConnectFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Connect(u, up, h)
    fn(this.request, this.response)
}

func (this *Client) Copy(u string, up, h interface{}) (*response.Response) {
    return this.Do(method.COPY +" "+ u, up, nil, h)
}
func (this *Client) CopyFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Copy(u, up, h)
    fn(this.request, this.response)
}

func (this *Client) Move(u string, up, h interface{}) (*response.Response) {
    return this.Do(method.MOVE +" "+ u, up, nil, h)
}
func (this *Client) MoveFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response)) {
    this.Move(u, up, h)
    fn(this.request, this.response)
}
