package coinbase

import (
	"net/http"
	"strconv"
	"testing"
)

func testok(r *http.Request) (int, string) {
	return 200, "OK"
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

//register mock http handlers
//path, method, response file name, function to test request
func init() {
	handlers =
		[]mock{
			mock{"/accounts", "GET", "accounts_test.json", testok},
			mock{"/accounts/123", "GET", "account_test.json", testok},
			mock{"/accounts", "POST", "create_account_test.json", testok},
			mock{"/contacts", "GET", "contacts_test.json", testPageLimit},
			mock{"/users/self", "POST", "current_user_test.json", testok},
			mock{"/prices/buy", "GET", "buy_price_test.json", testok},
			mock{"/prices/sell", "GET", "sell_price_test.json", testok},
			mock{"/prices/spot_rate?currency=USD", "GET", "spot_price_test.json", testok},
		}
}

var mclient = mockClient()

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
	expect(t, accts[0].AsFloat("/native_balance/amount"), 500.12)
}

func TestClientAccount(t *testing.T) {
	acct, err := mclient.Account("123")
	expect(t, err, nil)
	b := acct.props.AsNode("/balance/amount").AsStr()
	expect(t, b, "50.00000000")
	expect(t, acct.AsFloat("/balance/amount"), 50.00000000)
}

func TestCreateAccount(t *testing.T) {

	var json = `{"account": {"name": "Savings Wallet"}}`
	acct, err := mclient.CreateAccount(json)
	expect(t, err, nil)
	n := acct.props.AsNode("/name").AsStr()
	expect(t, n, "Savings Wallet")
}

func TestContacts(t *testing.T) {
	contacts, err := mclient.Contacts(1, 25, "")
	expect(t, err, nil)
	expect(t, len(contacts), 2)
	email := contacts[0].props.AsNode("/email").AsStr()
	expect(t, email, "user1@example.com")
}

func TestCurrentUser(t *testing.T) {
	user, err := mclient.CurrentUser()
	expect(t, err, nil)
	expect(t, user.props.AsNode("/email").AsStr(), "user1@example.com")
	expect(t, user.AsStr("/email"), "user1@example.com")
}

func TestBuyPrice(t *testing.T) {
	price, err := mclient.BuyPrice(1)
	expect(t, err, nil)
	expect(t, price.AsNode("/total/amount").AsStr(), "10.35")
}

func TestSellPrice(t *testing.T) {
	price, err := mclient.SellPrice(1)
	expect(t, err, nil)
	expect(t, price.AsNode("/total/amount").AsStr(), "9.65")
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
