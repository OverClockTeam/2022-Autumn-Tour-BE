package main

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Posts struct {
	Posts []Post
}

func (u *User) NewPost(title, content string) *Post {
	return &Post{
		Title:   title,
		Content: content,
		Author:  u.Name,
	}
}
