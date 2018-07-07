package entity

type TeamSummary struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	MemberCount uint64 `json:"member_count"`
}
