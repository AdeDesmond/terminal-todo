package main

import (
	"bufio"
	todo2 "example.com"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var todoFileName = ".todo.json"

func main() {
	l := &todo2.List{}

	if os.Getenv("TODO_FILENAME") == "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	add := flag.Bool("add", false, "Add tasks to the todo list")
	deletetask := flag.Int("delete", 0, "tasks to be deleted")
	task := flag.String("task", "", "Task to be included in the ToDo list") //-task
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool developed by desmond\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "Copywrite 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", "Usage information")
		flag.PrintDefaults()
	}

	flag.Parse()

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	switch {
	case *deletetask > 0:
		if err := l.Delete(*deletetask); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		l.Add(t)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	case *list:
		fmt.Print(l)
		//list current to do items
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		//complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		//save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		//	item := strings.Join(os.Args[1:], " ")
		//	l.Add(item)
		//	if err := l.Save(todoFileName); err != nil {
		//		fmt.Fprintf(os.Stderr, "%s\n", err)
		//		os.Exit(1)
		//	}
		//
		fmt.Fprintf(os.Stderr, "%s\n", "invalid option")
		os.Exit(1)
	}

}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("tasks cannot be blind")
	}
	return s.Text(), nil
}
