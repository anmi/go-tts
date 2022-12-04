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

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// AuthApiService is a service that implements the logic for the AuthApiServicer
// This service should implement the business logic for every endpoint for the AuthApi API.
// Include any external packages or services that will be required by this service.
type AuthApiService struct {
}

// NewAuthApiService creates a default api service
func NewAuthApiService() AuthApiServicer {
	return &AuthApiService{}
}

// GetCurrentUser -
func (s *AuthApiService) GetCurrentUser(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetCurrentUser with the required logic for this service method.
	// Add api_auth_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CurrentUser{}) or use other options such as http.Ok ...
	//return Response(200, CurrentUser{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetCurrentUser method not implemented")
}

// LogOut -
func (s *AuthApiService) LogOut(ctx context.Context) (ImplResponse, error) {
	// TODO - update LogOut with the required logic for this service method.
	// Add api_auth_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("LogOut method not implemented")
}

// SignIn -
func (s *AuthApiService) SignIn(ctx context.Context, signInForm SignInForm) (ImplResponse, error) {
	// TODO - update SignIn with the required logic for this service method.
	// Add api_auth_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("SignIn method not implemented")
}

// SignUp -
func (s *AuthApiService) SignUp(ctx context.Context, signUpForm SignUpForm) (ImplResponse, error) {
	// TODO - update SignUp with the required logic for this service method.
	// Add api_auth_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil
	fmt.Printf("%v, %v", signUpForm.Username, signUpForm.Password)

	return Response(http.StatusNotImplemented, nil), errors.New("SignUp method not implemented")
}
