package models

type Log struct {
	Room    string `json:"room" binding:"required"`
	Level   string `json:"level" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func (log *Log) Save() error {
	return nil
}
