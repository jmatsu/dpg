package entity

type InviteRecord struct {
	NameOrEmail string `json:"name"`
	Role        uint   `json:"role"`
}
