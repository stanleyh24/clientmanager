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

func (s *PostgresStore) GetAllClients() (models.Clients, error) {
	sql := "SELECT id, first_name,last_name, address,phone,userdata,payment_date,service_id,created_at, updated_at FROM clients"
	rows, err := s.db.QueryContext(context.Background(), sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var clients models.Clients

	for rows.Next() {
		client, err := scanRowClient(rows)

		if err != nil {
			return nil, fmt.Errorf("%s %w", "Client Row Scan(): ", err)
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (s *PostgresStore) CreateClient(client models.Client) (*models.Client, error) {
	sql := "INSERT INTO clients (id, first_name,last_name, address,phone,userdata,payment_date,service_id,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	client.ID = ID
	client.CreatedAt = time.Now().Unix()
	_, err = s.db.QueryContext(
		context.Background(),
		sql,
		client.ID,
		client.First_Name,
		client.Last_name,
		client.Address,
		client.Phone,
		client.User_Data,
		client.Payment_date,
		client.Service_id,
		client.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (s *PostgresStore) UpdateClient(client models.Client) (*models.Client, error) {
	sql := "UPDATE clients SET first_name=$2,last_name=$3, address=$4,phone=$5,userdata=$6,payment_date=$7,service_id=$8,updated_at=$9 WHERE id=$1"
	client.UpdatedAt = time.Now().Unix()

	_, err := s.db.QueryContext(
		context.Background(),
		sql,
		client.ID,
		client.First_Name,
		client.Last_name,
		client.Address,
		client.Phone,
		client.User_Data,
		client.Payment_date,
		client.Service_id,
		client.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (s *PostgresStore) DeleteClient(id string) error {
	sql := "DELETE FROM clients WHERE id = $1"
	row, err := s.db.Exec(sql, id)

	if err != nil {
		return fmt.Errorf("%s %w", "DeleteClient()", err)
	}
	ok, _ := row.RowsAffected()
	if ok < 1 {
		return fmt.Errorf("not found client whit id: %s", id)
	}

	return nil
}

func scanRowClient(s scanner) (*models.Client, error) {
	client := &models.Client{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(&client.ID, &client.First_Name, &client.Last_name, &client.Address, &client.Phone, &client.User_Data, &client.Payment_date, &client.Service_id, &client.CreatedAt, &updatedAtNull)
	if err != nil {
		return &models.Client{}, err
	}

	client.UpdatedAt = updatedAtNull.Int64

	return client, nil

}
