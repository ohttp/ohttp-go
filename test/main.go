package main

import (
    "http"
    "http/util"
    "http/util/params"
    // "http/util/query"
    "http/request"
    "http/response"
)

func main() {
    o := params.New()
    o.Set("debug", true)

    c := http.NewClient(o)

    // r := c.Do("GET http://localhost/foo.json", nil, nil, nil)
    // util.Dumpf("%#v", r)
    // util.Dumpf("%+v", c.Request().String())
    // util.Dumpf("%+v", c.Response().String())

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
