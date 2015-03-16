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

func (obj Account) Delete() (bool, error) {
	return false, nil
}

func (obj Account) SetPrimary() (bool, error) {
	return false, nil
}

func (obj Account) Modify(args dynjson.DynNode) (bool, error) {
	return false, nil
}

func (this Account) Balance() (dynjson.DynNode, error) {
	path := fmt.Sprintf("/accounts/%v/balance", this.AsStr("/id"))
	return this.client.GetDynNode(path, nil)
}

func (obj Account) Address() (dynjson.DynNode, error) {
	return nil, nil
}

func (obj Account) Addresses() (dynjson.DynNode, error) {
	return nil, nil
}

func (obj Account) NewAddress(args dynjson.DynNode) (dynjson.DynNode, error) {
	return nil, nil
}

func (obj Account) Transactions(page, limit int) ([]Transaction, error) {
	return nil, nil
}

func (obj Account) Transaction(id string) (Transaction, error) {
	return Transaction{}, nil
}

func (obj Account) Transfers(page, limit int) ([]Transfer, error) {
	return nil, nil
}

func (obj Account) Transfer(id string) (Transfer, error) {
	return Transfer{}, nil
}

func (obj Account) TransferMoney(args dynjson.DynNode) (Transaction, error) {
	return Transaction{}, nil
}

func (obj Account) SendMoney(args dynjson.DynNode, twofauth string) (Transaction, error) {
	return Transaction{}, nil
}

func (obj Account) RequestMoney(args dynjson.DynNode) (Transaction, error) {
	return Transaction{}, nil
}

func (obj Account) Button(code string) (Button, error) {
	return Button{}, nil
}

func (obj Account) NewButton(args dynjson.DynNode) (Button, error) {
	return Button{}, nil
}

func (obj Account) Orders(page, limit int) ([]Order, error) {
	return nil, nil
}

func (obj Account) Order(id string) (Order, error) {
	return Order{}, nil
}

func (obj Account) NewOrder(args dynjson.DynNode) (Order, error) {
	return Order{}, nil
}

func (obj Account) Buy(args dynjson.DynNode) (Transfer, error) {
	return Transfer{}, nil
}

func (obj Account) Sell(args dynjson.DynNode) (Transfer, error) {
	return Transfer{}, nil
}
