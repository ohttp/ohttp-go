package http

import (
    _str "strings"
)

import (
    "util"
)

type Request struct {
    Stream
    method string
    uri    *Uri
}

func NewRequest() (*Request) {
    return &Request{
        Stream: *NewStream(TYPE_REQUEST, HTTP_VERSION_1_0),
    }
}

func (this *Request) SetMethod(method string) (*Request) {
    this.method = _str.ToUpper(method)
    return this
}
func (this *Request) GetMethod() (string) {
    return this.method
}

func (this *Request) SetUri(uri string, uriParams map[string]string) (*Request) {
    if uriParams != nil {
        var query string
        for key, value := range uriParams {
            query += util.StringFormat("%s=%s", key, util.UrlEncode(value))
        }
        uri += "?"+ query
    }
    this.uri = NewUri(uri)
    return this
}
func (this *Request) GetUri() (*Uri) {
    return this.uri
}
