package scheme

import (
    _fmt "fmt"
    _net "net"
    _tls "crypto/tls"
)

const (
    PORT_HTTP  uint = 80
    PORT_HTTPS      = 443
)

func Dial(h string, p uint) (_net.Conn, error) {
    if p == PORT_HTTP {
        return DialHttp(h, PORT_HTTP)
    } else if p == PORT_HTTPS {
        return DialHttps(h, PORT_HTTPS)
    }
    return DialHttp(h, p)
}

func DialHttp(h string, p uint) (_net.Conn, error) {
    return _net.Dial("tcp", _fmt.Sprintf("%s:%v", h, p))
}

func DialHttps(h string, p uint) (_net.Conn, error) {
    return _tls.Dial("tcp", _fmt.Sprintf("%s:%v", h, p), nil)
}
