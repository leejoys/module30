package storage

//[ ]ТЗ:
//[ ]Получение списка задач,
//[ ]Получения информации о конкретной задаче по ее номеру,
//[ ]Создание задачи,
//[ ]Обновление задачи,
//[ ]Удаление задачи,
//[ ]Создание массива задач.

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Задача.
type Task struct {
	Id         int
	Opened     int64
	Closed     int64
	AuthorId   int
	AssignedId int
	Title      string
	Content    string
}

// Конструктор, принимает строку подключения к БД.
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

//NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(),
		`INSERT INTO tasks(title,content) 
		VALUES $1,$2 RETURNING id;`, t.Title, t.Content).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

//Tasks возвращает список задач из БД.
func (s *Storage) Tasks(id, authorId int) ([]Task, error) {

	rows, err := s.db.Query(context.Background(),
		`SELECT 
		id,
		opened,
		closed,
		author_id,
		assigned_id,
		title,
		content
		FROM tasks 
		WHERE ($1=0 OR id=$1) AND ($2=0 OR author_id=$2)
		ORDER BY id;`, id, authorId)

	if err != nil {
		return nil, err
	}

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorId,
			&t.AssignedId,
			&t.Title,
			&t.Content,
		)

		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

//Task - Получения информации о конкретной задаче по ее номеру,
func (s *Storage) Task(id int) (Task, error) {
	var t Task
	err := s.db.QueryRow(context.Background(),
		`SELECT 
		id,
		opened,
		closed,
		author_id,
		assigned_id,
		title,
		content
		FROM tasks 
		WHERE id=$1;`,
		id).Scan(
		&t.Id,
		&t.Opened,
		&t.Closed,
		&t.AuthorId,
		&t.AssignedId,
		&t.Title,
		&t.Content,
	)
	if err != nil {
		return Task{}, err
	}

	return t, nil
}

//Update - Обновление полей title и content задачи id,
func (s *Storage) Update(t Task) error {
	_, err := s.db.Exec(context.Background(), `
	UPDATE tasks
	SET title=$2, content=$3, assigned_id=$4
	WHERE id=$1;`, t.Id, t.Title, t.Content, t.AssignedId)
	return err
}

//Delete - Удаление задачи по id
func (s *Storage) Delete(id int) error {
	_, err := s.db.Exec(context.Background(), `
	DELETE FROM tasks 
	WHERE id=$1;`, id)
	return err
}

//NewTasks - Создание массива задач. Возвращает слайс
//с id задач соответственно их порядку на входе.
func (s *Storage) NewTasks(tasks []Task) ([]int, error) {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	var ids []int
	for _, t := range tasks {
		var id int
		err = tx.QueryRow(context.Background(), `
	INSERT INTO	
	tasks(title, content)
	VALUES ($1,$2);`, t.Title, t.Content).Scan(&id)
		if err != nil {
			tx.Rollback(context.Background())
			return nil, err
		}
		ids = append(ids, id)
	}
	tx.Commit(context.Background())
	return ids, nil
}
