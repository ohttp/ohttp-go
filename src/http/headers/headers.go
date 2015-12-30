package headers

type Headers struct {
    data map[string]string
}

func (this *Headers) Set(key, value string) (*Headers) {
    this.data[key] = value
    return this
}
func (this *Headers) Get(key string) (string) {
    if value, ok := this.data[key]; ok {
        return value
    }
    return ""
}
