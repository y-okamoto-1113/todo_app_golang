package models

import (
	"log"
	"time"
)

type Session struct {
	ID        uint64
	UUID      string
	UserID    uint64
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		user_id,
		email) values (?, ?, ?)`

	_, err = Db.Exec(cmd1, CreateUUID(), u.ID, u.Email)
	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select * from sessions where user_id = ? and email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.UserID,
		&session.Email,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return session, err
}

func (s *Session) CheckSession() (valid bool, err error) {
	cmd := `select * from sessions where uuid = ?`
	err = Db.QueryRow(cmd, s.UUID).Scan(
		&s.ID,
		&s.UUID,
		&s.UserID,
		&s.Email,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		valid = false
	}
	if s.ID != 0 {
		valid = true
	}
	return valid, err
}

func (s *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, s.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
