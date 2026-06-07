//go:build validator_novalidatefn

package stcvalidator

func isValidateFn(fl FieldLevel) bool {
	panic("validateFn is not supported with 'no-validate-fn' tag")
}
