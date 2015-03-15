package coinbase

import (
	"net/http"
	"strconv"
	"testing"
)

func testok(*http.Request) (int, string) {
	return 200, "OK"
}

var mclient = mockClient()

func init() {
	handlers =
		[]hdlr{
			hdlr{"/accounts", "GET", "accounts_test.json", testok},
			hdlr{"/accounts/123", "GET", "account_test.json", testok},
			hdlr{"/accounts", "POST", "create_account_test.json", testok},
			hdlr{"/contacts", "GET", "contacts_test.json", testok},
			hdlr{"/users/self", "POST", "current_user_test.json", testok},
		}
}

func TestClientAccounts(t *testing.T) {
	accts, err := mclient.Accounts()
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
	acct, err := mclient.Account("123")
	expect(t, err, nil)
	b := acct.props.AsNode("/balance/amount").AsStr()
	expect(t, b, "50.00000000")
}

func TestCreateAccount(t *testing.T) {

	var json = `{"account": {"name": "Savings Wallet"}}`
	//args := dynjson.NewFromBytes([]byte(json))
	acct, err := mclient.CreateAccount(json)
	expect(t, err, nil)
	n := acct.props.AsNode("/name").AsStr()
	expect(t, n, "Savings Wallet")
}

func TestContacts(t *testing.T) {
	contacts, err := mclient.Contacts(1, 10, "")
	expect(t, err, nil)
	expect(t, len(contacts), 2)
	email := contacts[0].props.AsNode("/email").AsStr()
	expect(t, email, "user1@example.com")
}

func TestCurrentUser(t *testing.T) {
	user, err := mclient.CurrentUser()
	expect(t, err, nil)
	expect(t, user.props.AsNode("/email").AsStr(), "user1@example.com")
}

func TestBuyPrice(t *testing.T) {
	//TODO
}

func TestSellPrice(t *testing.T) {
	//TODO
}

func TestSpotPrice(t *testing.T) {
	//TODO
}

func TestCurrencies(t *testing.T) {
	//TODO
}

func TestCreateUser(t *testing.T) {
	//TODO
}

func TestPayMethods(t *testing.T) {
	//TODO
}

func TestPayMethod(t *testing.T) {
	//TODO
}
