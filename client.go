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
		return nil, fmt.Errorf("Invalid HTTP response code: %d", resp.StatusCode)
	}

	return body, nil
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
	root, err := this.GetDynNode("/accounts", nil)
	if err != nil {
		return nil, err
	}
	acctsdyn, err := root.Node("/accounts")
	if err != nil {
		return nil, err
	}
	len := acctsdyn.Len()

	accts := make([]Account, len, len)
	for i := 0; i < len; i++ {
		acctdyn, err := acctsdyn.Node(fmt.Sprintf("/%v", i))
		if err != nil {
			return nil, err
		}
		accts[i] = NewAccountFromProps(acctdyn, this)
	}
	return accts, nil
}

func (this Client) Account(id string) (Account, error) {
	path := fmt.Sprintf("/accounts/%v", id)
	root, err := this.GetDynNode(path, nil)
	if err != nil {
		return Account{}, err
	}
	acctdyn, err := root.Node("/account")
	if err != nil {
		return Account{}, err
	}
	return NewAccountFromProps(acctdyn, this), nil
}

func (this Client) CreateAccount(args string) (Account, error) {
	root, err := this.PostDynNode("/accounts", args)
	if err != nil {
		return Account{}, err
	}
	acctdyn, err := root.Node("/account")
	if err != nil {
		return Account{}, err
	}
	return NewAccountFromProps(acctdyn, this), nil
}

func (this Client) Contacts(page int, limit int, query string) ([]Contact, error) {
	//TODO: page limit query impl
	root, err := this.GetDynNode("/contacts", nil)
	if err != nil {
		return nil, err
	}
	dyn, err := root.Node("/contacts")
	if err != nil {
		return nil, err
	}
	len := dyn.Len()

	contacts := make([]Contact, len, len)
	for i := 0; i < len; i++ {
		node, err := dyn.Node(fmt.Sprintf("/%v", i))
		if err != nil {
			return nil, err
		}
		props, err := node.Node("/contact")
		contacts[i] = Contact{Model{props: props, client: this}}
	}
	return contacts, nil
}

func (this Client) CurrentUser() (User, error) {
	root, err := this.PostDynNode("/users/self", "")
	if err != nil {
		return User{}, err
	}
	props, err := root.Node("/user")
	if err != nil {
		return User{}, err
	}
	return User{Model{props: props, client: this}}, nil
}

func (this Client) BuyPrice() (dynjson.DynNode, error) {
	return nil, nil
}

func (this Client) SellPrice() (dynjson.DynNode, error) {
	return nil, nil
}

func (this Client) SpotPrice() (dynjson.DynNode, error) {
	return nil, nil
}

func (this Client) Currencies() (dynjson.DynNode, error) {
	return nil, nil
}

func (this Client) Rates() (dynjson.DynNode, error) {
	return nil, nil
}

func (this Client) CreateUser(args dynjson.DynNode) (User, error) {
	return User{}, nil
}

func (this Client) PayMethods() ([]PayMethod, error) {
	return nil, nil
}

func (this Client) PayMethod(id string) (PayMethod, error) {
	return PayMethod{}, nil
}

/*
ClientBase.prototype._setAccessToken = function (url) {
ClientBase.prototype._generateSignature = function (url, bodyStr) {
ClientBase.prototype._generateReqOptions = function (url, body, method, headers) {
ClientBase.prototype._getHttp = function (path, args, callback, headers) {
ClientBase.prototype._postHttp = function (path, body, callback, headers) {
ClientBase.prototype._putHttp = function (path, body, callback, headers) {
ClientBase.prototype._deleteHttp = function (path, callback, headers) {
ClientBase.prototype._getAllHttp = function(opts, callback, headers) {
ClientBase.prototype._getOneHttp = function(args, callback, headers) {
ClientBase.prototype._postOneHttp = function (opts, callback, headers) {
*/
