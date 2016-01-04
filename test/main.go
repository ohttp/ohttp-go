package main

import (
    "ohttp"
    "ohttp/request"
    "ohttp/response"
    "ohttp/util"
    "ohttp/util/params"
    // "ohttp/util/query"
)

func init() {
    ohttp.Shutup()
    request.Shutup()
    response.Shutup()
    util.Shutup()
}

func main() {
    // r1, _, e1 := util.RegExpMatch(rs, re)
    // // r2, _, e2 := util.RegExpMatchName(rs, re)
    // util.Dumpf("%#v %#v %d", r1, e1, len(r1))
    // // util.Dumpf("%#v %#v %d", r2["method"], e2, len(r2))

    o := params.New()
    o.Set("debug", true)

    c := ohttp.NewClient(o)

    // default port 80
    r, err := c.Get("127.0.0.1:5984", params.Params{"a": 1}, params.Params{"X-foo": true})
    if err != nil {
        panic(err)
    }
    util.Dump(r.Status().Code())

    // r, err := c.Do("GET localhost:5984", nil, nil, nil)
    // if err != nil {
    //     panic(err)
    // }
    // util.Dumpf("%#v", r.Status().Code())
    // util.Dumpf("%+v", c.Request().String())
    // util.Dumpf("%+v", c.Response().String())
    // c.DoFunc("GET http://localhost/foo.json", nil, nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Options("http://localhost/foo.json", nil, nil)
    // util.Dumpf("%#v", r)
    // c.OptionsFunc("http://localhost/foo.json", nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Head("http://localhost/foo.json", nil, nil)
    // util.Dumpf("%#v", r)
    // c.HeadFunc("http://localhost/foo.json", nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Get("http://localhost/foo.json", nil, nil)
    // util.Dumpf("%#v", r)
    // c.GetFunc("http://localhost/foo.json", nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Post("http://localhost/foo.json", nil, nil, nil)
    // util.Dumpf("%#v", r)
    // c.PostFunc("http://localhost/foo.json", nil, nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Put("http://localhost/foo.json", nil, nil, nil)
    // util.Dumpf("%#v", r)
    // c.PutFunc("http://localhost/foo.json", nil, nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Patch("http://localhost/foo.json", nil, nil, nil)
    // util.Dumpf("%#v", r)
    // c.PatchFunc("http://localhost/foo.json", nil, nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Delete("http://localhost/foo.json", nil, nil)
    // util.Dumpf("%#v", r)
    // c.DeleteFunc("http://localhost/foo.json", nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Trace("http://localhost/foo.json", nil, nil)
    // util.Dumpf("%#v", r)
    // c.TraceFunc("http://localhost/foo.json", nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Connect("http://localhost/foo.json", nil, nil)
    // util.Dumpf("%#v", r)
    // c.ConnectFunc("http://localhost/foo.json", nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Copy("http://localhost/foo.json", nil, nil)
    // util.Dumpf("%#v", r)
    // c.CopyFunc("http://localhost/foo.json", nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // r, _ := c.Move("http://localhost/foo.json", nil, nil)
    // util.Dumpf("%#v", r)
    // c.MoveFunc("http://localhost/foo.json", nil, nil,
    //     func(req *request.Request, res *response.Response, err error) {
    //     if err != nil {
    //         panic(err)
    //     }
    //     util.Dump(req.String())
    //     util.Dump(res.String())
    // })

    // q := query.New(map[string]interface{}{"a": true, "b": 1})
    // util.Dumpf("%#v", q.String())
    // util.Dumpf("%#v", q.Params().String())

    // req := ohttp.NewRequest()
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

    // res := ohttp.NewResponse()
    // res.SetStatus("HTTP/1.0 200 OK")
    // util.Dumpf("%#v", res)
    // util.Dumpf("%#v", res.Status().Code())

    // uri := ohttp.NewUri("http://kerem:123@git.local.com")
    // uri := ohttp.NewUri("http://kerem:123@git.local.com:8080/foo?a=the%20a!#xxx")
    // util.Dumps(uri)
    // util.Dumps(uri.Segments())
    // util.Dumps(uri.Segment(0))
    // util.Dumps(uri.Query())
    // util.Dumps(uri.QueryParams())
    // util.Dumps(uri.QueryParam("aaa"))
}
