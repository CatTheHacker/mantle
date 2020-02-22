package db

import (
	"database/sql"
)

type Channel struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid" sqlite:"text"`
	Position    int    `json:"position" sqlite:"int"`
	Name        string `json:"name" sqlite:"text"`
	Description string `json:"description" sqlite:"text"`
}

func ScanChannel(rows *sql.Rows) *Channel {
	var v Channel
	rows.Scan(&v.ID, &v.UUID, &v.Position, &v.Name, &v.Description)
	return &v
}
