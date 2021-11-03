package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        uint64
	UserID    uint64
	Content   string
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		user_id,
		content,
		created_at) values (?, ?, ?)`
	_, err = Db.Exec(cmd, u.ID, content, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func FindTodo(id uint64) (todo Todo, err error) {
	cmd := `select id, user_id, content, created_at from todos where id = ?`
	todo = Todo{}
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Content,
		&todo.CreatedAt,
	)
	return todo, err
}
