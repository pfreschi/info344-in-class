package tasks

import (
	"database/sql"
)

type PGStore struct {
	DB *sql.DB
}

func (ps *PGStore) Insert(newtask *NewTask) (*Task, error) {
	t := newtask.ToTask()
	tx, err := ps.DB.Begin()
	if err != nil {
		return nil, err
	}

	sql := `insert into tasks 
    (title, createdAt, modifiedAt, complete) 
    values ($1,$2,$3,$4) returning id`
	row := tx.QueryRow(sql, t.Title, t.CreatedAt, t.ModifiedAt, t.Complete)
	err = row.Scan(&t.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	sql = `insert into tags(taskID, tag)
    values($1,$2)`
	for _, tag := range t.Tags {
		_, err := tx.Exec(sql, t.ID, tag)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return t, nil
}

func (ps *PGStore) Get(ID interface{}) (*Task, error) {
	return nil, nil
}

func (ps *PGStore) GetAll() ([]*Task, error) {
	return nil, nil
}

func (ps *PGStore) Update(task *Task) error {
	return nil
}
