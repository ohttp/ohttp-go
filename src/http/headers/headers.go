package headers

import (
    "http/util"
    "http/util/array/sarray"
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

func Parse(hs string) (map[string]string) {
    ret := util.MapString()
    if tmp := util.Explode(hs, util.CRLF, -1); tmp != nil {
        // status line (HTTP/1.0 200 OK)
        ret["0"] = sarray.Shift(&tmp)

        for _, tm := range tmp {
            if t := util.Explode(tm, ":", 2); len(t) == 2 {
                ret[util.Trim(t[0], "")] =  util.Trim(t[1], "")
            }
        }
    }

    return ret
}
