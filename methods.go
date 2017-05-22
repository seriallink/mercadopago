package mp

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseUrl = "https://api.mercadopago.com"

// Map data por post in the request
type Params map[string]interface{}

// Map extra request headers
type Headers map[string]string

// Make request and return the response
func (c *Client) execute(method string, endpoint string, params Params, headers Headers, model interface{}) error {

	// init vars
	var (
		err error
		data []byte
		url = baseUrl + endpoint
	)

	// validate method type
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		return errors.New("rest: Not supported method")
	}

	// init an empty payload
	var payload = strings.NewReader("")

	// set access_token
	if params == nil {
		params = Params{
			"access_token": c.AuthToken.AccessToken,
		}
	}

	// marshal params
	if data, err = json.Marshal(params); err != nil {
		return err
	}

	// set payload with params
	payload = strings.NewReader(string(data))

	// append access_token
	if endpoint != oauthTokenPath {
		url += "?access_token=" + c.AuthToken.AccessToken
	}

	request, _ := http.NewRequest(method, url, payload)
	request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", "application/json")

	// add headers
	if headers != nil {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// read response
	if data, err = ioutil.ReadAll(response.Body); err != nil {
		return err
	}

	// init MP custom error
	em := &ErrorMessage{}

	// check for error
	if err = json.Unmarshal(data,em); err != nil {
		return err
	}

	// found error
	if em.GetMessage() != "" {
		return errors.New(em.GetMessage())
	}

	// parse data
	return json.Unmarshal(data,model)

}

//// Execute GET requests
//func (c *Client) Get(endpoint string, params Params, headers Headers, data interface{}) error {
//	return c.execute("GET", endpoint, params, headers, data)
//}

// Execute POST requests
func (c *Client) Post(endpoint string, params Params, headers Headers, model interface{}) error {
	return c.execute("POST", endpoint, params, headers, model)
}

//// Execute PUT requests
//func (c *Client) Put(endpoint string, params Params, headers Headers, data interface{}) error {
//	return c.execute("PUT", endpoint, params, headers, data)
//}
//
//// Execute DELETE requests
//func (c *Client) Delete(endpoint string, params Params, headers Headers, data interface{}) error {
//	return c.execute("DELETE", endpoint, params, headers, data)
//}
