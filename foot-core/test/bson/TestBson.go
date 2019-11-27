package bson

import (
	"tesou.io/platform/foot-parent/foot-core/test"
	"gopkg.in/mgo.v2/bson"
)

type ParentBson struct {
	P     bson.ObjectId
	PName string
}

type TestBson struct {
	test.TestParentBson
	A bool
	B int    "myb"
	C string "myc,omitempty"
	D string `bson:",omitempty" json:"jsonkey"`
	E int64  ",minsize"
	F int64  "myf,omitempty,minsize"
}
