package coinbase

import (
	"strconv"
	"testing"
)

var mc = MockClient

func TestClientAccounts(t *testing.T) {
	accts, err := mc.Accounts()
	expect(t, err, nil)
	expect(t, len(accts), 2)
	b := accts[0].props.AsNode("/balance/amount").AsStr()
	fb, _ := strconv.ParseFloat(b, 64)
	expect(t, fb, float64(50))
	b = accts[0].props.AsNode("/native_balance/amount").AsStr()
	ffb, _ := strconv.ParseFloat(b, 64)
	expect(t, ffb, 500.12)
	f, _ := accts[0].Float("/native_balance/amount")
	expect(t, f, 500.12)
}

func TestClientAccount(t *testing.T) {
	acct, err := mc.Account("123")
	expect(t, err, nil)
	b := acct.props.AsNode("/balance/amount").AsStr()
	expect(t, b, "50.00000000")
	f, _ := acct.Float("/balance/amount")
	expect(t, f, 50.00000000)
}

func TestCreateAccount(t *testing.T) {

	var json = `{"account": {"name": "Savings Wallet"}}`
	acct, err := mc.CreateAccount(json)
	expect(t, err, nil)
	n, _ := acct.Str("/name")
	expect(t, n, "Savings Wallet")
}

func TestContacts(t *testing.T) {
	contacts, err := mc.Contacts(1, 25, "")
	expect(t, err, nil)
	expect(t, len(contacts), 2)
	email, _ := contacts[0].Str("/email")
	expect(t, email, "user1@example.com")
}

func TestCurrentUser(t *testing.T) {
	user, err := mc.CurrentUser()
	expect(t, err, nil)
	expect(t, user.props.AsNode("/email").AsStr(), "user1@example.com")
	email, _ := user.Str("/email")
	expect(t, email, "user1@example.com")
}

func TestBuyPrice(t *testing.T) {
	price, err := mc.BuyPrice(1)
	expect(t, err, nil)
	expect(t, price.AsNode("/total/amount").AsStr(), "10.35")
}

func TestSellPrice(t *testing.T) {
	price, err := mc.SellPrice(1)
	expect(t, err, nil)
	p, _ := price.AsNode("/total/amount").Str()
	expect(t, p, "9.65")
}

func TestSpotPrice(t *testing.T) {
	price, err := mc.SpotPrice("USD")
	expect(t, err, nil)
	expect(t, price.AsNode("/amount").AsStr(), "10.00")
}

func TestCurrencies(t *testing.T) {
	data, err := mc.Currencies()
	expect(t, err, nil)
	expect(t, data.AsNode("/2/1").AsStr(), "DZD")
}

func TestCreateUser(t *testing.T) {
	data, err := mc.Rates()
	expect(t, err, nil)
	expect(t, data.AsNode("/zwl_to_btc").AsStr(), "0.00001")
}

func TestPayMethods(t *testing.T) {
	pms, err := mc.PayMethods()
	expect(t, err, nil)
	expect(t, len(pms), 2)
	name, e := pms[0].Str("/name")
	expect(t, e, nil)
	expect(t, name, "US Bank ****4567")
}

func TestPayMethod(t *testing.T) {
	pm, err := mc.PayMethod("530eb5b217cb34e07a000011")
	expect(t, err, nil)
	name, _ := pm.Str("/name")
	expect(t, name, "US Bank ****4567")
}
