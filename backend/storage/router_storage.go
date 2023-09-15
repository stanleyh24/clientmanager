package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/stanleyh24/clientmanager/models"
)

func (s *Store) GetAllRouter() (models.Routers, error) {
	sql := "SELECT id,name,ip,username,password,created_at, updated_at FROM routers;"
	rows, err := s.db.Query(context.Background(), sql)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var routers models.Routers

	for rows.Next() {
		var router models.Router
		err = rows.Scan(&router.ID, &router.Name, &router.Ip, &router.Username, &router.Password, &router.CreatedAt, &router.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s %w", "Router Row Scan(): ", err)
		}

		routers = append(routers, &router)
	}
	return routers, nil

}

func (s *Store) GetRouterByID(id string) (*models.Router, error) {
	sql := "SELECT id,name,ip,username,password,created_at,updated_at FROM routers where id = $1;"
	var router models.Router
	err := s.db.QueryRow(context.Background(), sql, id).Scan(&router.ID, &router.Name, &router.Ip, &router.Username, &router.Password, &router.CreatedAt, &router.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &router, nil
}

func (s *Store) CreateRouter(r models.Router) (*models.Router, error) {
	sql := "insert into routers (id, name, ip, username, password, created_at) values ($1,$2,$3,$4,$5,$6);"

	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	r.ID = ID
	r.CreatedAt = time.Now().Unix()

	_, err = s.db.Query(context.Background(), sql, r.ID, r.Name, r.Ip, r.Username, r.Password, r.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *Store) UpdateRouter(r models.Router) (*models.Router, error) {
	sql := "update routers set name=$2, ip=$3 ,username=$4, password=$5, updated_at=$6 where id = $1;"
	r.UpdatedAt = time.Now().Unix()

	_, err := s.db.Query(context.Background(), sql, r.ID, r.Name, r.Ip, r.Username, r.Password, r.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("%s %w", "UpdateCategory()", err)
	}

	return &r, nil
}

func (s *Store) DeleteRouter(id string) error {
	sql := "delete from routers where id = $1"
	ok, err := s.db.Exec(context.Background(), sql, id)

	if err != nil {
		return fmt.Errorf("%s %w", "DeleteCategory()", err)
	}

	if ok.RowsAffected() < 1 {
		return fmt.Errorf("not found router whit id: %s", id)
	}

	return nil
}
