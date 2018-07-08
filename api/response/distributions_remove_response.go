package response

type DistributionsRemoveResponse struct {
	IsError bool `json:"error"`
	Result struct {
		Message string `json:"message"`
	} `json:"results"`
}
