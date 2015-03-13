package coinbase

import (
	"encoding/json"

	"github.com/chrhlnd/dynjson"
)

type Model struct {
	Base
	client Client
	props  dynjson.DynNode
}

func (obj Model) id() string {
	id, _ := obj.props.Node("/id")
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

func (this Model) String() string {
	if this.props == nil {
		return "nil"
	}
	return string(this.props.Data())
}
