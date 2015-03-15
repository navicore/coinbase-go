package coinbase

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func testok(*http.Request) (int, string) {
	return 200, "OK"
}
func initrouter() {
	handlers = append(handlers, hdlr{"/accounts", "GET", "test_accounts.json", testok})
	handlers = append(handlers, hdlr{"/account", "GET", "test_accounts.json", func(*http.Request) (int, string) { return 200, "OK" }})
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

type hdlr struct {
	Path     string
	Method   string
	Filename string
	Test     func(*http.Request) (int, string)
}

var handlers = make([]hdlr, 0, 50)

func route(r *http.Request) (int, string) {
	method := r.Method
	path := r.URL.Path
	for i := range handlers {
		if path == handlers[i].Path {
			if method == handlers[i].Method {
				code, msg := handlers[i].Test(r)
				if code != 200 {
					return code, msg
				}
				json, _ := ioutil.ReadFile(handlers[i].Filename)
				return 200, string(json)
			}
		}
	}
	return 404, "bad path or method"
}
func client(t *testing.T) Client {

	initrouter()

	mux := http.NewServeMux()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code, json := route(r)
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, json)
	})
	mux.Handle("/", handler)

	server := httptest.NewServer(mux)
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}

	c := Client{Base{}, *httpClient, server.URL, "mykey", "mysecret", "mytoken", "myrtoken"}
	return c
}

func TestAccountApi(t *testing.T) {

	c := client(t)
	accts, err := c.Accounts()
	expect(t, err, nil)
	expect(t, len(accts), 3)
	b := accts[1].props.AsNode("/balance/amount").AsStr()
	fb, _ := strconv.ParseFloat(b, 64)
	expect(t, fb, 508.94)
}
