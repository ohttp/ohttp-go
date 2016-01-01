package request

import (
    _fmt "fmt"
    _net "net"
    _bio "bufio"
)

import (
    "util"
    "http/message"
    "http/uri"
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

func (this *Request) SetMethod(m string) (*Request) {
    this.method = util.Upper(m)
    return this
}
func (this *Request) SetUri(u string, up interface{}) (*Request) {
    this.uri = uri.New(u, up)
    return this
}

func (this *Request) GetMethod() (string) {
    return this.method
}
func (this *Request) GetUri() (*uri.Uri) {
    return this.uri
}

func (this *Request) Send() (string, error) {
    this.SetHeader("Host", this.uri.Host())
    this.SetHeader("Connection", "close")

    uh := this.uri.Host()
    if s := this.uri.Scheme(); s != "" {
        if s == "http" {
            // dial via http
        } else if s == "https" {
            // dial via https
        } else {
            // ...
        }
    }

    // link, err := _net.Dial("tcp", uh)
    // if err != nil {
    //     panic(err)
    // }
    defer link.Close()

    us := "/"
    if s := this.uri.Path(); s != "" {
        us = s
    }
    if s := this.uri.Query().String(); s != "" {
        us += "?"+ s
    }

    var rs, rr string
    rs += _fmt.Sprintf("%s %s HTTP/%s\r\n", this.method, us, this.GetProtocolVersion())
    for k, v := range this.GetHeaderAll() {
        if v != "" {
            rs += _fmt.Sprintf("%s: %s\r\n", k, v)
        }
    }
    rs += "\r\n"
    rs += this.GetBody().Content()
    util.Dumps(rs)

    _fmt.Fprint(link, rs)

    r := _bio.NewReader(link)
    // status-line
    s, err := r.ReadString('\n')
    if s == "" {
        print("HTTP error: no response returned from server!\n")
        print("---------------------------------------------\n")
        print(rs)
        print("---------------------------------------------\n")
        panic(err)
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

    return rr, nil
}
