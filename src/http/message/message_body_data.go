package message

type MessageBodyData struct {
    content       interface{}
    contentType   string
}

func NewMessageBodyData(c, ct string) (*MessageBodyData) {
    // @todo parse by contentType
    return &MessageBodyData{
            content: c,
        contentType: ct,
    }
}

func (this *MessageBodyData) Content() (interface{}) {
    return this.content
}
