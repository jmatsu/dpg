package response

import "gopkg.in/guregu/null.v3"

type ErrorResponse struct {
	IsError bool   `json:"error"`
	Message string `json:"message"`
}

type ErrorResponseMock struct {
	IsError bool        `json:"error"`
	Message null.String `json:"message"`
	//Because null.String `json:"because"`
}

func (mock ErrorResponseMock) Ensure() *ErrorResponse {
	return &ErrorResponse{
		IsError: mock.IsError,
		Message: mock.Message.String,
	}
}
