/* ************************************************************************** */
/*                                                                            */
/*                                                            +#              */
/*   memory_repo_write.go                                   +##               */
/*                                                         +###+              */
/*   By: lucas <lucas@tsiaraso.com>                      +########+           */
/*                                                           +###+            */
/*   Created: 2026/03/12 12:36:09 by lucas                   ##+              */
/*   Updated: 2026/11/26 14:00:36 by lucas                  #+                */
/*                                                                            */
/* ************************************************************************** */

package infrastructure

import "go.clean/internal/entities"

func (r *MemoryRepo) Update(
	id int, title string, done bool,
) (*entities.Todo, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.todos[id]
	if !ok {
		return nil, false
	}
	t.Title = title
	t.Completed = done
	r.todos[id] = t
	return &t, true
}

func (r *MemoryRepo) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.todos[id]; !ok {
		return false
	}
	delete(r.todos, id)
	return true
}
