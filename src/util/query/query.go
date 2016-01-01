package query

import (
    "util/params"
)

type Query struct {
    params   *params.Params
}

func Shutup() {}

func New(p interface{}) (*Query) {
    this := &Query{}
    this.params = params.New()
    if _, ok := p.(params.Params); ok {
        for k, v := range p.(params.Params) {
            this.params.Set(k, v)
        }
    } else if _, ok := p.(*params.Params); ok {
        for k, v := range *p.(*params.Params) {
            this.params.Set(k, v)
        }
    } else if _, ok := p.(map[string]interface{}); ok {
        for k, v := range p.(map[string]interface{}) {
            this.params.Set(k, v)
        }
    }
    return this
}

func (this *Query) Set(k string, v interface{}) (*Query) {
    this.params.Set(k, v)
    return this
}

func (this *Query) Get(k string) (interface{}) {
    return this.params.Get(k)
}

func (this *Query) Params() (*params.Params) {
    return this.params
}

func (this *Query) String() (string) {
    return this.params.String()
}
