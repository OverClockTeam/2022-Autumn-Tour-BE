package model

import "time"

type Post struct {
	ID      int64     `json:"id" gorm:"primaryKey"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	Tag     string    `json:"tag" gorm:"null"`
	Comment *Comments `json:"comments" gorm:"null"`
}

func (p Post) NewComment(author, content string) {
	p.Comment.Comments = append(p.Comment.Comments, Comment{
		Author:    author,
		CommentAt: time.Now(),
		Content:   content,
	})
}
