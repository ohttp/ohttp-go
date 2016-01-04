Oh! HTTP, is a toolbox that make easy to dial with `http` URL's for you. It contains URL / header parser and much more utilities.

## Simply

```go
import "ohttp"
import "ohttp/util"

client := ohttp.NewClient(nil)

// default port 80
res, err := client.Get("github.com", nil, nil)
if err != nil {
    panic(err)
}
// => 301
util.Dump(res.Status().Code())
```

## URL's

Available formats: `github.com:80`, `http://github.com`, `github.com:443`, `https://github.com`...

```go
// connect via ssl
client.Get("github.com:443", nil, nil)
client.Get("https://github.com", nil, nil)

// connect couchdb
client.Get("127.0.0.1:5984", nil, nil)
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
b  := "hello=world!"
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

## Request

## Response
