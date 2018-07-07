package response

import "github.com/jmatsu/dpg/api/entity"

type OrganizationTeamsListResponse struct {
	Teams []entity.TeamSummary `json:"teams"`
	Error bool                 `json:"error"`
}
