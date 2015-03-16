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
	return false, nil
}

func (this Account) SetPrimary() (bool, error) {
	path := fmt.Sprintf("/accounts/%v/primary", this.AsStr("/id"))
	dyn, err := this.client.PostDynNode(path, "")
	if err != nil {
		return false, err
	}
	//ejs TODO make safe path checks, NO PANICS!
	//ejs TODO make safe path checks, NO PANICS!
	//ejs TODO make safe path checks, NO PANICS!
	//ejs TODO make safe path checks, NO PANICS!
	props, err := dyn.Node("/success")
	if err != nil {
		return false, nil
	}
	return props.AsBool(), nil
}

func (this Account) Modify(args dynjson.DynNode) (bool, error) {
	return false, nil
}

func (this Account) Balance() (dynjson.DynNode, error) {
	path := fmt.Sprintf("/accounts/%v/balance", this.AsStr("/id"))
	return this.client.GetDynNode(path, nil)
}

func (this Account) Address() (dynjson.DynNode, error) {
	return nil, nil
}

func (this Account) Addresses() (dynjson.DynNode, error) {
	return nil, nil
}

func (this Account) NewAddress(args dynjson.DynNode) (dynjson.DynNode, error) {
	return nil, nil
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
