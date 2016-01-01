package scheme

import (
    _net "net"
    _tls "crypto/tls"
)

const (
    HTTP            = "http"
    HTTPS           = "https"
    HTTP_PORT  uint = 80
    HTTPS_PORT uint = 443
)

const (
)

func Dial(h, s string) (_net.Conn, error) {
    if s == HTTP {
        return DialHttp(h)
    } else if s == HTTPS {
        return DialHttps(h)
    }
    panic("Unsupported scheme given!")
}

func DialHttp(h string) (_net.Conn, error) {
    return _net.Dial("tcp", h +":"+ HTTP)
}

func DialHttps(h string) (_net.Conn, error) {
    return _tls.Dial("tcp", h +":"+ HTTPS, nil)
}
