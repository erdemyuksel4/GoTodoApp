package core

import (
	"net/http"
)

func Route(pattern string, handler func(HttpVariables, *HttpAPIData), urlPattern UrlPattern, mux *http.ServeMux, middleware ...func(HttpVariables, *HttpAPIData)) {

	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		httpVars := HttpVariables{Writer: w, Request: r, UrlPattern: urlPattern}
		params := Params(httpVars)
		httpVars = HttpVariables{Writer: w, Request: r, pattern: pattern, UrlPattern: urlPattern}
		if middleware != nil {
			for _, m := range middleware {
				m(httpVars, params)
			}
		}
		handler(httpVars, params)
	})
}
func Params(httpVars HttpVariables) *HttpAPIData {

	switch httpVars.Request.Method {
	case "GET":
		return HandlerGET(httpVars)
	case "POST":
		return HandlerPOST(httpVars)
	case "PUT":
		return HandlerUPDATE(httpVars)
	case "PATCH":
		return HandlerPATCH(httpVars)
	case "DELETE":
		return HandlerDELETE(httpVars)
	}
	return nil
}
