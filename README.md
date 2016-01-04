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
