package request

import (
    "util"
    "http/uri"
    "http/message"
)

type Request struct {
    message.Message // extends
    method          string
    uri             *uri.Uri
}

func New(m *message.Message) (*Request) {
    return &Request{
        Message: *m,
    }
}

func (this *Request) SetMethod(m string) (*Request) {
    this.method = util.Upper(m)
    return this
}
func (this *Request) GetMethod() (string) {
    return this.method
}

func (this *Request) SetUri(u string, up map[string]string) (*Request) {
    if ups := util.UrlQueryUnparse(up); ups != "" {
        this.uri = uri.New(u +"?"+ ups)
    } else {
        this.uri = uri.New(u)
    }
    return this
}
func (this *Request) GetUri() (*uri.Uri) {
    return this.uri
}
