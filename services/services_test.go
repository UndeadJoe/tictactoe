package services

import (
	"testing"
	"labix.org/v2/mgo/bson"
)

func TestMakeMove(t *testing.T) {
	v := MakeMove(bson.ObjectIdHex("58d25a2ecb47275a068f6f32"), 1, 1, 1)

	if v == true {
		t.Error(
			"For", "58d25a2ecb47275a068f6f32",
			"expected", true,
			"got", v,
		)
	}
}
