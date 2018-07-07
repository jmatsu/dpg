package response

import "github.com/jmatsu/dpg/api/entity"

type AppInviteResponse struct {
	Result struct {
		Message        string               `json:"invite"`
		AddedRecords   []entity.UserSummary `json:"added"`
		RemovedRecords []entity.UserSummary `json:"removed"`
	} `json:"results"`
	Error bool `json:"error"`
}
