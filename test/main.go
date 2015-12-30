package main

import (
    "util"
    "http"
    // "http/status"
)

func main() {
    req := http.NewRequest()
    req.SetHeader("foo", "...")

    util.Dumpf("%#v", req)
}
