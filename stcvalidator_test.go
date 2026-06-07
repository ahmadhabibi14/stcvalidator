package stcvalidator_test

import (
	"testing"

	"github.com/ahmadhabibi14/stcvalidator"
)

var validate = stcvalidator.New()

type User struct {
	Name string `json:"name" validate:"required"`
	Age  int64  `json:"age" validate:"required,max=200"`
}

func TestValidate(t *testing.T) {
	myUser := User{
		Name: "Thor",
		Age:  1500,
	}

	err := validate.StructWithMapErr(myUser, stcvalidator.MapErrMsg{
		"Name": {
			"required": "Name cannot be empty",
		},
		"Age": {
			"required": "Age cannot be empty",
			"max":      "Maximal age is 200 year old",
		},
	})

	if err != nil {
		t.Log(err)
	}
}
