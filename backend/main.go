package main

import (
	"log"
	"os"

	"github.com/stanleyh24/clientmanager/api"
	"github.com/stanleyh24/clientmanager/storage"
	//"github.com/stanleyh24/clientmanager/models"
	//"github.com/stanleyh24/clientmanager/utils"
)

func main() {

	/* body := map[string]string{
		"name":     "prueba",
		"password": "asd096657",
		"profile":  "default",
		"routes":   "",
		"service":  "pppoe",
	}
	query := models.QueryParams{
		Method:   "DELETE",
		Username: "admin",
		Password: "Asd24690@",
		Ip:       "192.168.3.1",
		Resource: "ppp/secret/*6",
		Body:     nil,
	}

	req := utils.Querymaker(query)

	log.Println(utils.SendRequest(req))*/

	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}
	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewStore(dbPool)

	server := api.NewAPIServer(os.Getenv("SERVER_PORT"), *store)

	server.Run()
}
