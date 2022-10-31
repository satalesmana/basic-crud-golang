package usermodel

import (
	"app-basic-crud/app/database"
	"database/sql"
)

type Handler interface {
	Insert(data *User) (int64, error)
	GetByid(id string) (User, error)
	GetAll() ([]User, error)
	Delete(id string) (string, error)
	Update(data *User) (string, error)
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Db struct {
	database *sql.DB
}

func NewUserHandler() Handler {
	return &Db{database: database.GetCoon()}
}

// Insert implements Handler
func (db *Db) Insert(data *User) (int64, error) {
	query := `INSERT INTO user(
		name, email, password  
	) VALUES(?, ?, ?)`

	res, err := db.database.Exec(query,
		&data.Name,
		&data.Email,
		&data.Password,
	)

	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (db *Db) GetByid(id string) (User, error) {
	query := `SELECT id,name,email FROM user WHERE id=?`
	var output User

	res, err := db.database.Query(query, id)
	if err != nil {
		return output, err
	}
	defer res.Close()

	for res.Next() {
		if err := res.Scan(&output.ID, &output.Name, &output.Email); err != nil {
			return output, err
		}
	}

	return output, nil
}

func (db *Db) GetAll() ([]User, error) {
	query := `SELECT id,name,email FROM user`
	res, err := db.database.Query(query)
	var output []User
	if err != nil {
		return output, err
	}
	defer res.Close()

	for res.Next() {
		var data User
		if err := res.Scan(&data.ID, &data.Name, &data.Email); err != nil {
			return output, err
		}

		output = append(output, data)
	}

	return output, nil
}

func (db *Db) Delete(id string) (string, error) {
	query := `DELETE FROM user WHERE id=?`
	_, err := db.database.Exec(query, id)
	if err != nil {
		return "", err
	}
	return "Data berhasil dihapus", nil
}

func (db *Db) Update(data *User) (string, error) {
	if data.Password != "" {
		query := `UPDATE user SET name=?, email=?, password=? WHERE  id=?`
		_, err := db.database.Exec(query,
			&data.Name,
			&data.Email,
			&data.Password,
			&data.ID)
		if err != nil {
			return "", err
		}
	} else {
		query := `UPDATE user SET name=?, email=? WHERE  id=?`
		_, err := db.database.Exec(query,
			&data.Name,
			&data.Email,
			&data.ID)
		if err != nil {
			return "", err
		}
	}

	return "Data berhasil perbaharui", nil
}
