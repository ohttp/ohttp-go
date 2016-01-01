package query

import (
    // _fmt "fmt"
    // _str "strings"
)

import (
    // "util"
    "util/params"
)

type Query struct {
    params        *params.Params
    paramsString  string
}

func Shutup() {}

func New(p interface{}) (*Query) {
    this := &Query{}
    this.params = params.New()
    if _, ok := p.(*params.Params); ok {
        for k, v := range (*p.(*params.Params)) {
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
