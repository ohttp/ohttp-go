package message

type MessageBodyData struct {
    content   interface{}
}

func NewMessageBodyData(c string) (*MessageBodyData) {
    return &MessageBodyData{
        content: c,
    }
}

func (this *MessageBodyData) Content() (interface{}) {
    return this.content
}
