package stcvalidator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type MapErrMsg map[string]map[string]string

func validateWithMap(errMsg string, mapErrMsg MapErrMsg) string {
	dataArr := strings.Split(errMsg, ":")
	lenDataArr := len(dataArr)
	if lenDataArr != 3 && lenDataArr != 2 {
		return ""
	}

	structName := dataArr[0]
	tagPassed := dataArr[1]

	for k, v := range mapErrMsg {
		if k == structName {
			for k1, v1 := range v {
				if k1 == tagPassed {
					return v1
				}
			}
		}
	}

	return ""
}

func Validate(s any, mapErrMsg MapErrMsg) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	if len(validationErrors) == 0 {
		return nil
	}

	valErr := validationErrors[0]

	errMsg := fmt.Sprintf(
		"%s:%s:%v",
		valErr.Field(),
		valErr.ActualTag(),
		valErr.Value(),
	)

	if msg := validateWithMap(errMsg, mapErrMsg); msg != "" {
		return errors.New(msg)
	}

	return fmt.Errorf(
		"error when validating %s (%s): '%v'",
		validationErrors[0].Field(),
		validationErrors[0].ActualTag(),
		validationErrors[0].Value(),
	)
}
