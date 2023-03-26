// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/bayathy/go-api-server/ent/todo"
)

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Text      string
	CreatedAt *time.Time
	Status    *todo.Status
	Priority  *int
	ChildIDs  []int
	ParentID  *int
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetText(i.Text)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}
