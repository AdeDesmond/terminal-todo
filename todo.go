package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []Item

func (l *List) Add(task string) {
	t := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf(" item %d does not exist \n", i)
	}
	//adjust index for0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist on the list \n", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, js, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l) //l is the struct format that we will receive from the file
}
