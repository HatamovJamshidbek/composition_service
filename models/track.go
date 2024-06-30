package models

type Track struct {
	Id             string
	Composition_id string
	User_Id        string
	Title          string
	File_Url       string
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
	Limit, Offset  int
}
