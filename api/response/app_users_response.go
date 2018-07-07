package response

import "github.com/jmatsu/dpg/api/entity"

type AppUsersResponse struct {
	Result struct {
		Usage struct {
			UsedCount uint64 `json:"used"`
			Capacity  uint64 `json:"max"`
		} `json:"usage"`
		Users []entity.UserSummary `json:"users"`
	} `json:"results"`
	Error bool `json:"error"`
}
