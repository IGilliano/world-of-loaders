package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"worldOfLoaders/pkg/repository/repo_models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreatePlayer(player repo_models.Player) (int, error) {
	var id int
	row := ap.db.QueryRow("INSERT INTO players (login, password, class) VALUES ($1, $2,$3) RETURNING id", player.Login, player.Password, player.Class)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	switch player.Class {
	case "loader":
		if err := ap.CreateLoader(id); err != nil {
			return 0, err
		}
	case "client":
		if err := ap.CreateClient(id); err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (ap *AuthPostgres) CreateLoader(id int) error {
	loader := repo_models.NewLoader(id)
	_, err := ap.db.Query("INSERT INTO loaders (p_id, capacity, is_drinking, fatigue, salary) VALUES ($1, $2,$3, $4, $5)", loader.ID, loader.Capacity, loader.IsDrinking, loader.Fatigue, loader.Salary)
	return err
}

func (ap *AuthPostgres) CreateClient(id int) error {
	client := repo_models.NewClient(id)
	_, err := ap.db.Query("INSERT INTO clients (p_id, fund, in_game) VALUES ($1, $2, $3)", client.ID, client.Fund, true)
	return err
}

func (ap *AuthPostgres) GetPlayer(login, password string) (repo_models.Player, error) {
	var player repo_models.Player
	query := fmt.Sprintf(`SELECT * FROM players WHERE login = $1 AND password = $2`)
	err := ap.db.Get(&player, query, login, password)

	return player, err
}

func (ap *AuthPostgres) GetPlayers() ([]*repo_models.Player, error) {
	var users []*repo_models.Player
	rows, err := ap.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var player repo_models.Player
		err = rows.Scan(&player.ID, &player.Login, &player.Password, &player.Class)
		if err != nil {
			return nil, err
		}
		users = append(users, &player)
	}
	return users, nil
}

func (ap *AuthPostgres) PushTasks(tasks []repo_models.Task) ([]int, error) {
	tasksID := make([]int, 0, len(tasks))
	for _, task := range tasks {
		var id int
		row := ap.db.QueryRow("INSERT INTO tasks (name, item, weight, available) VALUES ($1, $2, $3, $4) RETURNING id", task.Name, task.Item, task.Weight, true)
		if err := row.Scan(&id); err != nil {
			log.Printf("Couldnt save task to database, err: %v", err.Error())
			continue
		}
		tasksID = append(tasksID, id)
	}
	return tasksID, nil
}
