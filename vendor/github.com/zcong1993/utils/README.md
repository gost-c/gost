# utils
[![Build Status](https://travis-ci.org/zcong1993/utils.svg?branch=master)](https://travis-ci.org/zcong1993/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/utils)](https://goreportcard.com/report/github.com/zcong1993/utils)
[![GoDoc](https://godoc.org/github.com/zcong1993/utils?status.svg)](https://godoc.org/github.com/zcong1993/utils)
<!--
[![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/utils)](https://goreportcard.com/report/github.com/zcong1993/utils)
[![Build Status](https://travis-ci.org/zcong1993/utils.svg?branch=master)](https://travis-ci.org/zcong1993/utils)
[![GoDoc](https://godoc.org/github.com/zcong1993/utils?status.svg)](https://godoc.org/github.com/zcong1993/utils)
[![codecov](https://codecov.io/gh/zcong1993/utils/branch/master/graph/badge.svg)](https://codecov.io/gh/zcong1993/utils)
-->

> Some helper functions for go

## Functions

### SliceIndex
return the index of the first element in the array that satisfies the provided testing function. Otherwise -1 is returned

### GetJson
GetJson make a get request use http.Get
```go
// define a response data type
type Resp struct {
    Status int    `decoder:"status"`
    Msg    string `decoder:"msg"`
}
var res Resp
// this api is always return json data {"status":200, "msg": "hello world}
err := utils.GetJSON("http://zcong-hello.getsandbox.com/hello", &res)
// check error
if err != nil {
    panic(err)
}
// res is response json data
fmt.Printf("%v", res)
//Output: {200 hello world}
```

### GetJsonWithHeaders
GetJsonWithHeaders make a http get request with custom headers
```go
// define your response data type
type Resp struct {}
// define some custom headers
customHeaders := map[string]string{"Foo": "bar"}
var res Resp
// this api return all the request headers
err := utils.GetJSONWithHeaders("http://zcong-hello.getsandbox.com/header", &res, customHeaders)
// res is response data
```
### Compile
Compile is a html template compiler with custom tpl and data
```go
tpl := "hello {{.data}}"
data := map[string]string{"data": "world"}
var d bytes.Buffer
err := utils.Compile(&d, tpl, &data)
// check error
if err != nil {
    panic(err)
}
fmt.Printf("%s", d.String())
// hello world
```

### CompileText
CompileText is same as Compile but use text/template
```go
tpl := "hello {{.data}}"
data := map[string]string{"data": "world"}
var d bytes.Buffer
err := utils.CompileText(&d, tpl, &data)
// check error
if err != nil {
    panic(err)
}
fmt.Printf("%s", d.String())
// hello world

### StringAddress
Get address of const string
```go
str := "zcong1993"
add := StringAddress(str)
// now add is &str without error
```

## License

MIT &copy; zcong1993
