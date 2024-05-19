package core

import (
	"encoding/json"
	"net/http"
	"todoapp/internal/core/helpers"
)

func HandlerGET(httpVars HttpVariables) *HttpAPIData {
	if httpVars.Request.Method == "GET" {
		var data HttpAPIData

		url := helpers.URLParse(httpVars.Request.URL.Path, httpVars.UrlPattern.Order)

		data.Query = url["Query"]
		return &data
	}
	http.Error(httpVars.Writer, "Method not allowed", http.StatusMethodNotAllowed)
	return nil
}
func HandlerPOST(httpVars HttpVariables) *HttpAPIData {
	if httpVars.Request.Method == "POST" {
		var data HttpAPIData

		err := json.NewDecoder(httpVars.Request.Body).Decode(&(data.Body))
		if err != nil {
			http.Error(httpVars.Writer, "Error parsing form data", http.StatusBadRequest)
			return nil
		}

		return &data
	}
	http.Error(httpVars.Writer, "Method not allowed", http.StatusMethodNotAllowed)
	return nil
}
func HandlerUPDATE(httpVars HttpVariables) *HttpAPIData {
	if httpVars.Request.Method == "PUT" {
		var data HttpAPIData
		url := helpers.URLParse(httpVars.Request.URL.Path, httpVars.UrlPattern.Order)
		data.Query = url["Query"]

		err := json.NewDecoder(httpVars.Request.Body).Decode(&(data.Body))
		if err != nil {
			http.Error(httpVars.Writer, "Error parsing form data", http.StatusBadRequest)
			return nil
		}
		return &data
	}
	http.Error(httpVars.Writer, "Method not allowed", http.StatusMethodNotAllowed)
	return nil
}
func HandlerPATCH(httpVars HttpVariables) *HttpAPIData {
	if httpVars.Request.Method == "PATCH" {
		var data HttpAPIData
		url := helpers.URLParse(httpVars.Request.URL.Path, httpVars.UrlPattern.Order)
		data.Query = url["Query"]

		err := json.NewDecoder(httpVars.Request.Body).Decode(&(data.Body))
		if err != nil {
			http.Error(httpVars.Writer, "Error parsing form data", http.StatusBadRequest)
			return nil
		}
		return &data
	}
	http.Error(httpVars.Writer, "Method not allowed", http.StatusMethodNotAllowed)
	return nil
}

func HandlerDELETE(httpVars HttpVariables) *HttpAPIData {
	if httpVars.Request.Method == "DELETE" {
		var data HttpAPIData

		url := helpers.URLParse(httpVars.Request.URL.Path, httpVars.UrlPattern.Order)
		data.Query = url["Query"]

		return &data
	}
	http.Error(httpVars.Writer, "Method not allowed", http.StatusMethodNotAllowed)
	return nil
}
