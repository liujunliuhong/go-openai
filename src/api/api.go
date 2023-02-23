package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	authToken string
}

func NewClient(authToken string) *Client {
	return &Client{
		authToken: authToken,
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	authorization := fmt.Sprintf("Bearer %s", c.authToken)
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", authorization)

	contentType := req.Header.Get("Content-Type")
	if len(contentType) == 0 {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if v != nil {
		err := json.NewDecoder(res.Body).Decode(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) fullURL(path string) string {
	if !strings.HasPrefix(path, "/") {
		return baseURL + "/" + path
	}
	return baseURL + path
}
