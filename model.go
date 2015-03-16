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

func (this Model) AsFloat(path string) float64 {
	//ejs TODO make safe path checks, NO PANICS!
	//ejs TODO use error return
	if this.props == nil {
		return -1
	}
	bn, err := this.props.Node(path)
	if err != nil {
		return -1
	}
	b := bn.AsStr()
	fb, _ := strconv.ParseFloat(b, 64)
	return fb
}

func (this Model) AsInt(path string) int {
	//ejs TODO make safe path checks, NO PANICS!
	//ejs TODO use error return
	if this.props == nil {
		return -1
	}
	bn, err := this.props.Node(path)
	if err != nil {
		return -1
	}
	return int(bn.AsI64())
}

func (this Model) AsStr(path string) string {
	//ejs TODO make safe path checks, NO PANICS!
	//ejs TODO use error return
	if this.props == nil {
		return ""
	}
	bn, err := this.props.Node(path)
	if err != nil {
		return ""
	}
	return bn.AsStr()
}

func (this Model) String() string {
	if this.props == nil {
		return "nil"
	}
	return string(this.props.Data())
}
