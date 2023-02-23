package request_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"openai_error"
	"request"
)

func Perform(r request.Request, v interface{}) (err error) {
	newRequest, err := http.NewRequest(r.Method(), r.URL(), r.Body())

	if err != nil {
		return
	}

	for k, v := range r.Headers() {
		newRequest.Header.Set(k, v)
	}

	client := &http.Client{}

	res, err := client.Do(newRequest)

	defer res.Body.Close()

	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {

		var errorResponse openai_error.APIErrorResponse
		err := json.NewDecoder(res.Body).Decode(&errorResponse)
		if errorResponse.Err == nil && err != nil {
			reqErr := &openai_error.RequestError{
				StatusCode: res.StatusCode,
				Err:        err,
			}
			return reqErr
		}
		if err != nil {
			return err
		}
		return fmt.Errorf("newRequest error, statusCode: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&v)

	return
}
