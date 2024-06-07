package main

import (
	todo2 "example.com"
	"flag"
	"fmt"
	"os"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo2.List{}

	task := flag.String("task", "", "Task to be included in the ToDo list") //-task
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	switch {
	case *list:
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
