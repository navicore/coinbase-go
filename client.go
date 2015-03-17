package coinbase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/chrhlnd/dynjson"
)

type Client struct {
	Base
	http.Client
	Uri    string
	Key    string
	Secret string
	Token  string
	Rtoken string
}

func NewClient(http *http.Client) Client {
	return Client{}
}

func (this Client) Refresh() (bool, error) {
	return false, nil
}

func (this *Client) authHeaders(url string, bodystr string, req *http.Request) {
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	msg := nonce + url + bodystr
	sign := this.getHMAC(msg)

	req.Header.Add("ACCESS_KEY", this.Key)
	req.Header.Add("ACCESS_NONCE", nonce)
	req.Header.Add("ACCESS_SIGNATURE", sign)
}

func (this *Client) headers(url string, req *http.Request) {
	req.Header.Add("User-Agent", "navicore/coinbase-go/1.0")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func (this *Client) getHMAC(msg string) string {
	key_bytes := []byte(this.Secret)
	msg_bytes := []byte(msg)

	mac := hmac.New(sha256.New, key_bytes)
	mac.Write(msg_bytes)

	return hex.EncodeToString(mac.Sum(nil))
}

func (this *Client) Delete(api_method string) ([]byte, error) {
	api_url := this.Uri + api_method

	var req *http.Request
	var err error

	req, err = http.NewRequest("DELETE", api_url, bytes.NewReader([]byte("")))
	if err != nil {
		return nil, err
	}

	this.headers(api_url, req)
	this.authHeaders(api_url, "", req)

	return this.request(req)
}

func (this *Client) Post(api_method string, bodystr string) ([]byte, error) {
	api_url := this.Uri + api_method

	var req *http.Request
	var err error

	req, err = http.NewRequest("POST", api_url, bytes.NewReader([]byte(bodystr)))
	if err != nil {
		return nil, err
	}

	this.headers(api_url, req)
	this.authHeaders(api_url, bodystr, req)

	return this.request(req)
}

func (this *Client) Put(api_method string, bodystr string) ([]byte, error) {
	api_url := this.Uri + api_method

	var req *http.Request
	var err error

	req, err = http.NewRequest("PUT", api_url, bytes.NewReader([]byte(bodystr)))
	if err != nil {
		return nil, err
	}

	this.headers(api_url, req)
	this.authHeaders(api_url, bodystr, req)

	return this.request(req)
}

func (this *Client) Get(method string, params url.Values) ([]byte, error) {

	api_url := this.Uri + method

	if params != nil {
		api_url = "/?" + params.Encode()
	}

	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		return nil, err
	}

	this.headers(api_url, req)
	this.authHeaders(api_url, "", req)

	return this.request(req)
}

func (this *Client) request(req *http.Request) ([]byte, error) {
	resp, err := this.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, CbHttpError{resp.StatusCode}
	}

	return body, nil
}

func (this Client) DelDynNode(api_method string, params string) (dynjson.DynNode, error) {
	buffer, err := this.Delete(api_method)

	if err != nil {
		return nil, err
	}

	return dynjson.NewFromBytes(buffer), nil
}

func (this Client) PutDynNode(api_method string, params string) (dynjson.DynNode, error) {
	buffer, err := this.Put(api_method, params)

	if err != nil {
		return nil, err
	}

	return dynjson.NewFromBytes(buffer), nil
}

func (this Client) PostDynNode(api_method string, params string) (dynjson.DynNode, error) {
	buffer, err := this.Post(api_method, params)

	if err != nil {
		return nil, err
	}

	return dynjson.NewFromBytes(buffer), nil
}

func (this Client) GetDynNode(api_method string, params url.Values) (dynjson.DynNode, error) {
	buffer, err := this.Get(api_method, params)

	if err != nil {
		return nil, err
	}

	return dynjson.NewFromBytes(buffer), nil
}

func (this Client) Accounts() ([]Account, error) {
	var root, props, acctsdyn dynjson.DynNode
	var err error
	if root, err = this.GetDynNode("/accounts", nil); err != nil {
		return nil, err
	}
	if acctsdyn, err = root.Node("/accounts"); err != nil {
		return nil, err
	}
	len := acctsdyn.Len()

	accts := make([]Account, len, len)
	for i := 0; i < len; i++ {
		if props, err = acctsdyn.Node(fmt.Sprintf("/%v", i)); err != nil {
			return nil, err
		}
		accts[i] = NewAccountFromProps(props, this)
	}
	return accts, nil
}

func (this Client) Account(id string) (Account, error) {
	var root, props dynjson.DynNode
	var err error
	path := fmt.Sprintf("/accounts/%v", id)
	if root, err = this.GetDynNode(path, nil); err != nil {
		return Account{}, err
	}
	if props, err = root.Node("/account"); err != nil {
		return Account{}, err
	}
	return NewAccountFromProps(props, this), nil
}

func (this Client) CreateAccount(args string) (Account, error) {
	var root, props dynjson.DynNode
	var err error
	if root, err = this.PostDynNode("/accounts", args); err != nil {
		return Account{}, err
	}
	if props, err = root.Node("/account"); err != nil {
		return Account{}, err
	}
	return NewAccountFromProps(props, this), nil
}

func (this Client) Contacts(page int, limit int, query string) ([]Contact, error) {
	var root, props, node, dyn dynjson.DynNode
	var err error
	path := fmt.Sprintf("/contacts?page=%v&limit=%v", page, limit)
	if query != "" {
		path = path + "&query=" + query
	}
	if root, err = this.GetDynNode(path, nil); err != nil {
		return nil, err
	}
	if dyn, err = root.Node("/contacts"); err != nil {
		return nil, err
	}
	len := dyn.Len()

	contacts := make([]Contact, len, len)
	for i := 0; i < len; i++ {
		if node, err = dyn.Node(fmt.Sprintf("/%v", i)); err != nil {
			return nil, err
		}
		if props, err = node.Node("/contact"); err != nil {
			return nil, err
		}
		contacts[i] = Contact{Model{props: props, client: this}}
	}
	return contacts, nil
}

func (this Client) CurrentUser() (User, error) {
	var root, props dynjson.DynNode
	var err error
	if root, err = this.PostDynNode("/users/self", ""); err != nil {
		return User{}, err
	}
	if props, err = root.Node("/user"); err != nil {
		return User{}, err
	}
	return User{Model{props: props, client: this}}, nil
}

func (this Client) BuyPrice(qty int) (dynjson.DynNode, error) {
	path := fmt.Sprintf("/prices/buy?qty=/%v", qty)
	return this.GetDynNode(path, nil)
}

func (this Client) SellPrice(qty int) (dynjson.DynNode, error) {
	path := fmt.Sprintf("/prices/sell?qty=/%v", qty)
	return this.GetDynNode(path, nil)
}

func (this Client) SpotPrice(currency string) (dynjson.DynNode, error) {
	path := fmt.Sprintf("/prices/spot_rate?currency=%v", currency)
	return this.GetDynNode(path, nil)
}

func (this Client) Currencies() (dynjson.DynNode, error) {
	return this.GetDynNode("/currencies", nil)
}

func (this Client) Rates() (dynjson.DynNode, error) {
	return this.GetDynNode("/exchange_rates", nil)
	return nil, nil
}

func (this Client) CreateUser(args dynjson.DynNode) (User, error) {
	return User{}, nil
}

func (this Client) PayMethods() ([]PayMethod, error) {
	var root, node, dyn, props dynjson.DynNode
	var err error
	if root, err = this.GetDynNode("/payment_methods", nil); err != nil {
		return nil, err
	}
	if dyn, err = root.Node("/payment_methods"); err != nil {
		return nil, err
	}
	len := dyn.Len()

	pms := make([]PayMethod, len, len)
	for i := 0; i < len; i++ {
		if node, err = dyn.Node(fmt.Sprintf("/%v", i)); err != nil {
			return nil, err
		}
		if props, err = node.Node("/payment_method"); err != nil {
			return nil, err
		}
		pms[i] = PayMethod{Model{props: props, client: this}}
	}
	return pms, nil
}

func (this Client) PayMethod(id string) (PayMethod, error) {
	var root, props dynjson.DynNode
	var err error
	path := fmt.Sprintf("/payment_methods/%v", id)
	if root, err = this.GetDynNode(path, nil); err != nil {
		return PayMethod{}, err
	}
	if props, err = root.Node("/payment_method"); err != nil {
		return PayMethod{}, err
	}
	return PayMethod{Model{props: props, client: this}}, nil
}
