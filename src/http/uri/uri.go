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

func New(source string) (*Uri) {
    this := &Uri{
        source: source,
    }

    if source != "" {
        source, _ := _url.Parse(util.UrlDecode(source))
        if scheme := source.Scheme; scheme != "" {
            this.scheme = scheme
        }
        if host := source.Host; host != "" {
            this.host = host
            if tmp := _str.Split(host, ":"); len(tmp) == 2 {
                this.host = tmp[0]
                this.port = util.UInt(tmp[1])
            }
        }
        if path := source.Path; path != "" {
            this.path = path
            if segments := _str.Split(path, "/"); len(segments) > 0 {
                this.segments = util.MapStringSlice(segments)
                for i, segment := range segments {
                    segment = _str.TrimSpace(segment)
                    if segment != "" {
                        this.segments[i] = segment
                    }
                }
                this.segments = scol.Filter(this.segments, nil)
            }
        }
        if source.User != nil {
            if username := source.User.Username(); username != "" {
                this.username = username
            }
            if password, _ := source.User.Password(); password != "" {
                this.password = password
            }
        }
        if query := source.RawQuery; query != "" {
            this.query = query
            this.queryParams = util.ParseQuery(query)
        }
        if fragment := source.Fragment; fragment != "" {
            this.fragment = fragment
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
func (this *Uri) QueryParam(key string) (string) {
    if param, ok := this.queryParams[key]; ok {
        return param
    }
    return ""
}
func (this *Uri) QueryParams() (map[string]string) {
    return this.queryParams
}
func (this *Uri) Fragment() (string) {
    return this.fragment
}
func (this *Uri) Segment(i int) (string) {
    if segment, ok := scol.FindIndex(this.segments, i); ok {
        return segment
    }
    return ""
}
func (this *Uri) Segments() ([]string) {
    return this.segments
}
