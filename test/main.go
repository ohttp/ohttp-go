package main

import (
    "http"
    "http/request"
    "http/response"
    "http/util"
    "http/util/params"
    // // "http/util/query"
)

func main() {
    // r1, _, e1 := util.RegExpMatch(rs, re)
    // // r2, _, e2 := util.RegExpMatchName(rs, re)
    // util.Dumpf("%#v %#v %d", r1, e1, len(r1))
    // // util.Dumpf("%#v %#v %d", r2["method"], e2, len(r2))

    o := params.New()
    // o.Set("debug", true)

    c := http.NewClient(o)

    // // r := c.Do("GET http://localhost/foo.json", nil, nil, nil)
    // // util.Dumpf("%#v", r)
    // // util.Dumpf("%+v", c.Request().String())
    // // util.Dumpf("%+v", c.Response().String())

    c.DoFunc("GET http://localhost/foo.json", nil, nil, nil,
        func(req *request.Request, res *response.Response) {
        util.Dump(req.String())
        util.Dump(res.String())
    })

    // c.Do("GET http://localhost/foo", nil, nil, nil)
    // util.Dumpf("%#v", c.Request())
    // util.Dumpf("%#v", c.Response())

    // q := query.New(map[string]interface{}{"a": true, "b": 1})
    // util.Dumpf("%#v", q.String())
    // util.Dumpf("%#v", q.Params().String())

    // req := http.NewRequest()
    // // req.SetMethod("GET")
    // req.SetUri("http://localhost/foo", nil)
    // // // util.Dumpf("%#v", req)
    // util.Dumpf("%#v", req.Uri())
    // // // util.Dumps("")

    // res, err := req.Send()
    // if err != nil {
    //     panic(err)
    // }
    // util.Dumps(res)

    // res := http.NewResponse()
    // res.SetStatus("HTTP/1.0 200 OK")
    // util.Dumpf("%#v", res)
    // util.Dumpf("%#v", res.Status().Code())

    // uri := http.NewUri("http://kerem:123@git.local.com")
    // uri := http.NewUri("http://kerem:123@git.local.com:8080/foo?a=the%20a!#xxx")
    // util.Dumps(uri)
    // util.Dumps(uri.Segments())
    // util.Dumps(uri.Segment(0))
    // util.Dumps(uri.Query())
    // util.Dumps(uri.QueryParams())
    // util.Dumps(uri.QueryParam("aaa"))
}
