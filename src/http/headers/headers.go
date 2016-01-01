package headers

import (
    "util"
)

type Headers struct {
    data map[string]string
}

func New() (*Headers) {
    return &Headers{
        data: util.MapString(),
    }
}

func (this *Headers) Set(k, v string) (*Headers) {
    this.data[k] = util.Trim(v, "")
    return this
}
func (this *Headers) SetAll(data map[string]string) {
    this.data = data
}

func (this *Headers) Get(k string) (string) {
    if v, ok := this.data[k]; ok {
        return v
    }
    return ""
}
func (this *Headers) GetAll() (map[string]string) {
    return this.data
}
