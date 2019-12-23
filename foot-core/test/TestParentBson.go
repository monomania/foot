package test

import "gopkg.in/mgo.v2/bson"

type TestParentBson struct {
	P     bson.ObjectId
	PName string
}
