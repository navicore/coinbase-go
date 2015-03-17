package coinbase

import (
	"testing"

	"github.com/chrhlnd/dynjson"
)

var jsonid = `
{
	"id" : "123"
}
`

func TestAccount(t *testing.T) {
	props := dynjson.NewFromBytes([]byte(jsonid))
	acct1 := Account{Model{Base{}, Client{}, props}}
	if "123" != acct1.id() {
		t.Errorf("whoops.  bad id.")
	}
}

func TestAccountFromJson(t *testing.T) {
	props := dynjson.NewFromBytes([]byte(jsonid))
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

func TestBalance(t *testing.T) {
	bal, err := MockAcct.Balance()
	expect(t, err, nil)
	expect(t, bal.AsNode("/amount").AsStr(), "36.62800000")
}

func TestSetPrimary(t *testing.T) {
	success, err := MockAcct.SetPrimary()
	expect(t, err, nil)
	expect(t, success, true)
}

func TestDeleteAccount(t *testing.T) {
	success, err := MockAcct.Delete()
	expect(t, err, nil)
	expect(t, success, true)
}

func TestModifyAccount(t *testing.T) {
	acct, err := MockAcct.Modify(`{"account": {"name": "Satoshi Wallet"}}`)
	expect(t, err, nil)
	name, err := acct.Str("/name")
	expect(t, err, nil)
	expect(t, name, "Satoshi Wallet")
}

func TestGetAddresses(t *testing.T) {
	data, err := MockAcct.Addresses(1, 25, "")
	expect(t, err, nil)
	label, err := data.Node("/addresses/0/address/label")
	expect(t, label.IsNull(), false)
	l, e := label.Str()
	expect(t, e, nil)
	expect(t, l, "My Label")
}

func TestGetAddress(t *testing.T) {
	data, err := MockAcct.Address("503c46a4f8182b10650000ad")
	expect(t, err, nil)
	label, err := data.Node("/address/label")
	expect(t, label.IsNull(), false)
	l, e := label.Str()
	expect(t, e, nil)
	expect(t, l, "My Label")
}

func TestCreateAddress(t *testing.T) {
	data, err := MockAcct.CreateAddress(`
	{
		"address": {
				"callback_url": "http://www.example.com/callback",
				"label": "Dalmation donations"
		}
	}
	`)
	expect(t, err, nil)
	label, err := data.Node("/label")
	expect(t, label.IsNull(), false)
	l, e := label.Str()
	expect(t, e, nil)
	expect(t, l, "Dalmation donations")

}

func TestGetTxns(t *testing.T) {
}

func TestGetTxn(t *testing.T) {
}

func TestGetXfers(t *testing.T) {
}

func TestGetXfer(t *testing.T) {
}

func TestTransferMoney(t *testing.T) {
}

func TestSendMoney(t *testing.T) {
}

func TestRequestMoney(t *testing.T) {
}

func TestGetButton(t *testing.T) {
}

func TestCreateButton(t *testing.T) {
}

func TestGetOrders(t *testing.T) {
}

func TestGetOrder(t *testing.T) {
}

func TestCreateOrder(t *testing.T) {
}

func TestBuy(t *testing.T) {
}

func TestSell(t *testing.T) {
}
