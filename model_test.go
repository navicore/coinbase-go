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
	acct1 := Account{AccountBase{Model{Base{}, props}}}
	if "123" != acct1.id() {
		t.Errorf("whoops.  bad id.")
	}
}
func TestAccountWithId(t *testing.T) {
	acct1 := NewAccount("4321")
	if "4321" != acct1.id() {
		t.Errorf("whoops.  bad id.")
	}
}
