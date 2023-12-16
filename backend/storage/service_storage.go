package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/stanleyh24/clientmanager/models"
)

func (s *PostgresStore) GetAllServices() (models.Services, error) {
	sql := "SELECT id,name,price,rate,router_identifier FROM services;"
	rows, err := s.db.QueryContext(context.Background(), sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var services models.Services

	for rows.Next() {

		s, err := scanRowService(rows)

		if err != nil {
			return nil, fmt.Errorf("%s %w", "Service Row Scan(): ", err)
		}
		services = append(services, s)
	}
	return services, nil
}

func (s *PostgresStore) CreateService(service models.Service) (*models.Service, error) {
	sql := "insert into services (id,name,price,rate,router_identifier) values ($1,$2,$3,$4,$5)"

	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	service.ID = ID

	_, err = s.db.QueryContext(context.Background(), sql, service.ID, service.Name, service.Price, service.Rate, service.RouterIdentifier)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *PostgresStore) UpdateService(service models.Service) (*models.Service, error) {
	sql := "UPDATE services SET name=$2, price=$3, rate=$4 WHERE id=$1"
	_, err := s.db.QueryContext(context.Background(), sql, service.ID, service.Name, service.Price, service.Rate)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "UpdateService()", err)
	}
	return &service, nil
}

func (s *PostgresStore) DeleteService(id string) error {
	sql := "DELETE FROM services WHERE id = $1"
	row, err := s.db.Exec(sql, id)

	if err != nil {
		return fmt.Errorf("%s %w", "DeleteService()", err)
	}
	ok, _ := row.RowsAffected()
	if ok < 1 {
		return fmt.Errorf("not found router whit id: %s", id)
	}

	return nil
}

func scanRowService(s scanner) (*models.Service, error) {
	service := &models.Service{}

	err := s.Scan(&service.ID, &service.Name, &service.Price, &service.Rate, &service.RouterIdentifier)
	if err != nil {
		return &models.Service{}, err
	}

	return service, nil
}
