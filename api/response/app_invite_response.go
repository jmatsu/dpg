package response

import "github.com/jmatsu/dpg/api/entity"

type AppInviteResponse struct {
	Result struct {
		Message        string                `json:"invite"`
		AddedRecords   []entity.InviteRecord `json:"added"`
		RemovedRecords []entity.InviteRecord `json:"removed"`
	} `json:"results"`
	Error bool `json:"error"`
}
