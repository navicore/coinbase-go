package coinbase

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var handlers = []mock{
	mock{"/accounts", "GET", "testjson/accounts_test.json", testok},
	mock{"/accounts/123", "GET", "testjson/account_test.json", testok},
	mock{"/accounts/536a541fa9393bb3c7000023", "GET", "testjson/account_test.json", testok},
	mock{"/accounts/536a541fa9393bb3c7000023/balance", "GET", "testjson/account_balance_test.json", testok},
	mock{"/accounts", "POST", "testjson/create_account_test.json", testok},
	mock{"/contacts", "GET", "testjson/contacts_test.json", testPageLimit},
	mock{"/users/self", "POST", "testjson/current_user_test.json", testok},
	mock{"/prices/buy", "GET", "testjson/buy_price_test.json", testok},
	mock{"/prices/sell", "GET", "testjson/sell_price_test.json", testok},
	mock{"/prices/spot_rate", "GET", "testjson/spot_price_test.json", testok},
	mock{"/currencies", "GET", "testjson/currencies_test.json", testok},
	mock{"/exchange_rates", "GET", "testjson/rates_test.json", testok},
	mock{"/payment_methods", "GET", "testjson/paymethods_test.json", testok},
	mock{"/payment_methods/530eb5b217cb34e07a000011", "GET", "testjson/paymethod_test.json", testok},
	mock{"/accounts/536a541fa9393bb3c7000023/primary", "POST", "testjson/success_test.json", testok},
	mock{"/accounts/536a541fa9393bb3c7000023", "DELETE", "testjson/success_test.json", testok},
}
var MockClient = mockClient()
var MockAcct, _ = MockClient.Account("536a541fa9393bb3c7000023")

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func testPageLimit(r *http.Request) (int, string) {
	page := r.FormValue("page")
	if page != "1" {
		return 400, "what?"
	}
	limit := r.FormValue("limit")
	if limit != "25" {
		return 400, "what?"
	}
	return 200, "OK"
}

func testok(r *http.Request) (int, string) {
	return 200, "OK"
}

type mock struct {
	Path     string
	Method   string
	Filename string
	Test     func(*http.Request) (int, string)
}

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

func mockClient() Client {

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
