package response

import "github.com/jmatsu/dpg/api/entity"

type AppUploadResponse struct {
	Apps  entity.App `json:"results"`
	Error bool       `json:"error"`
}
