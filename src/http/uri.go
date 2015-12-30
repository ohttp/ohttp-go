package http

import (
    _url "net/url"
    _str "strings"
)

import (
    "util"
)

type Uri struct {
    source   interface{}
    scheme   string
    host     string
    port     uint
    path     string
    username string
    password string
    query    map[string]string
    fragment string
    segments []string // @todo
}

func NewUri(source string) (*Uri) {
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
            }
        }
        if username := source.User.Username(); username != "" {
            this.username = username
        }
        if password, _ := source.User.Password(); password != "" {
            this.password = password
        }
        if query := source.RawQuery; query != "" {
            this.query = util.ParseQuery(query)
        }
        if fragment := source.Fragment; fragment != "" {
            this.fragment = fragment
        }
    }

    return this
}
