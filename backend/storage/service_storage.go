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

func (s *PostgresStore) GetAllServices() (models.Services, error) {
	query := "SELECT id,name,price,rate,created_at, updated_at FROM services;"
	rows, err := s.db.QueryContext(context.Background(), query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var services models.Services

	for rows.Next() {

		s, err := scanRowService(rows)

		if err != nil {
			return nil, fmt.Errorf("%s %w", "Router Row Scan(): ", err)
		}
		services = append(services, s)
	}
	return services, nil
}

func (s *PostgresStore) CreateService(service models.Service) (*models.Service, error) {
	sql := "insert into services (id,name,price,rate,created_at) values ($1,$2,$3,$4,$5)"

	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	service.ID = ID
	service.CreatedAt = time.Now().Unix()
	_, err = s.db.QueryContext(context.Background(), sql, service.ID, service.Name, service.Price, service.Rate, service.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *PostgresStore) UpdateService(id string) (*models.Service, error) {
	return nil, nil
}

func (s *PostgresStore) DeleteService(id string) error {
	return nil
}

func scanRowService(s scanner) (*models.Service, error) {
	service := &models.Service{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(&service.ID, &service.Name, &service.Price, &service.Rate, &service.CreatedAt, &updatedAtNull)
	if err != nil {
		return &models.Service{}, err
	}

	service.UpdatedAt = updatedAtNull.Int64

	return service, nil

}
