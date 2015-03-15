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

	"github.com/chrhlnd/dynjson"
)

var propsJson = `
{
	"id" : "123"
}
`

func TestAccount(t *testing.T) {
	props := dynjson.NewFromBytes([]byte(propsJson))
	acct1 := Account{Model{Base{}, Client{}, props}}
	if "123" != acct1.id() {
		t.Errorf("whoops.  bad id.")
	}
}

func TestAccountFromJson(t *testing.T) {
	props := dynjson.NewFromBytes([]byte(propsJson))
	acct1 := NewAccount(props, Client{})
	if "123" != acct1.id() {
		t.Errorf("whoops.  bad id.")
	}
}

func TestAccountFromId(t *testing.T) {
	acct1 := NewAccountFromId("4321", Client{})
	if "4321" != acct1.id() {
		t.Errorf("whoops.  bad id.")
	}
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func client(t *testing.T) Client {

	mux := http.NewServeMux()

	acctshandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		acctsjson, _ := ioutil.ReadFile("./test_accounts.json")
		fmt.Fprintln(w, string(acctsjson))
	})
	mux.Handle("/accounts", acctshandler)

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
