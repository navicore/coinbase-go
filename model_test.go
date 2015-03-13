package coinbase

import (
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
