package message

import (
    "util"
    "http/headers"
)

type Message struct {
    type_           uint
    protocolVersion string
    headers         *headers.Headers
    body            *MessageBody
    bodyData        *MessageBodyData // parsed

    MessageBodyInterface
    MessageStringInterface
}

type MessageBodyData struct {
    content         interface{}
}

type MessageBodyInterface interface {
    // SetBody(b string)
    // SetBodyData(b string, i interface{}) interface{}
    String() (string)
}

type MessageStringInterface interface {
    String() (string)
}

const (
    CRLF = "\r\n"
)

const (
    TYPE_REQUEST  = 1
    TYPE_RESPONSE = 2
)

const (
    PROTOCOL_VERSION_1_0 = "1.0"
    PROTOCOL_VERSION_1_1 = "1.1"
    PROTOCOL_VERSION_2_0 = "2.0"
)

func Shutup() {}

func NewMessage(t uint, pv string) (*Message) {
    return &Message{
        type_: t,
        protocolVersion: pv,
        headers: headers.New(),
    }
}

func (this *Message) SetType(t uint) (*Message) {
    this.type_ = t
    return this
}
func (this *Message) GetType() (uint) {
    return this.type_
}

func (this *Message) SetProtocolVersion(pv string) (*Message) {
    this.protocolVersion = pv
    return this
}
func (this *Message) GetProtocolVersion() (string) {
    return this.protocolVersion
}

func (this *Message) SetHeader(k, v string) (*Message) {
    this.headers.Set(k, v)
    return this
}
func (this *Message) GetHeader(k string) (string) {
    return this.headers.Get(k)
}

func (this *Message) SetHeaderAll(kv map[string]string) (*Message) {
    this.headers.SetAll(kv)
    return this
}
func (this *Message) GetHeaderAll() (map[string]string) {
    return this.headers.GetAll()
}

func (this *Message) SetBody(b interface{}) {
    if b != nil {
        var c, ct string
        ct = this.GetHeader("Content-Type")
        switch b := b.(type) {
            case string:
                if ct == "application/json" {
                    c = util.Quote(b)
                }
            default:
                bt := util.StringFormat("%T", b)
                if util.StringSearch(bt, "^u?int(\\d+)?|float(32|64)$") {
                    c = util.String(b)
                } else {
                    if ct == "application/json" {
                        b, err := util.JsonEncode(b)
                        if err != nil {
                            panic(err)
                        }
                        c = b
                    }
                }
        }
        this.body = NewMessageBody(c, ct)

        this.SetHeader("Content-Length", util.String(this.body.ContentLength()))
    }
}
