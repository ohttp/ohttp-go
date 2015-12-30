package http

import (
    "util"
    "http/headers"
)

type Stream struct {
    type_       string
    httpVersion string
    headers     *headers.Headers
    body        string
    StreamBody
    StreamString
}

type StreamBody interface {
    SetBody(body interface{})
}

type StreamString interface {
    String() (string)
}

const (
    TYPE_REQUEST  = "request"
    TYPE_RESPONSE = "response"
)

const (
    HTTP_VERSION_1_0 = "1.0"
    HTTP_VERSION_1_1 = "1.1"
    HTTP_VERSION_2_0 = "2.0"
)

func Shutup() {
    util.Shutup()
}

func NewStream(type_, httpVersion string) (*Stream) {
    return &Stream{
              type_: type_,
        httpVersion: httpVersion,
            headers: headers.New(),
               body: "",
    }
}

func (this *Stream) SetType(type_ string) {
    this.type_ = type_
}
func (this *Stream) GetType() (string) {
    return this.type_
}

func (this *Stream) SetHttpVersion(httpVersion string) {
    this.httpVersion = httpVersion
}
func (this *Stream) GetHttpVersion() (string) {
    return this.httpVersion
}

func (this *Stream) SetHeader(name, value string) {
    this.headers.Set(name, value)
}
func (this *Stream) GetHeader(name string) (string) {
    return this.headers.Get(name)
}

func (this *Stream) SetHeaderAll(data map[string]string) {
    this.headers.SetAll(data)
}
func (this *Stream) GetHeaderAll() (map[string]string) {
    return this.headers.GetAll(nil)
}

func (this *Stream) GetBody() (string) {
    return this.body
}
