package message

type MessageError struct {
    code   uint
    text   string
}

func NewMessageError(c uint, t string) (*MessageError) {
    return &MessageError{
        code: c,
        text: t,
    }
}

func (this *MessageError) Code() (uint) {
    return this.code
}

func (this *MessageError) Text() (string) {
    return this.text
}
