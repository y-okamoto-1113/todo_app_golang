package models

import (
	"log"
	"time"
)

type User struct {
	ID        uint64
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password) values (?, ?, ?, ?)`

	// Db は base.go で import されてる同じ models パッケージ内の処理なので参照できる。
	_, err = Db.Exec(cmd,
		CreateUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
	)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func FindUser(id uint64) (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
