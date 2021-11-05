package models

import (
	"fmt"
	"log"
	"time"
)

type Todo struct {
	ID        uint64
	UserID    uint64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		user_id,
		content) values (?, ?)`
	_, err = Db.Exec(cmd, u.ID, content)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func FindTodo(id uint64) (todo Todo, err error) {
	cmd := `select * from todos where id = ?`
	todo = Todo{}
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Content,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	return todo, err
}

func FindAllTodos() (todos []Todo, err error) {
	cmd := `select * from todos`
	rows, err := Db.Query(cmd)
	fmt.Println("rows =>", rows)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.UserID,
			&todo.Content,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (u *User) FindTodosByUser() (todos []Todo, err error) {
	cmd := `select * from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.UserID,
			&todo.Content,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
