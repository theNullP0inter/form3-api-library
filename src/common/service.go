package common

import "github.com/go-playground/validator"

type Service struct {
	Client    *Client
	BasePath  func() string
	Validator *validator.Validate
}
