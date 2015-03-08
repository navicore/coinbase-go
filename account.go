package coinbase

import "github.com/chrhlnd/dynjson"

type Account struct {
	AccountBase
}

func NewAccountFromId(id string) Account {
	props := Id(id)
	return Account{AccountBase{Model{Base{}, props}}}
}

func NewAccount(props dynjson.DynNode) Account {
	return Account{AccountBase{Model{Base{}, props}}}
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

func (obj Account) Balance() (dynjson.DynNode, error) {
	return nil, nil
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
