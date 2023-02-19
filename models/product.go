package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price float64    `json:"price"`
}

type CreateProduct struct {
	Name  string `json:"name"`
	Price float64    `json:"price"`
}

type UpdateProduct struct {
	Name  string `json:"name"`
	Price float64    `json:"price"`
}

type GetListProduct struct {
	Count    int       `json:"count"`
	Products []Product `json:"products"`
}
