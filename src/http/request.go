package http

type Request struct {
    Stream
    method string
    uri    *Uri
}

func NewRequest() (*Request) {
    this := &Request{}
    this.Stream = *NewStream(TYPE_REQUEST, HTTP_VERSION_1_0)
    return this
}
