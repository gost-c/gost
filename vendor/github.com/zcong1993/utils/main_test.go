package utils

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetJSON(t *testing.T) {
	type Resp struct {
		Status int    `decoder:"status"`
		Msg    string `decoder:"msg"`
	}
	var res Resp
	err := GetJSON("http://zcong-hello.getsandbox.com/hello", &res)
	assert.Nil(t, err)
	assert.Equal(t, res, Resp{200, "hello world"})
	err = GetJSON("http://example.com", &res)
	assert.NotNil(t, err)
}

func TestGetJSONWithHeaders(t *testing.T) {
	type Header struct {
		Name  string `decoder:"name"`
		Value string `decoder:"value"`
	}
	type Headers struct {
		Headers []Header `decoder:"headers"`
	}
	var res Headers
	err := GetJSONWithHeaders("http://zcong-hello.getsandbox.com/header", &res, map[string]string{"Content-Type": "application/json", "Foo": "bar"})
	assert.Nil(t, err)
	headers := res.Headers
	index1 := SliceIndex(len(headers), func(num int) bool {
		return headers[num] == Header{"Content-Type", "application/json"}
	})
	index2 := SliceIndex(len(headers), func(num int) bool {
		return headers[num] == Header{"Foo", "bar"}
	})
	assert.True(t, index1 > -1)
	assert.True(t, index2 > -1)
}

func TestSliceIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	condition := func(i int) bool {
		return arr[i] == 3
	}
	index := SliceIndex(len(arr), condition)
	noExists := SliceIndex(len(arr), func(i int) bool {
		return arr[i] == 10
	})
	assert.Equal(t, index, 2)
	assert.Equal(t, noExists, -1)
}

func TestPostJSON(t *testing.T) {
	type User struct {
		Username string `decoder:"username"`
		Age      int    `decoder:"age"`
	}
	type Resp struct {
		Status string `decoder:"status"`
		User   User   `decoder:"user"`
	}
	var res Resp
	err := PostJSON("http://zcong-hello.getsandbox.com/users", User{"zcong", 18}, &res, map[string]string{})
	assert.Nil(t, err)
	assert.Equal(t, res, Resp{"ok", User{"zcong", 18}})
}

func TestCompile(t *testing.T) {
	tpl := "hello {{.data}}"
	data := map[string]string{"data": "world"}
	var res bytes.Buffer
	err := CompileText(&res, tpl, data)
	assert.Nil(t, err)
	assert.Equal(t, res.String(), "hello world")
}

func TestCompileText(t *testing.T) {
	tpl := "hello {{.data}}"
	data := map[string]string{"data": "world"}
	var res bytes.Buffer
	err := Compile(&res, tpl, data)
	assert.Nil(t, err)
	assert.Equal(t, res.String(), "hello world")
}

func TestStringAddress(t *testing.T) {
	str := "zcong1993"
	add := StringAddress(str)
	var mock *string
	assert.Equal(t, reflect.TypeOf(add), reflect.TypeOf(mock))
	assert.Equal(t, *add, str)
}
