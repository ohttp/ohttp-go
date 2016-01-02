package message

type MessageError struct {
    code   int
    text   string
}

func NewMessageError(c int, t string) (*MessageError) {
    return &MessageError{
        code: c,
        text: t,
    }
}

func (this *MessageError) Code() (int) {
    return this.code
}

func (this *MessageError) Text() (string) {
    return this.text
}
