package coinbase

import (
	"net/http"
	"strconv"
	"testing"
)

func testok(*http.Request) (int, string) {
	return 200, "OK"
}

func init() {
	handlers =
		[]hdlr{
			hdlr{"/accounts", "GET", "accounts_test.json", testok},
			hdlr{"/account", "GET", "account_test.json", testok},
		}
}

func TestAccountApi(t *testing.T) {

	c := mockClient()
	accts, err := c.Accounts()
	expect(t, err, nil)
	expect(t, len(accts), 3)
	b := accts[1].props.AsNode("/balance/amount").AsStr()
	fb, _ := strconv.ParseFloat(b, 64)
	expect(t, fb, 508.94)
}
