package storage

import "github.com/stanleyh24/clientmanager/models"

func (s *PostgresStore) getAllUsers() (models.Clients, error) {
	return nil, nil
}

func (s *PostgresStore) createClient(client models.Client) (*models.Client, error) {
	return nil, nil
}

func (s *PostgresStore) updateClient(id string) (*models.Client, error) {
	return nil, nil
}

func (s *PostgresStore) deleteClient(id string) error {
	return nil
}
