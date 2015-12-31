package uri

import (
    _url "net/url"
    _str "strings"
)

import (
    "util"
    "util/col/scol"
)

type Uri struct {
    source        string
    scheme        string
    host          string
    port          uint
    path          string
    username      string
    password      string
    query         string
    queryParams   map[string]string
    fragment      string
    segments      []string // @todo
}

func New(s string) (*Uri) {
    this := &Uri{
        source: s,
    }

    if s != "" {
        s, _ := _url.Parse(util.UrlDecode(s))
        if ss := s.Scheme; ss != "" {
            this.scheme = ss
        }
        if sh := s.Host; sh != "" {
            this.host = sh
            if tmp := _str.Split(sh, ":"); len(tmp) == 2 {
                this.host = tmp[0]
                this.port = util.UInt(tmp[1])
            }
        }
        if sp := s.Path; sp != "" {
            this.path = sp
            if seg := _str.Split(sp, "/"); len(seg) > 0 {
                this.segments = util.MapStringSlice(seg)
                for i, se := range seg {
                    se = _str.TrimSpace(se)
                    if se != "" {
                        this.segments[i] = se
                    }
                }
                this.segments = scol.Filter(this.segments, nil)
            }
        }
        if s.User != nil {
            if su := s.User.Username(); su != "" {
                this.username = su
            }
            if sp, _ := s.User.Password(); sp != "" {
                this.password = sp
            }
        }
        if sq := s.RawQuery; sq != "" {
            this.query = sq
            this.queryParams = util.UrlQueryParse(sq)
        }
        if sf := s.Fragment; sf != "" {
            this.fragment = sf
        }
    }

    return this
}

func (this *Uri) Source() (string) {
    return this.source
}
func (this *Uri) Scheme() (string) {
    return this.scheme
}
func (this *Uri) Host() (string) {
    return this.host
}
func (this *Uri) Port() (uint) {
    return this.port
}
func (this *Uri) Path() (string) {
    return this.path
}
func (this *Uri) Username() (string) {
    return this.username
}
func (this *Uri) Password() (string) {
    return this.password
}
func (this *Uri) Query() (string) {
    return this.query
}
func (this *Uri) QueryParam(k string) (string) {
    if v, ok := this.queryParams[k]; ok {
        return v
    }
    return ""
}
func (this *Uri) QueryParams() (map[string]string) {
    return this.queryParams
}
func (this *Uri) Fragment() (string) {
    return this.fragment
}
func (this *Uri) Segments() ([]string) {
    return this.segments
}
func (this *Uri) Segment(i int) (string) {
    if se, ok := scol.FindIndex(this.segments, i); ok {
        return se
    }
    return ""
}
