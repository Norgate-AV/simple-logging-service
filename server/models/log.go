package models

import "github.com/Norgate-AV/simple-logging-service/db"

type Log struct {
	Room    string `json:"room" binding:"required"`
	Level   string `json:"level" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func (log *Log) Save() error {
	query := `
	INSERT INTO logs (room, level, message)
	VALUES (?, ?, ?)
	RETURNING id;
	`

	db, err := db.GetDatabase()
	if err != nil {
		return err
	}

	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(log.Room, log.Level, log.Message)
	if err != nil {
		return err
	}

	return nil
}
