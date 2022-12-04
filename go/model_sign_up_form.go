/*
 * Tiny tracker
 *
 * Tiny tracker server API
 *
 * API version: 1.0.0
 * Contact: anmi.asm@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SignUpForm struct {

	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`
}

// AssertSignUpFormRequired checks if the required fields are not zero-ed
func AssertSignUpFormRequired(obj SignUpForm) error {
	return nil
}

// AssertRecurseSignUpFormRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SignUpForm (e.g. [][]SignUpForm), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSignUpFormRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSignUpForm, ok := obj.(SignUpForm)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSignUpFormRequired(aSignUpForm)
	})
}