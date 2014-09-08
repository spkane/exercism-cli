package handlers

import (
	"fmt"

	"github.com/exercism/cli/api"
	"github.com/exercism/cli/config"
)

type Item struct {
	*api.Problem
	dir string
}

func (it *Item) Path() string {
	return fmt.Sprintf("%s/%s", it.dir, it.Problem.ID)
}

type Homework struct {
	Items    []*Item
	template string
}

func NewHomework(problems []*api.Problem, c *config.Config) *Homework {
	hw := Homework{}
	for _, problem := range problems {
		item := &Item{
			Problem: problem,
			dir:     c.Dir,
		}
		hw.Items = append(hw.Items, item)
	}

	hw.template = fmt.Sprintf("%%%ds %%s\n", hw.MaxTitleWidth())
	return &hw
}

func (hw *Homework) Report() {
	for _, item := range hw.Items {
		fmt.Printf(hw.template, item.String(), item.Path())
	}
}

func (hw *Homework) MaxTitleWidth() int {
	var max int
	for _, item := range hw.Items {
		if len(item.String()) > max {
			max = len(item.String())
		}
	}
	return max
}