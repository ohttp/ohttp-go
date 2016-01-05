Oh! HTTP is a toolbox that make easy to dial with HTTP URL's for you. It contains URL / header parser and much more utilities.

## Simply

```go
import "ohttp"
import "ohttp/request"
import "ohttp/response"
import "ohttp/util"

client := ohttp.NewClient(nil)

// default port 80
res, err := client.Get("github.com", nil, nil)
if err != nil {
    panic(err)
}
util.Dump(res.Status().Code()) // => 301

// with callback
client.GetFunc("github.com", nil, nil,
    func(req *request.Request, res *response.Response, err error) {
        if err != nil {
            panic(err)
        }
        util.Dump(req.String()) // dump whole message data
        util.Dump(res.String()) // dump whole message data
    })
```

## URL's

Available formats: `github.com:80`, `http://github.com`, `github.com:443`, `https://github.com`...

```go
// connect via SSL
client.Get("github.com:443", nil, nil)
client.Get("https://github.com", nil, nil)

// connect to CouchDB
client.Get("localhost:5984", nil, nil)
```

## Client

```go
import "ohttp/util/params"

/* available args */
// url
u  := "localhost"
// url params (nullable)
up := params.Params{"a": 1}
// body (nullable, string or map for JSON payloads, only for POST|PUT|PATCH)
b  := "Hello, world!"
// headers (nullable)
h  := params.Params{"X-foo": true}

// main request func is Do()
res, err := client.Do(u, up, b, h)

// or same with callback
client.DoFunc(u, up, b, h, func(req, res, err))

/* OPTIONS Request's */
// return
res, err := client.Options(u, up, h)
// callback
client.OptionsFunc(u, up, h, func(req, res, err))

/* HEAD Request's */
// return
res, err := client.Head(u, up, h)
// callback
client.HeadFunc(u, up, h, func(req, res, err))

/* GET Request's */
// return
res, err := client.Get(u, up, h)
// callback
client.GetFunc(u, up, h, func(req, res, err))

/* POST Request's */
// return
res, err := client.Post(u, up, b, h)
// callback
client.PostFunc(u, up, b, h, func(req, res, err))

/* PUT Request's */
// return
res, err := client.Put(u, up, b, h)
// callback
client.PutFunc(u, up, b, h, func(req, res, err))

/* PATCH Request's */
// return
res, err := client.Patch(u, up, b, h)
// callback
client.PatchFunc(u, up, b, h, func(req, res, err))

/* DELETE Request's */
// return
res, err := client.Delete(u, up, h)
// callback
client.DeleteFunc(u, up, h, func(req, res, err))

/* TRACE Request's */
// return
res, err := client.Trace(u, up, h)
// callback
client.TraceFunc(u, up, h, func(req, res, err))

/* CONNECT Request's */
// return
res, err := client.Connect(u, up, h)
// callback
client.ConnectFunc(u, up, h, func(req, res, err))

/* COPY Request's */
// return
res, err := client.Copy(u, up, h)
// callback
client.CopyFunc(u, up, h, func(req, res, err))

/* MOVE Request's */
// return
res, err := client.Move(u, up, h)
// callback
client.MoveFunc(u, up, h, func(req, res, err))
```

Notice: Constructor methods could be named as just `New`, also be named as `NewClient` etc. But all setter methods have a `Set` prefix, and getters have not `Get` prefix. For example, if you want to get client's request object, just do it like `req := client.Request()`.

## Client in Action

```go
// check a service status
client.HeadFunc("github.com:443", nil, nil,
  func(req *request.Request, res *response.Response, err error) {
    if err != nil {
      panic(err)
    }
    if !res.OK() {
      panic("Mayday mayday! GitHub down, it's fuckin' down!")
    }
  })
```

## Message

Actually, even you don't need many stuffs below, but suppose should be explained all.

Both `Request` and `Response` object extends / implements `ohttp.message.Message` object. So it has many methods that comes from `Message` object. Here are some of them that you may want/available to use.

```go
// m is just a pseudo, it could be client's request or response
m := client.Request()

// @return ohttp.message.(TYPE_REQUEST|TYPE_RESPONSE)
t := m.Type()

// @return ohttp.message.(PROTOCOL_VERSION_1_0|PROTOCOL_VERSION_1_1|PROTOCOL_VERSION_2_0)
// @default 1.0
pv := m.ProtocolVersion()

// @return string|""
h := m.Header("Content-Type")

// @return map[string]string
ha := m.HeaderAll()

// @return *ohttp.message.MessageBody
b := m.Body()
  bc  := m.Body().Content()       => string
  bct := m.Body().ContentType()   => string
  bcl := m.Body().ContentLength() => int

// @return *ohttp.util.params.Params
o := m.Options()
  x := m.Options().Get()          => interface{}
  x := m.Options().GetInt()       => int
  x := m.Options().GetUInt()      => uint
  x := m.Options().GetString()    => string
  x := m.Options().GetBool()      => bool
  x := m.Options().Empty()        => bool
  x := m.Options().Array()        => map[string]interface{}
  x := m.Options().String()       => string // a=1&b=2...

// @return *ohttp.message.MessageError
e := m.Error()
  ec := m.Error().Code()          => int
  ec := m.Error().Text()          => string

// @return bool
ok := req.OK()

// @param uint (see types above)
m := m.SetType(t)

// @param string (see protocols above)
m := m.SetProtocolVersion(pv)

// @param k string
// @param v string
m := m.SetHeader(k, v)

// @param kv map[string]string|params.Params|*params.Params
m := m.SetHeaderAll(kv)

// @param ec uint
// @param et string
m.SetError(ec, et)

// @param b interface{}
m := m.SetBody(b)
```

## Request

```go
req := client.Request()

// @return GET|POST..
m := req.Method()

// @return *ohttp.uri.Uri
u := req.Uri()
  x := req.Uri().Source()         => string // github.com:443/foo/?a=1#xyz
  x := req.Uri().Scheme()         => string
  x := req.Uri().Host()           => string
  x := req.Uri().Port()           => uint
  x := req.Uri().Path()           => string
  x := req.Uri().Username()       => string
  x := req.Uri().Password()       => string
  x := req.Uri().Query()          => *ohttp.util.query.Query
    x := req.Uri().Get(k string)  => interface{}
    x := req.Uri().Params()       => *ohttp.util.params.Params
      // see above
    x := req.Uri().String()       => string // a=1
  x := req.Uri().Fragment()       => string
  x := req.Uri().Segment(i int)   => string|""
  x := req.Uri().Segments()       => []string
  x := req.Uri().Authorization()  => "username:password"|"username"|""

// @return string
s := req.String()
// GET /foo?a=1 HTTP/1.0
// Host: github.com
// ...
// [body]
```

## Response

```go
res := client.Response()

// @return *ohttp.response.status.Status
s := res.Status()
  x := res.Status().Code()        => int
  x := res.Status().Text()        => string
  x := res.Status().Status()      => string // HTTP/1.0 200 OK

// @return string
s := res.String()
// HTTP/1.1 200 OK
// Server: GitHub.com
// Date: Tue, 05 Jan 2016 19:49:00 GMT
// ...
// [body]
```
