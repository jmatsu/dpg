package entity

type LiteUser struct {
	Id             uint64 `json:"id"`
	Name           string `json:"name"`
	ProfileIconUrl string `json:"profile_icon"`
}
