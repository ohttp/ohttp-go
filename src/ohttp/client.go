// Copyright (c) 2015-2016 Kerem Güneş
//   <http://qeremy.com>
//
// GNU General Public License v3.0
//   <http://www.gnu.org/licenses/gpl-3.0.txt>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
//
// @package ohttp
// @imports ohttp.util, ohttp.http.params, ohttp.headers,
//          ohttp.request, ohttp.request.method, ohttp.response
// @author  Kerem Güneş <qeremy[at]gmail[dot]com>
package ohttp

import (
    "ohttp/util"
    "ohttp/util/params"
    "ohttp/headers"
    "ohttp/request"
    "ohttp/request/method"
    "ohttp/response"
)

// @object ohttp.Client
type Client struct {
    request    *request.Request
    response   *response.Response
}

func Shutup() {}

// Constructor.
//
// @param  o ohttp.params.Params
// @return (*ohttp.Client)
func NewClient(o *params.Params) (*Client) {
    return &Client{
         request: NewRequest(o),
        response: NewResponse(o),
    }
}

// Get: Request object.
//
// @return (*ohttp.request.Request)
func (this *Client) Request() (*request.Request) {
    return this.request
}

// Get: Response object.
//
// @return (*ohttp.response.Response)
func (this *Client) Response() (*response.Response) {
    return this.response
}

// Perform a HTTP Request and return Response
//
// @param  u  string      Request URL
// @param  up interface{} Request URL parameters
// @param  b  interface{} Request body
// @param  h  interface{} Request headers
// @return (*ohttp.response.Response)
// @panics
func (this *Client) Do(u string, up, b, h interface{}) (*response.Response, error) {
    m, _, err := util.RegExpMatch(u, "^([A-Z]+)\\s+(.+)")
    if len(m) < 3 {
        panic("Usage: <Method GET|POST...> <Scheme http|https>://<Host>/<Path>... !")
    }

    // set request method, uri, body and all headers
    this.request.
        SetMethod(m[1]).
        SetUri(m[2], up).
        SetBody(b).
        SetHeaderAll(h)

    // send request
    rs, err := this.request.Send()
    if err != nil {
        return nil, err
    }

    // split headers/body parts
    rt := util.Explode(rs, util.CRLF + util.CRLF, 2)
    if len(rt) != 2 {
        return nil, util.Error("No valid response returned from server!", nil)
    }

    // parse headers
    rh := headers.Parse(rt[0])
    if _, ok := rh["0"]; ok {
        // set status-line
        this.response.SetStatus(rh["0"])
    }

    // set response body and all headers
    this.response.
        SetHeaderAll(rh).
        SetBody(rt[1])

    // check http error by response status code
    if sc := this.response.Status().Code(); sc >= 400 {
        st := this.response.Status().Text()
        this.response.SetError(sc, st)
    }

    return this.response, nil
}

// Perform a HTTP Request and return Response
//
// @param  u  string      Request URL
// @param  up interface{} Request URL parameters
// @param  b  interface{} Request body
// @param  h  interface{} Request headers
// @param  fn func        Callback function
// @return (*ohttp.response.Response)
// @panics
func (this *Client) DoFunc(u string, up, b, h interface{}, fn func (
    req *request.Request, res *response.Response, err error)) {
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
