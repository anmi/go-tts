/*
 * Tiny tracker
 *
 * Tiny tracker server API
 *
 * API version: 1.0.0
 * Contact: anmi.asm@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	AuthApiService := openapi.NewAuthApiService()
	AuthApiController := openapi.NewAuthApiController(AuthApiService)

	router := openapi.NewRouter(AuthApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
