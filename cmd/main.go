package main

import (
	"fmt"
	"log"
	"module30/pkg/storage"
	"os"
	"strconv"
)

func main() {

	pwd := os.Getenv("pgpass")
	connstr := "postgres://postgres:" + pwd + "@0.0.0.0/catalog3"

	s, err := storage.New(connstr)
	if err != nil {
		log.Fatal(err)
	}

	//Тест на пустой БД
	fmt.Println(s.Tasks(0, 0))
	fmt.Println(s.Task(0))
	fmt.Println("Тест на пустой БД выполнен")

	//Тест создания одной записи и вывода её двумя методами
	id, err := s.NewTask(storage.Task{Title: "first", Content: "first task"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Создание записи выполнено, id=", id)
	fmt.Println(s.Task(id))
	fmt.Println(s.Tasks(0, 0))
	fmt.Println("Вывод содержимого БД выполнен")

	//Тест удаления записи
	err = s.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Tasks(0, 0))
	fmt.Println("Удаление записи выполнено")

	//Тест создания массива записей и вывода их двумя методами
	var tasks []storage.Task
	for i := 2; i <= 10; i++ {
		t := storage.Task{Title: strconv.Itoa(i), Content: strconv.Itoa(i)}
		tasks = append(tasks, t)
	}
	fmt.Println("Создание массива записей выполнено, tasks=", tasks)
	tasksIds, err := s.NewTasks(tasks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Создание записей выполнено, tasksIds=", tasksIds)
	fmt.Println(s.Tasks(0, 0))
	fmt.Println("Вывод содержимого БД выполнен")

	for _, id := range tasksIds {
		fmt.Println(s.Task(id))
		fmt.Println(s.Tasks(id, 0))
	}
	fmt.Println("Вывод содержимого БД выполнен")

	//Тест изменения записи
	err = s.Update(storage.Task{Id: 2, Title: "Second", Content: "Second task"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Tasks(0, 0))
	fmt.Println("Изменение записи выполнено")

	//Удаление записей
	for _, id := range tasksIds {
		err = s.Delete(id)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(s.Tasks(0, 0))
	fmt.Println("Удаление записи выполнено")

}
