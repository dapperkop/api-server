package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ErrorResponse type ...
type ErrorResponse struct {
	Error  string `json:"error"`
	Status bool   `json:"status"`
}

// SuccessResponse type ...
type SuccessResponse struct {
	Data   interface{} `json:"data"`
	Status bool        `json:"status"`
}

func getResponseData(data interface{}, httpStatus int) interface{} {
	var responseData interface{}

	switch {
	case httpStatus < 400:
		responseData = SuccessResponse{
			Data:   data,
			Status: true,
		}
	default:
		responseData = ErrorResponse{
			Error:  fmt.Sprint(data),
			Status: false,
		}
	}

	return responseData
}

// Response func ...
func Response(data interface{}, httpStatus int, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(httpStatus)

	var (
		err      error
		response []byte
	)

	switch r.Header.Get("Content-Type") {
	case "application/json":
		w.Header().Set("Content-Type", "application/json")

		response, err = json.Marshal(getResponseData(data, httpStatus))

		if err != nil {
			Response(errors.New(http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError, w, r)
		}
	default:
		w.Header().Set("Content-Type", "text/plain")

		response = []byte(fmt.Sprint(getResponseData(data, httpStatus)))
	}

	_, err = w.Write(response)

	if err != nil {
		Response(errors.New(http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError, w, r)
	}
}
