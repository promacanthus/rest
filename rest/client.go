package rest

import (
	"github.com/go-resty/resty/v2"
)

type Interface interface {
	Post(path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error)
	Put(path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error)
	Patch(path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error)
	Get(path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error)
	Delete(path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error)
}

type RESTClient struct {
	*resty.Client
}

// HTTPClientFor returns an http.Client that use the default transport.
// Will return the default http.DefaultClient if no special case behavior is needed.
func HTTPClientFor(c *Config) *RESTClient {
	client := resty.New()
	client.
		SetRetryCount(c.RetryCount).
		SetRetryWaitTime(c.RetryWaitTime).
		SetRetryMaxWaitTime(c.RetryMaxWaitTime).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		})
	return &RESTClient{
		client,
	}
}

// Post sends a POST request to the specified path with optional query parameters
// and request/response body, and returns the resty.Response and an error.
func (c *RESTClient) Post(
	path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error) {
	request := c.R()
	for key, value := range headers {
		request.SetHeader(key, value)
	}
	if queries != nil {
		request.SetQueryParams(queries)
	}
	return request.SetBody(reqBody).SetResult(respBody).Post(path)
}

// Put sends a PUT request to the specified path with the given query parameters
// and request/response body, and returns the resty.Response and an error.
func (c *RESTClient) Put(
	path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error) {
	request := c.R()
	for key, value := range headers {
		request.SetHeader(key, value)
	}
	if queries != nil {
		request.SetQueryParams(queries)
	}
	return request.SetBody(reqBody).SetResult(respBody).Put(path)
}

// Patch sends a PATCH request to the specified path with optional query parameters
// and request/response body, and returns the resty.Response and an error.
func (c *RESTClient) Patch(
	path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error) {
	request := c.R()
	for key, value := range headers {
		request.SetHeader(key, value)
	}
	if queries != nil {
		request.SetQueryParams(queries)
	}
	return request.SetBody(reqBody).SetResult(respBody).Patch(path)
}

// Get sends a GET request to the specified path with optional query parameters
// and request/response body, and returns the resty.Response and an error.
func (c *RESTClient) Get(
	path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error) {
	request := c.R()
	for key, value := range headers {
		request.SetHeader(key, value)
	}
	if queries != nil {
		request.SetQueryParams(queries)
	}
	return request.SetBody(reqBody).SetResult(respBody).Get(path)
}

// Delete sends a DELETE request to the specified path with optional query parameters
// and request/response body, and returns the resty.Response and an error.
func (c *RESTClient) Delete(
	path string, headers, queries map[string]string, reqBody, respBody any) (*resty.Response, error) {
	request := c.R()
	for key, value := range headers {
		request.SetHeader(key, value)
	}
	if queries != nil {
		request.SetQueryParams(queries)
	}
	return request.SetBody(reqBody).SetResult(respBody).Delete(path)
}
