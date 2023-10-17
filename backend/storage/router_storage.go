package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/stanleyh24/clientmanager/models"
)

func (s *PostgresStore) GetAllRouter() (models.Routers, error) {
	query := "SELECT id,name,ip,username,password,created_at, updated_at FROM routers;"
	rows, err := s.db.QueryContext(context.Background(), query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var routers models.Routers

	for rows.Next() {

		r, err := scanRowRouter(rows)

		if err != nil {
			return nil, fmt.Errorf("%s %w", "Router Row Scan(): ", err)
		}
		routers = append(routers, r)
	}
	return routers, nil
}

func (s *PostgresStore) CreateRouter(r models.Router) (*models.Router, error) {
	sql := "insert into routers (id, name, ip, username, password, created_at) values ($1,$2,$3,$4,$5,$6);"

	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	r.ID = ID
	r.CreatedAt = time.Now().Unix()

	_, err = s.db.QueryContext(context.Background(), sql, r.ID, r.Name, r.Ip, r.Username, r.Password, r.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *PostgresStore) UpdateRouter(r models.Router) (*models.Router, error) {
	sql := "update routers set name=$2, ip=$3 ,username=$4, password=$5, updated_at=$6 where id = $1;"
	r.UpdatedAt = time.Now().Unix()

	_, err := s.db.QueryContext(context.Background(), sql, r.ID, r.Name, r.Ip, r.Username, r.Password, r.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("%s %w", "UpdateCategory()", err)
	}

	return &r, nil
}

func (s *PostgresStore) DeleteRouter(id string) error {
	sql := "delete from routers where id = $1"
	row, err := s.db.Exec(sql, id)

	if err != nil {
		return fmt.Errorf("%s %w", "DeleteCategory()", err)
	}
	ok, _ := row.RowsAffected()
	if ok < 1 {
		return fmt.Errorf("not found router whit id: %s", id)
	}

	return nil
}

func scanRowRouter(s scanner) (*models.Router, error) {
	router := &models.Router{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(&router.ID, &router.Name, &router.Ip, &router.Username, &router.Password, &router.CreatedAt, &updatedAtNull)
	if err != nil {
		return &models.Router{}, err
	}

	router.UpdatedAt = updatedAtNull.Int64

	return router, nil

}
