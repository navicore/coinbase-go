package coinbase

import (
	"net/http"
	"strconv"
	"testing"
)

func testok(*http.Request) (int, string) {
	return 200, "OK"
}

var client = mockClient()

func init() {
	handlers =
		[]hdlr{
			hdlr{"/accounts", "GET", "accounts_test.json", testok},
			hdlr{"/accounts/123", "GET", "account_test.json", testok},
		}
}

func TestClientAccounts(t *testing.T) {

	accts, err := client.Accounts()
	expect(t, err, nil)
	expect(t, len(accts), 2)
	b := accts[0].props.AsNode("/balance/amount").AsStr()
	fb, _ := strconv.ParseFloat(b, 64)
	expect(t, fb, float64(50))
	b = accts[0].props.AsNode("/native_balance/amount").AsStr()
	ffb, _ := strconv.ParseFloat(b, 64)
	expect(t, ffb, 500.12)
}

func TestClientAccount(t *testing.T) {

	acct, err := client.Account("123")
	expect(t, err, nil)
	b := acct.props.AsNode("/balance/amount").AsStr()
	expect(t, b, "50.00000000")
}
