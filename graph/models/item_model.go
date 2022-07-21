package models

type Item struct {
	ID           string  `json:"id"`
	TypeID       string  `json:"typeId"`
	Lon          float64 `json:"lon"`
	Lat          float64 `json:"lat"`
	BrokenID     *string `json:"brokenId"`
	DeleteStatus *int    `json:"deleteStatus"`
}

type Rent struct {
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
	ItemsID    string `json:"itemsId"`
	UsersID    string `json:"usersId"`
}
