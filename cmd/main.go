package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsonDb"
	"log"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)

	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
	err = c.MoneyTransfer(sender, receiver, 500_000)
	if err != nil {
		log.Println(err)
	}
}
