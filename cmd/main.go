package main

import (
	"fmt"
	"log"
	"os"
	"storage"
	"strconv"
)

func main() {
	pwd := os.Getenv("pgpass")
	connstr := "postgres://postgres:" + pwd + "@0.0.0.0/catalog3"

	s, err := storage.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Tasks(0, 0))
	fmt.Println(s.Task(0))

	id, err := s.NewTask(storage.Task{Title: "first", Content: "first task"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Task(id))

	var tasks []storage.Task
	for i := 1; i < 10; i++ {
		t := storage.Task{Title: strconv.Itoa(i), Content: strconv.Itoa(i)}
		tasks = append(tasks, t)
	}

	fmt.Println(s.Tasks(0, 0))
	fmt.Println(s.Tasks(0, 0))
	fmt.Println(s.Tasks(0, 0))
}
