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
func (c *Client) execute(method string, path string, params interface{}, headers Headers, model interface{}) error {

	// init vars
	var (
		err  error
		data []byte
		url  = baseUrl + path
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
	if path != oauthTokenPath {
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

	// check for error message
	if err = json.Unmarshal(data, em); err == nil {
		if em.GetMessage() != "" {
			if em.HasCauses() {
				return errors.New(em.Causes[0].Description)
			}
			return errors.New(em.GetMessage())
		}
	}

	// parse data
	return json.Unmarshal(data, model)

}

// Execute GET requests
func (c *Client) Get(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute("GET", path, params, headers, model)
}

// Execute POST requests
func (c *Client) Post(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute("POST", path, params, headers, model)
}

// Execute PUT requests
func (c *Client) Put(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute("PUT", path, params, headers, model)
}

// Execute DELETE requests
func (c *Client) Delete(path string, params interface{}, headers Headers, model interface{}) error {
	return c.execute("DELETE", path, params, headers, model)
}
