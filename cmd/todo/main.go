package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Kachyr/todo-app"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add new todo")
	complete := flag.Int("complete", 0, "mark todo as complete")
	delete := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todo")

	flag.Parse()

	todos := &todo.TodoList{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
		// log.Fatal(err.Error())
	}

	switch {
	case *add:

		addTodo(todos)
	case *complete > 0:
		completeTodo(todos, *complete)
	case *delete > 0:
		deleteTodo(todos, *delete)
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stderr, "invalid command")
		os.Exit(0)
	}
}

func addTodo(todos *todo.TodoList) {
	task, err := getInput(os.Stdin, flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}

	todos.Add(task)
	storeFile(todos)

}

func completeTodo(todos *todo.TodoList, index int) {
	err := todos.Complete(index)
	if err != nil {
		log.Fatal("cant complete todo")
	}
	storeFile(todos)
}

func deleteTodo(todos *todo.TodoList, index int) {
	err := todos.Delete(index)
	if err != nil {
		log.Fatal("cant delete todo")
	}
	storeFile(todos)
}

func storeFile(todos *todo.TodoList) {
	err := todos.Store(todoFile)
	if err != nil {
		log.Fatal("cant store file")
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	// return "", errors.New("not enough args")

	scanner := bufio.NewScanner(r)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(scanner.Text()) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return scanner.Text(), nil
}
