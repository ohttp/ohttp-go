package request

import (
    _fmt "fmt"
    _bio "bufio"
)

import (
    "http/util"
    "http/message"
    "http/uri"
    "http/useragent"
    "http/connection"
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

    link, err := connection.Dial(this.uri.Host(), this.uri.Port())
    if err != nil {
        return "", err
    }
    defer link.Close()

    var rs, rr string
    rs += this.TheRequestLine()
    for k, v := range this.HeaderAll() {
        if v != "" {
            rs += _fmt.Sprintf("%s: %s%s", k, v, util.CRLF)
        }
    }
    rs += util.CRLF
    rs += this.Body().Content()

    _fmt.Fprint(link, rs)

    r := _bio.NewReader(link)
    // status-line
    sl, err := r.ReadString('\n')
    if sl == "" {
        return "", err
    }
    rr += sl

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
    return this.ToString(this.TheRequestLine())
}


func (this *Request) TheRequestLine() (string) {
    rm  := this.Method()
    rp  := "/"
    rpv := this.ProtocolVersion()
    if s := this.uri.Path(); s != "" {
        rp = s
    }
    if s := this.uri.Query().String(); s != "" {
        rp += "?"+ s
    }
    return _fmt.Sprintf("%s %s HTTP/%s%s", rm, rp, rpv, util.CRLF)
}
