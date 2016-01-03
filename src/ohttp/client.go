package ohttp

import (
    "ohttp/headers"
    "ohttp/request"
    "ohttp/request/method"
    "ohttp/response"
    "ohttp/util"
    "ohttp/util/params"
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

func (this *Client) Do(u string, up, b, h interface{}) (*response.Response, error) {
    m, _, err := util.RegExpMatch(u, "^([A-Z]+)\\s+(.+)")
    if len(m) < 3 {
        panic("Usage: <Method GET|POST...> <Scheme http|https>://<Host>/<Path>... !")
    }

    this.request.
        SetMethod(m[1]).
        SetUri(m[2], up).
        SetBody(b).
        SetHeaderAll(h)

    rs, err := this.request.Send()
    if err != nil {
        return nil, err
    }

    rt := util.Explode(rs, util.CRLF + util.CRLF, 2)
    if len(rt) != 2 {
        return nil, util.Error("No valid response returned from server!", nil)
    }

    rh := headers.Parse(rt[0])
    if _, ok := rh["0"]; ok {
        this.response.SetStatus(rh["0"])
    }
    this.response.
        SetHeaderAll(rh).
        SetBody(rt[1])

    if sc := this.response.Status().Code(); sc >= 400 {
        st := this.response.Status().Text()
        this.response.SetError(sc, st)
    }

    return this.response, nil
}
func (this *Client) DoFunc(u string, up, b, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Do(u, up, b, h)
    fn(this.request, this.response, err)
}

func (this *Client) Options(u string, up, h interface{}) (*response.Response, error) {
    return this.Do(method.OPTIONS +" "+ u, up, nil, h)
}
func (this *Client) OptionsFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Options(u, up, h)
    fn(this.request, this.response, err)
}

func (this *Client) Head(u string, up, h interface{}) (*response.Response, error) {
    return this.Do(method.HEAD +" "+ u, up, nil, h)
}
func (this *Client) HeadFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Head(u, up, h)
    fn(this.request, this.response, err)
}

func (this *Client) Get(u string, up, h interface{}) (*response.Response, error) {
    return this.Do(method.GET +" "+ u, up, nil, h)
}
func (this *Client) GetFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Get(u, up, h)
    fn(this.request, this.response, err)
}

func (this *Client) Post(u string, up, b, h interface{}) (*response.Response, error) {
    return this.Do(method.POST +" "+ u, up, b, h)
}
func (this *Client) PostFunc(u string, up, b, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Post(u, up, b, h)
    fn(this.request, this.response, err)
}

func (this *Client) Put(u string, up, b, h interface{}) (*response.Response, error) {
    return this.Do(method.PUT +" "+ u, up, b, h)
}
func (this *Client) PutFunc(u string, up, b, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Put(u, up, b, h)
    fn(this.request, this.response, err)
}

func (this *Client) Patch(u string, up, b, h interface{}) (*response.Response, error) {
    return this.Do(method.PATCH +" "+ u, up, b, h)
}
func (this *Client) PatchFunc(u string, up, b, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Patch(u, up, b, h)
    fn(this.request, this.response, err)
}

func (this *Client) Delete(u string, up, h interface{}) (*response.Response, error) {
    return this.Do(method.DELETE +" "+ u, up, nil, h)
}
func (this *Client) DeleteFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Delete(u, up, h)
    fn(this.request, this.response, err)
}

func (this *Client) Trace(u string, up, h interface{}) (*response.Response, error) {
    return this.Do(method.TRACE +" "+ u, up, nil, h)
}
func (this *Client) TraceFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Trace(u, up, h)
    fn(this.request, this.response, err)
}

func (this *Client) Connect(u string, up, h interface{}) (*response.Response, error) {
    return this.Do(method.CONNECT +" "+ u, up, nil, h)
}
func (this *Client) ConnectFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Connect(u, up, h)
    fn(this.request, this.response, err)
}

func (this *Client) Copy(u string, up, h interface{}) (*response.Response, error) {
    return this.Do(method.COPY +" "+ u, up, nil, h)
}
func (this *Client) CopyFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Copy(u, up, h)
    fn(this.request, this.response, err)
}

func (this *Client) Move(u string, up, h interface{}) (*response.Response, error) {
    return this.Do(method.MOVE +" "+ u, up, nil, h)
}
func (this *Client) MoveFunc(u string, up, h interface{},
    fn func (req *request.Request, res *response.Response, err error)) {
    _, err := this.Move(u, up, h)
    fn(this.request, this.response, err)
}
