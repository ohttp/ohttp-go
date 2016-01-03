package message

import (
    "ohttp/util"
    "ohttp/util/params"
    "ohttp/headers"
)

type Message struct {
    type_           uint
    protocolVersion string
    headers         *headers.Headers
    body            *MessageBody
    bodyData        *MessageBodyData // parsed
    options         *params.Params
    error           *MessageError
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
    TYPE_REQUEST  uint = 1
    TYPE_RESPONSE      = 2
)

const (
    PROTOCOL_VERSION_1_0 = "1.0"
    PROTOCOL_VERSION_1_1 = "1.1"
    PROTOCOL_VERSION_2_0 = "2.0"
)

func Shutup() {}

func NewMessage(t uint, pv string, o *params.Params) (*Message) {
    return &Message{
        type_: t,
        protocolVersion: pv,
        headers: headers.New(),
        body: NewMessageBody("", "", 0),
        bodyData: NewMessageBodyData("", ""),
        error: NewMessageError(0, ""),
        options: o,
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
func (this *Message) BodyData() (*MessageBodyData) {
    return this.bodyData
}
func (this *Message) Options() (*params.Params) {
    if this.options == nil {
        this.options = params.New()
    }
    return this.options
}
func (this *Message) Error() (*MessageError) {
    return this.error
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

func (this *Message) SetHeaderAll(kv interface{}) (*Message) {
    if kv, _ := kv.(map[string]string); kv != nil {
        this.headers.SetAll(kv)
    }
    return this
}

func (this *Message) SetError(ec int, et string) {
    this.error.code = ec
    this.error.text = et
}

func (this *Message) SetBody(b interface{}) (*Message) {
    if b == nil {
        return this
    }

    if this.type_ == TYPE_REQUEST {
        var c string
        ct := this.Header("Content-Type")
        switch b := b.(type) {
            case string:
                if util.StringSearch(ct, "application/json") {
                    c = util.Quote(b)
                } else {
                    c = util.String(b)
                }
            default:
                if util.StringSearch(ct, "application/json") {
                    b, err := util.JsonEncode(b)
                    if err != nil {
                        panic(err)
                    }
                    c = b
                } else {
                    c = util.String(b)
                }
        }
        cl := len(c)
        // @overwrite
        this.body = NewMessageBody(c, ct, cl)
        this.SetHeader("Content-Length", util.String(cl))
    } else if this.type_ == TYPE_RESPONSE {
        // @overwrite
        this.body = NewMessageBody(
            util.String(b),
            this.Header("Content-Type"),
            util.Int(this.Header("Content-Length")),
        )
    }

    return this
}

func (this *Message) ToString(sl string) (string) {
    s := sl
    for k, v := range this.HeaderAll() {
        if k == "0" { // response only
            continue
        }
        if (v != "") {
            s += util.StringFormat("%s: %s%s", k, v, util.CRLF)
        }
    }
    s += util.CRLF
    if this.body != nil {
        s += this.body.Content()
    }
    return s
}
