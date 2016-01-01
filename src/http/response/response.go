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
        this.SetStatusCode(c)
        this.SetStatusText(t)
    }
    return this
}
func (this *Response) SetStatusCode(c int) (*Response) {
    this.status.SetCode(c)
    return this
}
func (this *Response) SetStatusText(t string) (*Response) {
    this.status.SetText(t)
    return this
}


func (this *Response) GetStatus() (*status.Status) {
    return this.status
}
func (this *Response) GetStatusCode() (int) {
    return this.status.GetCode()
}
func (this *Response) GetStatusText() (string) {
    return this.status.GetText()
}
