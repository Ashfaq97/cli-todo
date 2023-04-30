package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/Ashfaq97/cli-todo"
)

const (
	todoFile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	delete := flag.Int("delete", 0, "deletes a todo item")
	list := flag.Bool("list", false, "list all todos")
	count := flag.Bool("pending", false, "lists the count of pending items")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.LoadFile(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {

	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		todos.Add(task)

		err = todos.StoreInFile(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *complete > 0:
		err := todos.Complete(*complete)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.StoreInFile(todoFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *delete > 0:
		err := todos.Delete(*delete)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.StoreInFile(todoFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *list:
		todos.Print()

	case *count:
		pendingCount := todos.CountPending()

		fmt.Printf("Total number of pending tasks: %d \n", pendingCount)

	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil

}
