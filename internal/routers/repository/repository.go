package repository

import (
	"context"

	"github.com/stanley24/clientmanager/internal/config"
	"github.com/stanley24/clientmanager/internal/routers/models"
)

func GetAll(ctx context.Context) ([]models.Router, error) {
	db, err := config.CreateConnection()
	if err != nil {
		panic(err)
	}
	qry := `select id,ip,name,username,password from router`

	rows, err := db.QueryContext(ctx, qry)

	if err != nil {
		return nil, err
	}
	routers := []models.Router{}
	for rows.Next() {
		r := models.Router{}
		err := rows.Scan(&r.Id, &r.Ip, &r.Name, &r.Username, &r.Password)

		if err != nil {
			return nil, err
		}
		routers = append(routers, r)
	}
	return routers, nil

}
