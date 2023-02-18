package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()

	store, err := storage.NewFileJson(&cfg)
	if err != nil {
		panic("error while connecting json file")
	}

	c := controller.NewController(&cfg, store)

	// CREATE
	// _, err = c.CreateUser(
	// 	&models.CreateUser{
	// 		Name: "Ibrohim",
	// 		Surname: "Artikov",
	// 	},
	// )
	// if err != nil {
	// 	log.Println("error while CreateUser")
	// }

	// UPDATE
	// err = c.UpdateUser(
	// 	&models.UpdateUser{
	// 		Name:    "Shokhrukh",
	// 		Surname: "Safarov",
	// 	},
	// 	4,
	// )

	// if err != nil {
	// 	log.Println("error while UpdateUser")
	// }

	// GETBYId
	// user, err := c.GetByIdUser(
	// 	&models.UserPrimaryKey{
	// 		Id: 3,
	// 	},
	// )
	// if err != nil {
	// 	log.Println("error while GetById")
	// }
	// fmt.Println(user)

	// GetList
	users, err := c.GetListUser(&models.GetListRequest{
		Offset: 2,
		Limit:  2,
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(users)
}
