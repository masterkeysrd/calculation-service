package random

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client interface {
	GenerateRandom(o Operator) (*Response, error)
}

func NewClient(config *Config) Client {
	return &client{
		config: *config,
	}
}

type client struct {
	config Config
}

func (c *client) GenerateRandom(o Operator) (*Response, error) {
	request := c.createRequest(o)
	requestReader, err := c.createRequestReader(request)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.config.ApiUrl, "application/json", requestReader)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var apiResponse Response
	responseDecoder := json.NewDecoder(resp.Body)
	err = responseDecoder.Decode(&apiResponse)

	if err != nil {
		return nil, err
	}

	if apiResponse.Error != nil {
		return nil, fmt.Errorf("error response from random.org: %v", apiResponse.Error)
	}

	return &apiResponse, nil
}

func (c *client) createRequest(o Operator) *request {
	params := o.GetParams()

	params["apiKey"] = c.config.ApiKey

	r := NewRequest(c.config.ApiKey, o.GetMethod(), params)
	return &r
}

func (c *client) createRequestReader(r *request) (*bytes.Reader, error) {
	requestJson, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(requestJson), nil
}
