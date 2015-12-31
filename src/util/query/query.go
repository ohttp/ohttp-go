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

/*func (this *Query) DataString() (string) {
    if this.dataString != "" {
        return this.dataString
    }

    for k, v := range this.data {
        vt := util.TypeReal(v)
        if vt == "[]int" {
            v = _fmt.Sprintf("[\"%s\"]", _str.Join(v.([]int), "\",\""))
        } else if vt == "[]string" {
            v = _fmt.Sprintf("[\"%s\"]", _str.Join(v.([]string), "\",\""))
        }
        this.dataString += _fmt.Sprintf(
            "%s=%s&", util.UrlEncode(k), util.UrlEncode(util.String(v)))
    }

    if this.dataString != "" {
        // drop last "&"
        this.dataString = this.dataString[0 : len(this.dataString) - 1]
        // purify some encoded stuff
        this.dataString = _str.NewReplacer(
            "%5B", "[",
            "%5D", "]",
            "%2C", ",",
        ).Replace(this.dataString)
    }

    return this.dataString
}

*/
