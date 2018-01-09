package utils_test

import (
	"fmt"
	"github.com/zcong1993/utils"
	"os"
)

// Example (simple) make a simple http get request for json data
func ExampleGetJSON() {
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
	fmt.Printf("%+v", res)
	//Output: {Status:200 Msg:hello world}
}

// Example (custom headers) can add some custom headers when request
func ExampleGetJSONWithHeaders() {
	// define a response data type
	type Header struct {
		Name  string `decoder:"name"`
		Value string `decoder:"value"`
	}
	type Headers struct {
		Headers []Header `decoder:"headers"`
	}
	// define some custom headers
	customHeaders := map[string]string{"Foo": "bar"}
	var res Headers
	// this api return all the request headers
	err := utils.GetJSONWithHeaders("http://zcong-hello.getsandbox.com/header", &res, customHeaders)
	if err != nil {
		panic(err)
	}
	// custom headers should in response data
	index := utils.SliceIndex(len(res.Headers), func(num int) bool {
		return res.Headers[num] == Header{"Foo", "bar"}
	})
	fmt.Printf("%+v", res.Headers[index])
	//Output: {Name:Foo Value:bar}
}

// Example PostJSON make a simple http post request
func ExamplePostJSON() {
	// define a response data type
	type User struct {
		Username string `decoder:"username"`
		Age      int    `decoder:"age"`
	}
	type Resp struct {
		Status string `decoder:"status"`
		User   User   `decoder:"user"`
	}
	// post body
	user := User{"zcong", 18}
	var res Resp
	// this api return post body in response.user
	err := utils.PostJSON("http://zcong-hello.getsandbox.com/users", &user, &res, map[string]string{})
	// check error
	if err != nil {
		panic(err)
	}
	// res is response json data
	fmt.Printf("%+v", res)
	//Output: {Status:ok User:{Username:zcong Age:18}}
}

// Example compile inline template with data, put result to io.Write you defined
func ExampleCompile() {
	tpl := "hello {{.data}}"
	data := map[string]string{"data": "world"}
	err := utils.Compile(os.Stdout, tpl, &data)
	// check error
	if err != nil {
		panic(err)
	}
	//Output: hello world
}

// Example SliceIndex returns the index of the first element in the array that satisfies the provided testing function. Otherwise -1 is returned
func ExampleSliceIndex() {
	arr := []int{1, 2, 3, 4}
	// test function, find index of 3 in arr
	condition := func(i int) bool {
		return arr[i] == 3
	}
	index := utils.SliceIndex(len(arr), condition)
	fmt.Println(index)
	//Output: 2
}
