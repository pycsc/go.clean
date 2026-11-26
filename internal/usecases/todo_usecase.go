/* ************************************************************************** */
/*                                                                            */
/*                                                            +#              */
/*   todo_usecase.go                                        +##               */
/*                                                         +###+              */
/*   By: lucas <lucas@tsiaraso.com>                      +########+           */
/*                                                           +###+            */
/*   Created: 2026/03/12 12:36:09 by lucas                   ##+              */
/*   Updated: 2026/11/26 14:00:36 by lucas                  #+                */
/*                                                                            */
/* ************************************************************************** */

package usecases

import (
	"go.clean/internal/entities"
	"go.clean/internal/interfaces"
)

// TodoUseCase holds all business logic.
type TodoUseCase struct {
	repo interfaces.TodoRepository
}

func NewUseCase(r interfaces.TodoRepository) *TodoUseCase {
	return &TodoUseCase{repo: r}
}

func (u *TodoUseCase) GetAll() []entities.Todo {
	return u.repo.GetAll()
}

func (u *TodoUseCase) GetByID(id int) (*entities.Todo, bool) {
	return u.repo.GetByID(id)
}

func (u *TodoUseCase) Create(title string) entities.Todo {
	return u.repo.Save(entities.NewTodo(0, title))
}
