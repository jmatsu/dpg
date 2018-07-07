package entity

type UserSummary struct {
	NameOrEmail string `json:"name"`
	Role        uint   `json:"role"`
}
