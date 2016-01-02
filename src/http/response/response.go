package response

import (
    "http/util"
    "http/message"
    "http/response/status"
)

type Response struct {
    message.Message // extends
    status          *status.Status
}

func New(m *message.Message) (*Response) {
    this := &Response{
        Message: *m,
        status: status.New(0, "", ""),
    }
    return this
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

func (this *Response) String() (string) {
    return this.ToString(this.TheStatusLine() + util.CRLF)
}

func (this *Response) TheStatusLine() (string) {
    return this.status.Status()
}
