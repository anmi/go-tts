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

type SignInForm struct {

	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`
}

// AssertSignInFormRequired checks if the required fields are not zero-ed
func AssertSignInFormRequired(obj SignInForm) error {
	return nil
}

// AssertRecurseSignInFormRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SignInForm (e.g. [][]SignInForm), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSignInFormRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSignInForm, ok := obj.(SignInForm)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSignInFormRequired(aSignInForm)
	})
}
