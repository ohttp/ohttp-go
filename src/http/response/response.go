package response

import (
    "util"
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

func (this *Response) SetStatus(s string) (*Response) {
    m, _, err := util.RegExpMatch(s, "^HTTP/\\d+\\.\\d+\\s+(\\d+)\\s+(.+)")
    if err != nil {
        panic(err)
    }
    if len(m) == 3 {
        this.status.SetTextPhrase(s)
        c, t := util.Int(m[1]), util.Trim(m[2], "")
        this.status.SetCode(c)
        this.status.SetText(t)
    }
    return this
}

func (this *Response) GetStatus() (*status.Status) {
    return this.status
}
