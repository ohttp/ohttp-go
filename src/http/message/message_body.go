package message

type MessageBody struct {
    content       string
    contentType   string
    contentLength int
}

func NewMessageBody(c, ct string) (*MessageBody) {
    return &MessageBody{
        content: c,
        contentType: ct,
        contentLength: len(c),
    }
}

func (this *MessageBody) Content() (string) {
    return this.content
}

func (this *MessageBody) ContentType() (string) {
    return this.contentType
}

func (this *MessageBody) ContentLength() (int) {
    return this.contentLength
}
