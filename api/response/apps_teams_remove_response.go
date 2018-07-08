package response

type AppsTeamsRemoveResponse struct {
	IsError bool   `json:"error"`
	Message string `json:"message"`
}
