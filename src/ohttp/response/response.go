package response

import (
    "ohttp/util"
    "ohttp/message"
    "ohttp/response/status"
)

type Response struct {
    message.Message // extends
    status          *status.Status
}

func Shutup() {}

func New(m message.Message) (*Response) {
    return &Response{
        Message: m,
        status: status.New(0, "", ""),
    }
}

func (this *Response) Status() (*status.Status) {
    return this.status
}

func (this *Response) SetStatus(s string) (*Response) {
    m, _, err := util.RegExpMatch(s, "^HTTP/\\d+\\.\\d+\\s+(\\d+)\\s+(.+)")
    if err != nil {
        panic(err)
    }
    if len(m) == 3 {
        this.status.SetStatus(s)
        this.status.SetCode(util.Int(m[1]))
        this.status.SetText(util.Trim(m[2], ""))
    }
    return this
}

func (this *Response) OK() (bool) {
    return (this.Error().Code() == 0 && this.Error().Text() == "")
}

func (this *Response) String() (string) {
    return this.ToString(this.TheStatusLine() + util.CRLF)
}

func (this *Response) TheStatusLine() (string) {
    return this.status.Status()
}
