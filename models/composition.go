package models

type Composition struct {
	Id            string
	User_id       string
	Title         string
	Description   string
	Status        string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	Limit, Offset int
}
