package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"worldOfLoaders/pkg/models"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClient(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c *ClientPostgres) GetClientFromDB(ID int) (models.Client, error) {
	var client models.Client
	query := fmt.Sprintf(`SELECT * FROM clients WHERE p_id = $1`)
	err := c.db.Get(&client, query, ID)

	return client, err
}

func (c *ClientPostgres) UpdateClient(client models.Client) error {
	_, err := c.db.Query("UPDATE clients SET fund = $1, in_game = $2 WHERE p_id = $3", client.Fund, client.InGame, client.ID)
	return err
}
