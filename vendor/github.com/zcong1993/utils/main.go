package utils

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strings"
	text "text/template"
)

// GetJSON make a get request use http.Get
func GetJSON(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}

// GetJSONWithHeaders make a http get request with custom headers
func GetJSONWithHeaders(url string, v interface{}, headers map[string]string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}

// SliceIndex return the index of the first element in the array that satisfies the provided testing function.
//
// Otherwise -1 is returned
func SliceIndex(limit int, condition func(num int) bool) int {
	for i := 0; i < limit; i++ {
		if condition(i) {
			return i
		}
	}
	return -1
}

// PostJSON make a http post request with custom body and headers
func PostJSON(url string, body interface{}, v interface{}, headers map[string]string) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}

// Compile is a html template compiler with custom tpl and data
func Compile(w io.Writer, tpl string, data interface{}) error {
	t := template.New("compiler")
	t, err := t.Parse(tpl)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

// CompileText is same as Compile but use text/template
func CompileText(w io.Writer, tpl string, data interface{}) error {
	t := text.New("compiler-text")
	t, err := t.Parse(tpl)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

// StringAddress is a helper function to get const string address
func StringAddress(v string) *string {
	return &v
}
