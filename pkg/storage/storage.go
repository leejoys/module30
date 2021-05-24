package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// type Interface interface{
// 	NewTask(Task) (int, error),
// 	Tasks(int, int) ([]Tasks, error)
// }

type Storage struct {
	db *pgxpool.Pool
}
type Task struct {
	id          int
	opened      int
	closed      int
	author_id   int
	assigned_id int
	title       string
	context     string
}

func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err:=s.db.QueryRow(`INSERT INTO tasks VALUES`,t).Scan(&id)

if err != nil {
	return 0, err
}

	return id, nil
}

func (s *Storage) Tasks(id, authorId int) ([]Task, error) {
	
	rows,err:=s.db.Query(`SELECT 
	id,
	opened,
	closed,
	author_id,
	assigned_id,
	title,
	context
	FROM TABLE tasks 
	WHERE 
	ORDER BY id`,id,authorId)
	
	if err != nil {
		return nil, err
	}

	var tasks []Task
for range rows{
	rows.
}

	return tasks, nil
}

func New(connstr string) (*Storage, error) {

	db, err := pgxpool.Connect(context.Background(), connstr)

	if err != nil {
		return nil, err
	}

	s := &Storage{
		db: db,
	}

	return s, nil
}
