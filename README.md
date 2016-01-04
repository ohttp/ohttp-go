Oh! HTTP, is a toolbox that make easy to dial with `http` URL's for you. It contains URL / header parser and much more utilities.

## Simply
```go
import "ohttp"
import "ohttp/util"

c := ohttp.NewClient(nil)

// default port 80
r, err := c.Get("github.com", nil, nil)
if err != nil {
    panic(err)
}
// => 301
util.Dump(r.Status().Code())
```
