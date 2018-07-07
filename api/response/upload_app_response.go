package response

import "github.com/jmatsu/dpg/api/entity"

type UploadAppResponse struct {
	Apps  entity.App `json:"results"`
	Error bool       `json:"error"`
}
