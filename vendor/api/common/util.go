/*
Package common - Utility for use everywhere in project
*/
package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LicenseText
var LicenseText = "Copyright 2021 dev Nitesh Joshi. All rights reserved."

// JSONMsgContent :
type JSONMsgContent struct {
	License    string `json:"license"`
	StatusCode int    `json:"statuscode"`
	MSG        string `json:"message"`
}

// JSONObjectContent :
type JSONObjectContent struct {
	License    string      `json:"license"`
	StatusCode int         `json:"statuscode"`
	Content    interface{} `json:"content"`
}

// JSONMSG returns a preformatted Object data
func JSONMSG(code int, msg string) []byte {
	jsonString := JSONMsgContent{
		License:    LicenseText,
		StatusCode: code,
		MSG:        msg,
	}

	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	return result
}

// JSONMSGWrappedObj returns an encoded JSON of the object provided.
func JSONMSGWrappedObj(code int, obj interface{}) []byte {
	jsonString := JSONObjectContent{
		License:    LicenseText,
		StatusCode: code,
		Content:    obj,
	}

	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// OutputResponse is a wrapper function for returning a output response that includes a standard text message.
func OutputResponse(w http.ResponseWriter, r *http.Request, code int, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(JSONMSG(code, message))
}

// OutputResponse is a wrapper function for returning a output response that includes a JSON content output.
func OutputResponseeObject(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(JSONMSGWrappedObj(code, obj))
}

// OutputResponseJSONObject is a wrapper function that returns an already prepared JSON object as web response.
func OutputResponseJSONObject(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(obj.([]byte))
}
