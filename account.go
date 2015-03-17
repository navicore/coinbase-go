package coinbase

import (
	"fmt"

	"github.com/chrhlnd/dynjson"
)

type Account struct {
	Model
}

func NewAccountFromId(id string, client Client) Account {
	props := Id(id)
	return Account{Model{Base{}, client, props}}
}

func NewAccountFromProps(props dynjson.DynNode, client Client) Account {
	return Account{Model{Base{}, client, props}}
}

func NewAccount(props dynjson.DynNode, client Client) Account {
	return Account{Model{Base{}, client, props}}
}

func (this Account) Delete() (bool, error) {
	var id string
	var root, props dynjson.DynNode
	var err error
	if id, err = this.Str("/id"); err != nil {
		return false, err
	}
	path := fmt.Sprintf("/accounts/%v", id)
	root, err = this.client.DelDynNode(path, "")
	if err != nil {
		return false, err
	}
	if props, err = root.Node("/success"); err != nil {
		return false, err
	}
	if props.IsNull() {
		return false, fmt.Errorf("node not found")
	}
	return props.AsBool(), nil
}

func (this Account) SetPrimary() (bool, error) {
	var id string
	var root, props dynjson.DynNode
	var err error
	if id, err = this.Str("/id"); err != nil {
		return false, err
	}
	path := fmt.Sprintf("/accounts/%v/primary", id)
	root, err = this.client.PostDynNode(path, "")
	if err != nil {
		return false, err
	}
	if props, err = root.Node("/success"); err != nil {
		return false, err
	}
	if props.IsNull() {
		return false, fmt.Errorf("node not found")
	}
	return props.AsBool(), nil
}

func (this Account) Modify(args string) (Account, error) {
	var id string
	var root, props dynjson.DynNode
	var err error
	if id, err = this.Str("/id"); err != nil {
		return Account{}, err
	}
	path := fmt.Sprintf("/accounts/%v", id)
	if root, err = this.client.PutDynNode(path, args); err != nil {
		return Account{}, err
	}
	if root.IsNull() {
		return Account{}, fmt.Errorf("node not found")
	}
	if props, err = root.Node("/account"); err != nil {
		return Account{}, err
	}
	return NewAccountFromProps(props, this.client), nil
}

func (this Account) Balance() (dynjson.DynNode, error) {
	var id string
	var err error
	if id, err = this.Str("/id"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/accounts/%v/balance", id)
	return this.client.GetDynNode(path, nil)
}

func (this Account) Addresses(page int, limit int, query string) (dynjson.DynNode, error) {
	var id string
	var err error
	if id, err = this.Str("/id"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/addresses?account_id=%v&page=%v&limit=%v", id, page, limit)
	if query != "" {
		path = path + "&query=" + query
	}
	return this.client.GetDynNode(path, nil)
}

func (this Account) Address(address_id string) (dynjson.DynNode, error) {
	var id string
	var err error
	if id, err = this.Str("/id"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/addresses/%v?account_id=%v", address_id, id)
	return this.client.GetDynNode(path, nil)
}

func (this Account) CreateAddress(args string) (dynjson.DynNode, error) {
	var id string
	var err error
	if id, err = this.Str("/id"); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/accounts/%v/address", id)
	return this.client.PostDynNode(path, args)
}

func (this Account) Transactions(page, limit int) ([]Transaction, error) {
	return nil, nil
}

func (this Account) Transaction(id string) (Transaction, error) {
	return Transaction{}, nil
}

func (this Account) Transfers(page, limit int) ([]Transfer, error) {
	return nil, nil
}

func (this Account) Transfer(id string) (Transfer, error) {
	return Transfer{}, nil
}

func (this Account) TransferMoney(args dynjson.DynNode) (Transaction, error) {
	return Transaction{}, nil
}

func (this Account) SendMoney(args dynjson.DynNode, twofauth string) (Transaction, error) {
	return Transaction{}, nil
}

func (this Account) RequestMoney(args dynjson.DynNode) (Transaction, error) {
	return Transaction{}, nil
}

func (this Account) Button(code string) (Button, error) {
	return Button{}, nil
}

func (this Account) NewButton(args dynjson.DynNode) (Button, error) {
	return Button{}, nil
}

func (this Account) Orders(page, limit int) ([]Order, error) {
	return nil, nil
}

func (this Account) Order(id string) (Order, error) {
	return Order{}, nil
}

func (this Account) NewOrder(args dynjson.DynNode) (Order, error) {
	return Order{}, nil
}

func (this Account) Buy(args dynjson.DynNode) (Transfer, error) {
	return Transfer{}, nil
}

func (this Account) Sell(args dynjson.DynNode) (Transfer, error) {
	return Transfer{}, nil
}
