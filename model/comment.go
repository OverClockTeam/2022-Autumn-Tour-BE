package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Comment struct {
	Author    string    `json:"author"`
	CommentAt time.Time `json:"commentAt"`
	Content   string    `json:"content"`
}

type Comments struct {
	Comments []Comment
}

func (c *Comments) Scan(v any) error {
	bytes := v.([]byte)
	return json.Unmarshal(bytes, c)
}

func (c *Comments) Value() (value driver.Value, err error) {
	return json.Marshal(c)
}
