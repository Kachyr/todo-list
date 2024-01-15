package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList []item

func (t *TodoList) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *TodoList) Complete(index int) error {
	item, err := t.findItem(index)
	if err != nil {
		return fmt.Errorf("cant complete: %v", err)
	}

	item.Done = true
	item.CompletedAt = time.Now()
	return nil
}

func (t *TodoList) Delete(index int) error {

	if index < 0 || index > len(*t) {
		return errors.New("invalid index")
	}
	*t = append((*t)[:index-1], (*t)[index:]...)

	return nil
}

func (t *TodoList) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(os.ErrNotExist, err) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return errors.New("empty file")
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t TodoList) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *TodoList) Print() {
	for i, item := range *t {
		fmt.Printf("%d - %s\n", i, item.Task)
	}
}

func (t *TodoList) findItem(index int) (*item, error) {
	if index < 1 || index > len(*t)+1 {
		return nil, errors.New("invalid index")
	}

	return &(*t)[index-1], nil
}
