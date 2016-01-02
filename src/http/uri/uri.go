package uri

import (
    _url "net/url"
    _str "strings"
)

import (
    "http/util"
    "http/util/query"
    "http/util/array/sarray"
)

type Uri struct {
    source     string
    scheme     string
    host       string
    port       uint
    path       string
    username   string
    password   string
    query      *query.Query
    fragment   string
    segments   []string
}

func New(s string, q interface{}) (*Uri) {
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
                this.segments = sarray.Filter(this.segments, nil)
            }
        }
        if s.User != nil {
            this.username = s.User.Username()
            this.password, _ = s.User.Password()
        }
        if sq := s.RawQuery; sq != "" {
            this.query = query.New(sq)
        } else {
            this.query = query.New(q)
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
func (this *Uri) Query() (*query.Query) {
    return this.query
}
func (this *Uri) Fragment() (string) {
    return this.fragment
}
func (this *Uri) Segments() ([]string) {
    return this.segments
}
func (this *Uri) Segment(i int) (string) {
    if se, ok := sarray.FindIndex(this.segments, i); ok {
        return se
    }
    return ""
}
func (this *Uri) Authorization() (string) {
    if this.username != "" && this.password != "" {
        return this.username +":"+ this.password
    } else if this.username != "" {
        return this.username
    }
    return ""
}
