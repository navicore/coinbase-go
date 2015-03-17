package coinbase

import (
	"encoding/json"
	"strconv"

	"github.com/chrhlnd/dynjson"
)

type Model struct {
	Base
	client Client
	props  dynjson.DynNode
}

func (this Model) id() string {
	id, _ := this.props.Node("/id")
	return id.AsStr()
}

func Id(id string) dynjson.DynNode {
	type idprop struct {
		Id string `json:"id"`
	}
	myid := idprop{id}
	b, _ := json.Marshal(myid)
	props := dynjson.NewFromBytes(b)
	return props
}

func (this Model) node(path string) (dynjson.DynNode, error) {
	if this.props == nil {
		return nil, CbError{"props are not set"}
	}
	if this.props.IsNull() {
		return nil, CbError{"props are nil"}
	}
	n, e := this.props.Node(path)
	if e != nil {
		return nil, e
	}
	if n.IsNull() {
		return nil, CbError{"field is nil"}
	}
	return n, nil
}

func (this Model) Float(path string) (float64, error) {
	n, err := this.node(path)
	if err != nil {
		return -1, err
	}
	b, e := n.Str()
	if e != nil {
		return -1, e
	}
	fb, _ := strconv.ParseFloat(b, 64)
	return fb, nil
}

func (this Model) Int(path string) (int, error) {
	n, err := this.node(path)
	if err != nil {
		return -1, err
	}
	return int(n.AsI64()), nil
}

func (this Model) Str(path string) (string, error) {
	n, err := this.node(path)
	if err != nil {
		return "", err
	}
	return n.Str()
}

func (this Model) String() string {
	//TODO: baked??
	if this.props == nil {
		return "nil"
	}
	return string(this.props.Data())
}
