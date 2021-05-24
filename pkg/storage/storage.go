package storage

import "github.com/jackc/pgx/v4/pgxpool"

// type Interface interface{
// 	NewTask(Task) (int, error),
// 	Tasks(int, int) ([]Tasks, error)
// }

type Storage struct {
	*pgxpool.Pool
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
	return id, nil
}

func (s *Storage) Tasks(id, authorId int) ([]Task, error) {
	var tasks []Task
	return tasks, nil
}
