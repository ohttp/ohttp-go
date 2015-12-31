package message

import (
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

type MessageBody struct {
    content         string
}
type MessageBodyData struct {
    content         interface{}
}

type MessageBodyInterface interface {
    SetBody(body string)
    SetBodyData(body string, i interface{}) interface{}
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

func NewMessage(type_ uint, protocolVersion string) (*Message) {
    return &Message{
                  type_: type_,
        protocolVersion: protocolVersion,
                headers: headers.New(),
    }
}

func (this *Message) SetType(type_ uint) (*Message) {
    this.type_ = type_
    return this
}
func (this *Message) GetType() (uint) {
    return this.type_
}

func (this *Message) SetProtocolVersion(protocolVersion string) (*Message) {
    this.protocolVersion = protocolVersion
    return this
}
func (this *Message) GetProtocolVersion() (string) {
    return this.protocolVersion
}

func (this *Message) SetHeader(name, value string) (*Message) {
    this.headers.Set(name, value)
    return this
}
func (this *Message) GetHeader(name string) (string) {
    return this.headers.Get(name)
}

func (this *Message) SetHeaderAll(data map[string]string) (*Message) {
    this.headers.SetAll(data)
    return this
}
func (this *Message) GetHeaderAll() (map[string]string) {
    return this.headers.GetAll(nil)
}
