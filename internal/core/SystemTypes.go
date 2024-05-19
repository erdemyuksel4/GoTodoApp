package core

import (
	"net/http"
	"net/url"
)

type HttpVariables struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	pattern    string
	UrlPattern UrlPattern
}
type HttpAPIData struct {
	Query string
	Form  url.Values
	Body  interface{}
}
type UrlPattern struct {
	Order []string
}
