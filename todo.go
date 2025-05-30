package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// non-exported type (lowercase)
type item struct {
	task        string
	done        bool
	createdat   time.time
	completedat time.time
}

// exported type (visible outside package) (uppercase)
type list []item

func (l *list) Add(task string)	{ // receiver using a pointer since we want the method to modify the content of the receiver
	t := item {
		task: task,
		done: false,
		createdat: time.now(),
		completedat: time.time{},
	}

	*l = append(*l, t) // dereferencing the pointer to access the underlying slice.
}

func (l *list) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.errorf("item %d does not exist", i)
	}

	ls[i-1].done = true
	ls[i-1].completedat = time.now()

	return nil
}

func (l *list) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.errorf("item does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)
}

func (l *List) Save(filename string) error {
	js, err := json.Marshall(l)
	
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)
}

func (l *list) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)

	if err !nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) = 0 {
		return nil
	}

	return json.Usmarshal(file, l)
}
