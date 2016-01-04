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

```go
// all available below
"github.com:80"
"http://github.com"
"github.com:443"
"https://github.com"

// ie: connect local couchdb
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

// the main request func's: Do()
res, err := client.Do(u, up, b, h)

// or with callback
client.DoFunc(u, up, b, h, func(req, res, err))
```
