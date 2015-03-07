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
	acct1 := Account{AccountBase{Base{props}}}
	if "123" != acct1.id() {
		t.Errorf("whoops.  bad id.")
	}
}
