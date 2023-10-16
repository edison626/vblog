package token

import (
	"github.com/edison626/vblog/exception"
)

const (
	TOKEN_COOKIE_NAME  = "access_token"
	TOKEN_GIN_KEY_NAME = "access_token"
)

var (
	CookieNotFound = exception.NewAuthFailed("cookie %s not found", TOKEN_COOKIE_NAME)
)
