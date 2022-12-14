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
	"net/http"
)



// AuthApiRouter defines the required methods for binding the api requests to a responses for the AuthApi
// The AuthApiRouter implementation should parse necessary information from the http request,
// pass the data to a AuthApiServicer to perform the required actions, then write the service results to the http response.
type AuthApiRouter interface { 
	GetCurrentUser(http.ResponseWriter, *http.Request)
	LogOut(http.ResponseWriter, *http.Request)
	SignIn(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
}


// AuthApiServicer defines the api actions for the AuthApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AuthApiServicer interface { 
	GetCurrentUser(context.Context) (ImplResponse, error)
	LogOut(context.Context) (ImplResponse, error)
	SignIn(context.Context, SignInForm) (ImplResponse, error)
	SignUp(context.Context, SignUpForm) (ImplResponse, error)
}
