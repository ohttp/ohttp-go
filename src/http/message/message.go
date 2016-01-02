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
    MessageOKInterface
    MessageStringInterface
}

type MessageOKInterface interface {
    OK() (bool)
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
        body: NewMessageBody("", ""),
        bodyData: NewMessageBodyData(""),
    }
}

func (this *Message) Type() (uint) {
    return this.type_
}
func (this *Message) ProtocolVersion() (string) {
    return this.protocolVersion
}
func (this *Message) Header(k string) (string) {
    return this.headers.Get(k)
}
func (this *Message) HeaderAll() (map[string]string) {
    return this.headers.GetAll()
}
func (this *Message) Body() (*MessageBody) {
    return this.body
}

func (this *Message) SetType(t uint) (*Message) {
    this.type_ = t
    return this
}

func (this *Message) SetProtocolVersion(pv string) (*Message) {
    this.protocolVersion = pv
    return this
}

func (this *Message) SetHeader(k, v string) (*Message) {
    this.headers.Set(k, v)
    return this
}

func (this *Message) SetHeaderAll(kv map[string]string) (*Message) {
    this.headers.SetAll(kv)
    return this
}

func (this *Message) SetBody(b interface{}) {
    if b != nil {
        var c, ct string
        ct = this.Header("Content-Type")
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
        // @overwrite
        this.body = NewMessageBody(c, ct)
        this.SetHeader("Content-Length", util.String(this.body.ContentLength()))
    }
}
