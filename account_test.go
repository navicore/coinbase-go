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
	_, err := MockAcct.Balance()
	expect(t, err, nil)
	//expect(t, bal.AsNode("/amount").AsStr(), "50.00000000")
}
