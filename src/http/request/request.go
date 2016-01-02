package request

import (
    _fmt "fmt"
    _bio "bufio"
)

import (
    "util"
    "http/message"
    "http/uri"
    "http/useragent"
    "http/request/scheme"
)

type Request struct {
    message.Message // extends
    method          string
    uri             *uri.Uri
}

func New(m *message.Message) (*Request) {
    this := &Request{
        Message: *m,
    }
    return this
}

func (this *Request) Method() (string) {
    return this.method
}
func (this *Request) Uri() (*uri.Uri) {
    return this.uri
}

func (this *Request) SetMethod(m string) (*Request) {
    this.method = util.Upper(m)
    return this
}
func (this *Request) SetUri(u string, up interface{}) (*Request) {
    this.uri = uri.New(u, up)
    return this
}

func (this *Request) Send() (string, error) {
    debug := this.Options().GetBool("debug")

    this.SetHeader("Host", this.uri.Host())
    this.SetHeader("Connection", "close")
    if this.Header("User-Agent") == "" {
        this.SetHeader("User-Agent", _fmt.Sprintf("%s/v%s (+%s)",
            useragent.OH_NAME, useragent.OH_VERSION, useragent.OH_LINK))
    }

    var uh, us string
    uh = this.uri.Host()
    us = this.uri.Scheme()
    if this.uri.Port() == scheme.HTTPS_PORT {
        us = scheme.HTTPS
    }

    link, err := scheme.Dial(uh, us)
    if err != nil {
        return "", err
    }
    defer link.Close()

    var rs, rr string
    rs += this.RequestLine()
    for k, v := range this.HeaderAll() {
        if v != "" {
            rs += _fmt.Sprintf("%s: %s\r\n", k, v)
        }
    }
    rs += "\r\n"
    rs += this.Body().Content()

    _fmt.Fprint(link, rs)

    r := _bio.NewReader(link)
    // status-line
    s, err := r.ReadString('\n')
    if s == "" {
        return "", err
    }
    rr += s

    for {
        b := make([]byte, 1024)
        if bl, _ := r.Read(b); bl == 0 {
            break // eof
        }
        rr += util.Trim(string(b), "\x00")
    }

    link.Close()

    if debug == true {
        util.Dumps(rs)
        util.Dumps(rr)
    }

    return rr, nil
}

func (this *Request) String() (string) {
    return this.ToString(this.RequestLine())
}


func (this *Request) RequestLine() (string) {
    rm  := this.Method()
    rp  := "/"
    rpv := this.ProtocolVersion()
    if s := this.uri.Path(); s != "" {
        rp = s
    }
    if s := this.uri.Query().String(); s != "" {
        rp += "?"+ s
    }
    return _fmt.Sprintf("%s %s HTTP/%s\r\n", rm, rp, rpv)
}
