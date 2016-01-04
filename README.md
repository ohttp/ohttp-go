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

/* GET Request */
// return
res, err := client.Get(u, up, h)
// callback
client.GetFunc(u, up, h, func(req, res, err))
```
