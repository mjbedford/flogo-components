package array2string

import (
	"github.com/project-flogo/core/data/coerce"
)

// Settings struct
type Settings struct {
	// Method        string                 `md:"method,required,allowed(GET,POST,PUT,PATCH,DELETE)"` // The HTTP method to invoke
	Delimeter string `md:"delimeter"` // The delimeter
	Prefix    string `md:"prefix"`    // Optional record prefix
	Suffix    string `md:"suffix"`    // Optional record suffix
	// Uri           string                 `md:"uri,required"`                                       // The URI of the service to invoke
	// Headers       map[string]string      `md:"headers"`                                            // The HTTP header parameters
	// Proxy         string                 `md:"proxy"`                                              // The address of the proxy server to be use
	// Timeout       int                    `md:"timeout"`                                            // The request timeout in seconds
	// SkipSSLVerify bool                   `md:"skipSSLVerify"`                                      // Skip SSL validation
	// CertFile      string                 `md:"certFile"`                                           // Path to PEM encoded client certificate
	// KeyFile       string                 `md:"keyFile"`                                            // Path to PEM encoded client key
	// CAFile        string                 `md:"CAFile"`                                             // Path to PEM encoded root certificates file
	// SSLConfig     map[string]interface{} `md:"sslConfig"`                                          // SSL Configuration
}

// Input struct
type Input struct {
	// PathParams  map[string]string `md:"pathParams"`  // The query parameters (e.g., 'id' in http://.../pet?id=someValue )
	// QueryParams map[string]string `md:"queryParams"` // The path parameters (e.g., 'id' in http://.../pet/:id/name )
	// Headers     map[string]string `md:"headers"`     // The HTTP header parameters
	// Content     interface{}       `md:"content"`     // The message content to send. This is only used in POST, PUT, and PATCH
	InputArray map[string]string `md:"inputarray"` // The Input Array
}

// Output struct
type Output struct {
	// Status  int               `md:"status"`  // The HTTP status code
	// Data    interface{}       `md:"data"`    // The HTTP response data
	// Headers map[string]string `md:"headers"` // The HTTP response headers
	// Cookies      []interface{} `md:"cookies"`      // The response cookies (from 'Set-Cookie')
	ResultString string `md:"resultstring"` // the Output Result String
}

// ToMap Input
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"inputarray": i.InputArray,
	}
}

// FromMap Input
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error

	i.InputArray, err = coerce.ToArray(values["inputarray"])
	if err != nil {
		return err
	}
	return nil
}

// ToMap Output
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"resultstring": o.ResultString,
	}
}

// FromMap Output
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	// o.ResultString , err := coerce.ToString(values["resultstring"])
	o.ResultString  , err := coerce.ToString(values["resultstring"]) //= string(values["resultstring"])
	if err != nil {
		return err
	}
	// o.Status, err = coerce.ToInt(values["status"])
	// if err != nil {
	// 	return err
	// }
	// o.Data, _ = values["data"]

	// o.Headers, err = coerce.ToParams(values["headers"])
	// if err != nil {
	// 	return err
	// }

	// o.Cookies, err = coerce.ToArray(values["cookies"])
	// if err != nil {
	// 	return err
	// }

	return nil
}
