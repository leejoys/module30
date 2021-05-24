package storage

type Interface interface{
	NewTask(Task) (int, error),
	Tasks(int, int) ([]Tasks, error)
}

type Storage struct{
	*pgxpool.Pool
}
 type Task struct{
	id int,
opened int,
closed int,
author_id int,
assigned_id int,
title string,
context string
 }
