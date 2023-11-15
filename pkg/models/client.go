package models

import "math/rand"

type Client struct {
	ID     int  `json:"id" db:"p_id"`
	Fund   int  `json:"fund" db:"fund"`
	InGame bool `json:"in_game" db:"in_game"`
}

func NewClient(id int) *Client {
	fund := rand.Int31n(100000-10000) + 10000
	if fund < 10000 {
		fund += 10000
	}
	return &Client{ID: id, Fund: int(fund)}
}
