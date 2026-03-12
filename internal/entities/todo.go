/* ************************************************************************** */
/*                                                                            */
/*                                                           +#               */
/*    todo.go                                               +##               */
/*                                                        +###+               */
/*    By: lucas <lucas@tsiaraso.com>                   +########+             */
/*                                                         +###+              */
/*    Created: 2026/03/12 12:36:09 by lucas                  ##+              */
/*    Updated: 2026/03/12 12:36:09 by lucas                  #+               */
/*                                                                            */
/* ************************************************************************** */

package entities

import "time"

// Todo is the core domain entity.
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTodo(id int, title string) Todo {
	return Todo{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now().UTC(),
	}
}
