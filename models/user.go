package models

type UserPrimaryKey struct {
	Id int `json:"id"`
}

type CreateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type UpdateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type GetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListResponse struct {
	Count int     `json:"count"`
	Users []User `json:"users"`
}
